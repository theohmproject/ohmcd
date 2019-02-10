// Copyright (c) 2015-2017 The ohmcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ohmcjson

// SessionResult models the data from the session command.
type SessionResult struct {
	SessionID uint64 `json:"sessionid"`
}

// RescannedBlock contains the hash and all discovered transactions of a single
// rescanned block.
//
// NOTE: This is a ohmcsuite extension ported from
// github.com/decred/dcrd/dcrjson.
type RescannedBlock struct {
	Hash         string   `json:"hash"`
	Transactions []string `json:"transactions"`
}
