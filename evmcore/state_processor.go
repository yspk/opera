// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package evmcore

import (
	"github.com/Fantom-foundation/go-ethereum/common"
	"github.com/Fantom-foundation/go-ethereum/core/state"
	"github.com/Fantom-foundation/go-ethereum/core/types"
	"github.com/Fantom-foundation/go-ethereum/core/vm"
	"github.com/Fantom-foundation/go-ethereum/crypto"
	"github.com/Fantom-foundation/go-ethereum/params"
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *params.ChainConfig // Chain configuration options
	bc     DummyChain          // Canonical block chain
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *params.ChainConfig, bc DummyChain) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
	}
}

// Process processes the state changes according to the Ethereum rules by running
// the transaction messages using the statedb and applying any rewards to both
// the processor (coinbase) and any included uncles.
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of gas that was used in the process. If any of the
// transactions failed to execute due to insufficient gas it will return an error.
func (p *StateProcessor) Process(
	block *EvmBlock, statedb *state.StateDB, cfg vm.Config, internal bool, onNewLog func(*types.Log, *state.StateDB),
) (
	receipts types.Receipts, allLogs []*types.Log, usedGas uint64, skipped []uint32, err error,
) {
	skipped = make([]uint32, 0, len(block.Transactions))
	var (
		gp      = new(GasPool).AddGas(block.GasLimit)
		receipt *types.Receipt
		skip    bool
	)
	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions {
		statedb.Prepare(tx.Hash(), block.Hash, i)
		receipt, _, skip, err = ApplyTransaction(p.config, p.bc, nil, gp, statedb, block.Header(), tx, &usedGas, cfg, internal, onNewLog)
		if skip {
			skipped = append(skipped, uint32(i))
			err = nil
			continue
		}
		if err != nil {
			return
		}
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}

	return
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(
	config *params.ChainConfig,
	bc DummyChain,
	author *common.Address,
	gp *GasPool,
	statedb *state.StateDB,
	header *EvmHeader,
	tx *types.Transaction,
	usedGas *uint64,
	cfg vm.Config,
	internal bool,
	onNewLog func(*types.Log, *state.StateDB),
) (
	*types.Receipt,
	uint64,
	bool,
	error,
) {
	var msg types.Message
	var err error
	if !internal {
		msg, err = tx.AsMessage(types.MakeSigner(config, header.Number))
		if err != nil {
			return nil, 0, false, err
		}
	} else {
		msg = types.NewMessage(common.Address{}, tx.To(), tx.Nonce(), tx.Value(), tx.Gas(), tx.GasPrice(), tx.Data(), false)
	}

	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc, author)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	// Apply the transaction to the current state (included in the env)
	result, err := ApplyMessage(vmenv, msg, gp)
	if err != nil {
		return nil, 0, result == nil, err
	}
	// Notify about logs with potential state changes
	logs := statedb.GetLogs(tx.Hash())
	for _, l := range logs {
		onNewLog(l, statedb)
	}
	// Update the state with pending changes
	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	*usedGas += result.UsedGas

	// Create a new receipt for the transaction, storing the intermediate root and gas used by the tx
	// based on the eip phase, we're passing whether the root touch-delete accounts.
	receipt := types.NewReceipt(root, result.Failed(), *usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = result.UsedGas
	// if the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
	}
	// Set the receipt logs
	receipt.Logs = logs
	receipt.BlockHash = statedb.BlockHash()
	receipt.BlockNumber = header.Number
	receipt.TransactionIndex = uint(statedb.TxIndex())

	return receipt, result.UsedGas, false, err
}
