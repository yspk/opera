package driverpos

import (
	"github.com/Fantom-foundation/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Events
var (
	// Topics of Driver contract logs
	Topics = struct {
		UpdateValidatorWeight common.Hash
		UpdateValidatorPubkey common.Hash
		UpdateNetworkRules    common.Hash
		UpdateNetworkVersion  common.Hash
		AdvanceEpochs         common.Hash
	}{
		UpdateValidatorWeight: common.BytesToHash(crypto.Keccak256Hash([]byte("UpdateValidatorWeight(uint256,uint256)")).Bytes()),
		UpdateValidatorPubkey: common.BytesToHash(crypto.Keccak256Hash([]byte("UpdateValidatorPubkey(uint256,bytes)")).Bytes()),
		UpdateNetworkRules:    common.BytesToHash(crypto.Keccak256Hash([]byte("UpdateNetworkRules(bytes)")).Bytes()),
		UpdateNetworkVersion:  common.BytesToHash(crypto.Keccak256Hash([]byte("UpdateNetworkVersion(uint256)")).Bytes()),
		AdvanceEpochs:         common.BytesToHash(crypto.Keccak256Hash([]byte("AdvanceEpochs(uint256)")).Bytes()),
	}
)
