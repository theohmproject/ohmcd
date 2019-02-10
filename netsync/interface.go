// Copyright (c) 2017 The ohmcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/ohmcsuite/ohmcd/blockchain"
	"github.com/ohmcsuite/ohmcd/chaincfg"
	"github.com/ohmcsuite/ohmcd/chaincfg/chainhash"
	"github.com/ohmcsuite/ohmcd/mempool"
	"github.com/ohmcsuite/ohmcd/peer"
	"github.com/ohmcsuite/ohmcd/wire"
	"github.com/ohmcsuite/ohmcutil"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *ohmcutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockchain.BlockChain
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int

	FeeEstimator *mempool.FeeEstimator
}
