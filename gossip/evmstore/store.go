package evmstore

import (
	"sync"

	"github.com/Fantom-foundation/go-ethereum/common"
	"github.com/Fantom-foundation/go-ethereum/core/rawdb"
	"github.com/Fantom-foundation/go-ethereum/core/state"
	"github.com/Fantom-foundation/go-ethereum/core/types"
	"github.com/Fantom-foundation/go-ethereum/ethdb"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/kvdb"
	"github.com/Fantom-foundation/lachesis-base/kvdb/nokeyiserr"
	"github.com/Fantom-foundation/lachesis-base/kvdb/table"
	"github.com/Fantom-foundation/lachesis-base/utils/wlru"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/Fantom-foundation/go-opera/topicsdb"
	"github.com/Fantom-foundation/go-opera/utils/adapters/kvdb2ethdb"
	"github.com/Fantom-foundation/go-opera/utils/rlpstore"
)

// Store is a node persistent storage working over physical key-value database.
type Store struct {
	cfg StoreConfig

	mainDB kvdb.Store
	table  struct {
		// API-only tables
		Receipts    kvdb.Store `table:"r"`
		TxPositions kvdb.Store `table:"x"`
		Txs         kvdb.Store `table:"X"`

		Evm      ethdb.Database
		EvmState state.Database
		EvmLogs  *topicsdb.Index
	}

	cache struct {
		TxPositions *wlru.Cache `cache:"-"` // store by pointer
		Receipts    *wlru.Cache `cache:"-"` // store by value
	}

	mutex struct {
		Inc sync.Mutex
	}

	rlp rlpstore.Helper

	logger.Instance
}

// NewStore creates store over key-value db.
func NewStore(mainDB kvdb.Store, cfg StoreConfig) *Store {
	s := &Store{
		cfg:      cfg,
		mainDB:   mainDB,
		Instance: logger.MakeInstance(),
		rlp:      rlpstore.Helper{logger.MakeInstance()},
	}

	table.MigrateTables(&s.table, s.mainDB)

	evmTable := nokeyiserr.Wrap(table.New(s.mainDB, []byte("M"))) // ETH expects that "not found" is an error
	s.table.Evm = rawdb.NewDatabase(kvdb2ethdb.Wrap(evmTable))
	s.table.EvmState = state.NewDatabaseWithCache(s.table.Evm, cfg.Cache.EvmDatabase/opt.MiB, "")
	s.table.EvmLogs = topicsdb.New(table.New(s.mainDB, []byte("L")))

	s.initCache()

	return s
}

func (s *Store) initCache() {
	s.cache.Receipts = s.makeCache(s.cfg.Cache.ReceiptsSize, s.cfg.Cache.ReceiptsBlocks)
	s.cache.TxPositions = s.makeCache(uint(s.cfg.Cache.TxPositions), s.cfg.Cache.TxPositions)
}

// Commit changes.
func (s *Store) Commit(root hash.Hash) error {
	// Flush trie on the DB
	err := s.table.EvmState.TrieDB().Commit(common.Hash(root), false, nil)
	if err != nil {
		s.Log.Error("Failed to flush trie DB into main DB", "err", err)
	}
	return err
}

func (s *Store) Cap(max, min int) {
	maxSize := common.StorageSize(max)
	minSize := common.StorageSize(min)
	size, preimagesSize := s.table.EvmState.TrieDB().Size()
	if size >= maxSize || preimagesSize >= maxSize {
		_ = s.table.EvmState.TrieDB().Cap(minSize)
	}
}

// StateDB returns state database.
func (s *Store) StateDB(from hash.Hash) (*state.StateDB, error) {
	return state.New(common.Hash(from), s.table.EvmState, nil)
}

// IndexLogs indexes EVM logs
func (s *Store) IndexLogs(recs ...*types.Log) {
	err := s.table.EvmLogs.Push(recs...)
	if err != nil {
		s.Log.Crit("DB logs index error", "err", err)
	}
}

func (s *Store) EvmTable() ethdb.Database {
	return s.table.Evm
}

func (s *Store) EvmLogs() *topicsdb.Index {
	return s.table.EvmLogs
}

/*
 * Utils:
 */

func (s *Store) makeCache(weight uint, size int) *wlru.Cache {
	cache, err := wlru.New(weight, size)
	if err != nil {
		s.Log.Crit("Failed to create LRU cache", "err", err)
		return nil
	}
	return cache
}
