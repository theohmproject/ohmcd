// Copyright (c) 2014 The ohmcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ohmcjson_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/ohmcsuite/ohmcd/ohmcjson"
)

// TestWalletSvrCmds tests all of the wallet server commands marshal and
// unmarshal into valid results include handling of optional fields being
// omitted in the marshalled command, while optional fields with defaults have
// the default assigned on unmarshalled commands.
func TestWalletSvrCmds(t *testing.T) {
	t.Parallel()

	testID := int(1)
	tests := []struct {
		name         string
		newCmd       func() (interface{}, error)
		staticCmd    func() interface{}
		marshalled   string
		unmarshalled interface{}
	}{
		{
			name: "addmultisigaddress",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("addmultisigaddress", 2, []string{"031234", "035678"})
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return ohmcjson.NewAddMultisigAddressCmd(2, keys, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"addmultisigaddress","params":[2,["031234","035678"]],"id":1}`,
			unmarshalled: &ohmcjson.AddMultisigAddressCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
				Account:   nil,
			},
		},
		{
			name: "addmultisigaddress optional",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("addmultisigaddress", 2, []string{"031234", "035678"}, "test")
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return ohmcjson.NewAddMultisigAddressCmd(2, keys, ohmcjson.String("test"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"addmultisigaddress","params":[2,["031234","035678"],"test"],"id":1}`,
			unmarshalled: &ohmcjson.AddMultisigAddressCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
				Account:   ohmcjson.String("test"),
			},
		},
		{
			name: "addwitnessaddress",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("addwitnessaddress", "1address")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewAddWitnessAddressCmd("1address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"addwitnessaddress","params":["1address"],"id":1}`,
			unmarshalled: &ohmcjson.AddWitnessAddressCmd{
				Address: "1address",
			},
		},
		{
			name: "createmultisig",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("createmultisig", 2, []string{"031234", "035678"})
			},
			staticCmd: func() interface{} {
				keys := []string{"031234", "035678"}
				return ohmcjson.NewCreateMultisigCmd(2, keys)
			},
			marshalled: `{"jsonrpc":"1.0","method":"createmultisig","params":[2,["031234","035678"]],"id":1}`,
			unmarshalled: &ohmcjson.CreateMultisigCmd{
				NRequired: 2,
				Keys:      []string{"031234", "035678"},
			},
		},
		{
			name: "dumpprivkey",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("dumpprivkey", "1Address")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewDumpPrivKeyCmd("1Address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"dumpprivkey","params":["1Address"],"id":1}`,
			unmarshalled: &ohmcjson.DumpPrivKeyCmd{
				Address: "1Address",
			},
		},
		{
			name: "encryptwallet",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("encryptwallet", "pass")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewEncryptWalletCmd("pass")
			},
			marshalled: `{"jsonrpc":"1.0","method":"encryptwallet","params":["pass"],"id":1}`,
			unmarshalled: &ohmcjson.EncryptWalletCmd{
				Passphrase: "pass",
			},
		},
		{
			name: "estimatefee",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("estimatefee", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewEstimateFeeCmd(6)
			},
			marshalled: `{"jsonrpc":"1.0","method":"estimatefee","params":[6],"id":1}`,
			unmarshalled: &ohmcjson.EstimateFeeCmd{
				NumBlocks: 6,
			},
		},
		{
			name: "estimatepriority",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("estimatepriority", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewEstimatePriorityCmd(6)
			},
			marshalled: `{"jsonrpc":"1.0","method":"estimatepriority","params":[6],"id":1}`,
			unmarshalled: &ohmcjson.EstimatePriorityCmd{
				NumBlocks: 6,
			},
		},
		{
			name: "getaccount",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getaccount", "1Address")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetAccountCmd("1Address")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaccount","params":["1Address"],"id":1}`,
			unmarshalled: &ohmcjson.GetAccountCmd{
				Address: "1Address",
			},
		},
		{
			name: "getaccountaddress",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getaccountaddress", "acct")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetAccountAddressCmd("acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaccountaddress","params":["acct"],"id":1}`,
			unmarshalled: &ohmcjson.GetAccountAddressCmd{
				Account: "acct",
			},
		},
		{
			name: "getaddressesbyaccount",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getaddressesbyaccount", "acct")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetAddressesByAccountCmd("acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"getaddressesbyaccount","params":["acct"],"id":1}`,
			unmarshalled: &ohmcjson.GetAddressesByAccountCmd{
				Account: "acct",
			},
		},
		{
			name: "getbalance",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getbalance")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetBalanceCmd(nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":[],"id":1}`,
			unmarshalled: &ohmcjson.GetBalanceCmd{
				Account: nil,
				MinConf: ohmcjson.Int(1),
			},
		},
		{
			name: "getbalance optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getbalance", "acct")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetBalanceCmd(ohmcjson.String("acct"), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":["acct"],"id":1}`,
			unmarshalled: &ohmcjson.GetBalanceCmd{
				Account: ohmcjson.String("acct"),
				MinConf: ohmcjson.Int(1),
			},
		},
		{
			name: "getbalance optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getbalance", "acct", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetBalanceCmd(ohmcjson.String("acct"), ohmcjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getbalance","params":["acct",6],"id":1}`,
			unmarshalled: &ohmcjson.GetBalanceCmd{
				Account: ohmcjson.String("acct"),
				MinConf: ohmcjson.Int(6),
			},
		},
		{
			name: "getnewaddress",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getnewaddress")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetNewAddressCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getnewaddress","params":[],"id":1}`,
			unmarshalled: &ohmcjson.GetNewAddressCmd{
				Account: nil,
			},
		},
		{
			name: "getnewaddress optional",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getnewaddress", "acct")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetNewAddressCmd(ohmcjson.String("acct"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getnewaddress","params":["acct"],"id":1}`,
			unmarshalled: &ohmcjson.GetNewAddressCmd{
				Account: ohmcjson.String("acct"),
			},
		},
		{
			name: "getrawchangeaddress",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getrawchangeaddress")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetRawChangeAddressCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawchangeaddress","params":[],"id":1}`,
			unmarshalled: &ohmcjson.GetRawChangeAddressCmd{
				Account: nil,
			},
		},
		{
			name: "getrawchangeaddress optional",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getrawchangeaddress", "acct")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetRawChangeAddressCmd(ohmcjson.String("acct"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getrawchangeaddress","params":["acct"],"id":1}`,
			unmarshalled: &ohmcjson.GetRawChangeAddressCmd{
				Account: ohmcjson.String("acct"),
			},
		},
		{
			name: "getreceivedbyaccount",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getreceivedbyaccount", "acct")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetReceivedByAccountCmd("acct", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaccount","params":["acct"],"id":1}`,
			unmarshalled: &ohmcjson.GetReceivedByAccountCmd{
				Account: "acct",
				MinConf: ohmcjson.Int(1),
			},
		},
		{
			name: "getreceivedbyaccount optional",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getreceivedbyaccount", "acct", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetReceivedByAccountCmd("acct", ohmcjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaccount","params":["acct",6],"id":1}`,
			unmarshalled: &ohmcjson.GetReceivedByAccountCmd{
				Account: "acct",
				MinConf: ohmcjson.Int(6),
			},
		},
		{
			name: "getreceivedbyaddress",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getreceivedbyaddress", "1Address")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetReceivedByAddressCmd("1Address", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaddress","params":["1Address"],"id":1}`,
			unmarshalled: &ohmcjson.GetReceivedByAddressCmd{
				Address: "1Address",
				MinConf: ohmcjson.Int(1),
			},
		},
		{
			name: "getreceivedbyaddress optional",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getreceivedbyaddress", "1Address", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetReceivedByAddressCmd("1Address", ohmcjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"getreceivedbyaddress","params":["1Address",6],"id":1}`,
			unmarshalled: &ohmcjson.GetReceivedByAddressCmd{
				Address: "1Address",
				MinConf: ohmcjson.Int(6),
			},
		},
		{
			name: "gettransaction",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("gettransaction", "123")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetTransactionCmd("123", nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"gettransaction","params":["123"],"id":1}`,
			unmarshalled: &ohmcjson.GetTransactionCmd{
				Txid:             "123",
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "gettransaction optional",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("gettransaction", "123", true)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetTransactionCmd("123", ohmcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"gettransaction","params":["123",true],"id":1}`,
			unmarshalled: &ohmcjson.GetTransactionCmd{
				Txid:             "123",
				IncludeWatchOnly: ohmcjson.Bool(true),
			},
		},
		{
			name: "getwalletinfo",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("getwalletinfo")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewGetWalletInfoCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"getwalletinfo","params":[],"id":1}`,
			unmarshalled: &ohmcjson.GetWalletInfoCmd{},
		},
		{
			name: "importprivkey",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("importprivkey", "abc")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewImportPrivKeyCmd("abc", nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc"],"id":1}`,
			unmarshalled: &ohmcjson.ImportPrivKeyCmd{
				PrivKey: "abc",
				Label:   nil,
				Rescan:  ohmcjson.Bool(true),
			},
		},
		{
			name: "importprivkey optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("importprivkey", "abc", "label")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewImportPrivKeyCmd("abc", ohmcjson.String("label"), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label"],"id":1}`,
			unmarshalled: &ohmcjson.ImportPrivKeyCmd{
				PrivKey: "abc",
				Label:   ohmcjson.String("label"),
				Rescan:  ohmcjson.Bool(true),
			},
		},
		{
			name: "importprivkey optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("importprivkey", "abc", "label", false)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewImportPrivKeyCmd("abc", ohmcjson.String("label"), ohmcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"importprivkey","params":["abc","label",false],"id":1}`,
			unmarshalled: &ohmcjson.ImportPrivKeyCmd{
				PrivKey: "abc",
				Label:   ohmcjson.String("label"),
				Rescan:  ohmcjson.Bool(false),
			},
		},
		{
			name: "keypoolrefill",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("keypoolrefill")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewKeyPoolRefillCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"keypoolrefill","params":[],"id":1}`,
			unmarshalled: &ohmcjson.KeyPoolRefillCmd{
				NewSize: ohmcjson.Uint(100),
			},
		},
		{
			name: "keypoolrefill optional",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("keypoolrefill", 200)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewKeyPoolRefillCmd(ohmcjson.Uint(200))
			},
			marshalled: `{"jsonrpc":"1.0","method":"keypoolrefill","params":[200],"id":1}`,
			unmarshalled: &ohmcjson.KeyPoolRefillCmd{
				NewSize: ohmcjson.Uint(200),
			},
		},
		{
			name: "listaccounts",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listaccounts")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListAccountsCmd(nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listaccounts","params":[],"id":1}`,
			unmarshalled: &ohmcjson.ListAccountsCmd{
				MinConf: ohmcjson.Int(1),
			},
		},
		{
			name: "listaccounts optional",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listaccounts", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListAccountsCmd(ohmcjson.Int(6))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listaccounts","params":[6],"id":1}`,
			unmarshalled: &ohmcjson.ListAccountsCmd{
				MinConf: ohmcjson.Int(6),
			},
		},
		{
			name: "listaddressgroupings",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listaddressgroupings")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListAddressGroupingsCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"listaddressgroupings","params":[],"id":1}`,
			unmarshalled: &ohmcjson.ListAddressGroupingsCmd{},
		},
		{
			name: "listlockunspent",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listlockunspent")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListLockUnspentCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"listlockunspent","params":[],"id":1}`,
			unmarshalled: &ohmcjson.ListLockUnspentCmd{},
		},
		{
			name: "listreceivedbyaccount",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listreceivedbyaccount")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListReceivedByAccountCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[],"id":1}`,
			unmarshalled: &ohmcjson.ListReceivedByAccountCmd{
				MinConf:          ohmcjson.Int(1),
				IncludeEmpty:     ohmcjson.Bool(false),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaccount optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listreceivedbyaccount", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListReceivedByAccountCmd(ohmcjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6],"id":1}`,
			unmarshalled: &ohmcjson.ListReceivedByAccountCmd{
				MinConf:          ohmcjson.Int(6),
				IncludeEmpty:     ohmcjson.Bool(false),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaccount optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listreceivedbyaccount", 6, true)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListReceivedByAccountCmd(ohmcjson.Int(6), ohmcjson.Bool(true), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6,true],"id":1}`,
			unmarshalled: &ohmcjson.ListReceivedByAccountCmd{
				MinConf:          ohmcjson.Int(6),
				IncludeEmpty:     ohmcjson.Bool(true),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaccount optional3",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listreceivedbyaccount", 6, true, false)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListReceivedByAccountCmd(ohmcjson.Int(6), ohmcjson.Bool(true), ohmcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaccount","params":[6,true,false],"id":1}`,
			unmarshalled: &ohmcjson.ListReceivedByAccountCmd{
				MinConf:          ohmcjson.Int(6),
				IncludeEmpty:     ohmcjson.Bool(true),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listreceivedbyaddress")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListReceivedByAddressCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[],"id":1}`,
			unmarshalled: &ohmcjson.ListReceivedByAddressCmd{
				MinConf:          ohmcjson.Int(1),
				IncludeEmpty:     ohmcjson.Bool(false),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listreceivedbyaddress", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListReceivedByAddressCmd(ohmcjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6],"id":1}`,
			unmarshalled: &ohmcjson.ListReceivedByAddressCmd{
				MinConf:          ohmcjson.Int(6),
				IncludeEmpty:     ohmcjson.Bool(false),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listreceivedbyaddress", 6, true)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListReceivedByAddressCmd(ohmcjson.Int(6), ohmcjson.Bool(true), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6,true],"id":1}`,
			unmarshalled: &ohmcjson.ListReceivedByAddressCmd{
				MinConf:          ohmcjson.Int(6),
				IncludeEmpty:     ohmcjson.Bool(true),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listreceivedbyaddress optional3",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listreceivedbyaddress", 6, true, false)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListReceivedByAddressCmd(ohmcjson.Int(6), ohmcjson.Bool(true), ohmcjson.Bool(false))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listreceivedbyaddress","params":[6,true,false],"id":1}`,
			unmarshalled: &ohmcjson.ListReceivedByAddressCmd{
				MinConf:          ohmcjson.Int(6),
				IncludeEmpty:     ohmcjson.Bool(true),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listsinceblock",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listsinceblock")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListSinceBlockCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":[],"id":1}`,
			unmarshalled: &ohmcjson.ListSinceBlockCmd{
				BlockHash:           nil,
				TargetConfirmations: ohmcjson.Int(1),
				IncludeWatchOnly:    ohmcjson.Bool(false),
			},
		},
		{
			name: "listsinceblock optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listsinceblock", "123")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListSinceBlockCmd(ohmcjson.String("123"), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123"],"id":1}`,
			unmarshalled: &ohmcjson.ListSinceBlockCmd{
				BlockHash:           ohmcjson.String("123"),
				TargetConfirmations: ohmcjson.Int(1),
				IncludeWatchOnly:    ohmcjson.Bool(false),
			},
		},
		{
			name: "listsinceblock optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listsinceblock", "123", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListSinceBlockCmd(ohmcjson.String("123"), ohmcjson.Int(6), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123",6],"id":1}`,
			unmarshalled: &ohmcjson.ListSinceBlockCmd{
				BlockHash:           ohmcjson.String("123"),
				TargetConfirmations: ohmcjson.Int(6),
				IncludeWatchOnly:    ohmcjson.Bool(false),
			},
		},
		{
			name: "listsinceblock optional3",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listsinceblock", "123", 6, true)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListSinceBlockCmd(ohmcjson.String("123"), ohmcjson.Int(6), ohmcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listsinceblock","params":["123",6,true],"id":1}`,
			unmarshalled: &ohmcjson.ListSinceBlockCmd{
				BlockHash:           ohmcjson.String("123"),
				TargetConfirmations: ohmcjson.Int(6),
				IncludeWatchOnly:    ohmcjson.Bool(true),
			},
		},
		{
			name: "listtransactions",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listtransactions")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListTransactionsCmd(nil, nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":[],"id":1}`,
			unmarshalled: &ohmcjson.ListTransactionsCmd{
				Account:          nil,
				Count:            ohmcjson.Int(10),
				From:             ohmcjson.Int(0),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listtransactions", "acct")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListTransactionsCmd(ohmcjson.String("acct"), nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct"],"id":1}`,
			unmarshalled: &ohmcjson.ListTransactionsCmd{
				Account:          ohmcjson.String("acct"),
				Count:            ohmcjson.Int(10),
				From:             ohmcjson.Int(0),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listtransactions", "acct", 20)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListTransactionsCmd(ohmcjson.String("acct"), ohmcjson.Int(20), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20],"id":1}`,
			unmarshalled: &ohmcjson.ListTransactionsCmd{
				Account:          ohmcjson.String("acct"),
				Count:            ohmcjson.Int(20),
				From:             ohmcjson.Int(0),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional3",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listtransactions", "acct", 20, 1)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListTransactionsCmd(ohmcjson.String("acct"), ohmcjson.Int(20),
					ohmcjson.Int(1), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20,1],"id":1}`,
			unmarshalled: &ohmcjson.ListTransactionsCmd{
				Account:          ohmcjson.String("acct"),
				Count:            ohmcjson.Int(20),
				From:             ohmcjson.Int(1),
				IncludeWatchOnly: ohmcjson.Bool(false),
			},
		},
		{
			name: "listtransactions optional4",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listtransactions", "acct", 20, 1, true)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListTransactionsCmd(ohmcjson.String("acct"), ohmcjson.Int(20),
					ohmcjson.Int(1), ohmcjson.Bool(true))
			},
			marshalled: `{"jsonrpc":"1.0","method":"listtransactions","params":["acct",20,1,true],"id":1}`,
			unmarshalled: &ohmcjson.ListTransactionsCmd{
				Account:          ohmcjson.String("acct"),
				Count:            ohmcjson.Int(20),
				From:             ohmcjson.Int(1),
				IncludeWatchOnly: ohmcjson.Bool(true),
			},
		},
		{
			name: "listunspent",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listunspent")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListUnspentCmd(nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[],"id":1}`,
			unmarshalled: &ohmcjson.ListUnspentCmd{
				MinConf:   ohmcjson.Int(1),
				MaxConf:   ohmcjson.Int(9999999),
				Addresses: nil,
			},
		},
		{
			name: "listunspent optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listunspent", 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListUnspentCmd(ohmcjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6],"id":1}`,
			unmarshalled: &ohmcjson.ListUnspentCmd{
				MinConf:   ohmcjson.Int(6),
				MaxConf:   ohmcjson.Int(9999999),
				Addresses: nil,
			},
		},
		{
			name: "listunspent optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listunspent", 6, 100)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListUnspentCmd(ohmcjson.Int(6), ohmcjson.Int(100), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6,100],"id":1}`,
			unmarshalled: &ohmcjson.ListUnspentCmd{
				MinConf:   ohmcjson.Int(6),
				MaxConf:   ohmcjson.Int(100),
				Addresses: nil,
			},
		},
		{
			name: "listunspent optional3",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("listunspent", 6, 100, []string{"1Address", "1Address2"})
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewListUnspentCmd(ohmcjson.Int(6), ohmcjson.Int(100),
					&[]string{"1Address", "1Address2"})
			},
			marshalled: `{"jsonrpc":"1.0","method":"listunspent","params":[6,100,["1Address","1Address2"]],"id":1}`,
			unmarshalled: &ohmcjson.ListUnspentCmd{
				MinConf:   ohmcjson.Int(6),
				MaxConf:   ohmcjson.Int(100),
				Addresses: &[]string{"1Address", "1Address2"},
			},
		},
		{
			name: "lockunspent",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("lockunspent", true, `[{"txid":"123","vout":1}]`)
			},
			staticCmd: func() interface{} {
				txInputs := []ohmcjson.TransactionInput{
					{Txid: "123", Vout: 1},
				}
				return ohmcjson.NewLockUnspentCmd(true, txInputs)
			},
			marshalled: `{"jsonrpc":"1.0","method":"lockunspent","params":[true,[{"txid":"123","vout":1}]],"id":1}`,
			unmarshalled: &ohmcjson.LockUnspentCmd{
				Unlock: true,
				Transactions: []ohmcjson.TransactionInput{
					{Txid: "123", Vout: 1},
				},
			},
		},
		{
			name: "move",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("move", "from", "to", 0.5)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewMoveCmd("from", "to", 0.5, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"move","params":["from","to",0.5],"id":1}`,
			unmarshalled: &ohmcjson.MoveCmd{
				FromAccount: "from",
				ToAccount:   "to",
				Amount:      0.5,
				MinConf:     ohmcjson.Int(1),
				Comment:     nil,
			},
		},
		{
			name: "move optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("move", "from", "to", 0.5, 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewMoveCmd("from", "to", 0.5, ohmcjson.Int(6), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"move","params":["from","to",0.5,6],"id":1}`,
			unmarshalled: &ohmcjson.MoveCmd{
				FromAccount: "from",
				ToAccount:   "to",
				Amount:      0.5,
				MinConf:     ohmcjson.Int(6),
				Comment:     nil,
			},
		},
		{
			name: "move optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("move", "from", "to", 0.5, 6, "comment")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewMoveCmd("from", "to", 0.5, ohmcjson.Int(6), ohmcjson.String("comment"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"move","params":["from","to",0.5,6,"comment"],"id":1}`,
			unmarshalled: &ohmcjson.MoveCmd{
				FromAccount: "from",
				ToAccount:   "to",
				Amount:      0.5,
				MinConf:     ohmcjson.Int(6),
				Comment:     ohmcjson.String("comment"),
			},
		},
		{
			name: "sendfrom",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("sendfrom", "from", "1Address", 0.5)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSendFromCmd("from", "1Address", 0.5, nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5],"id":1}`,
			unmarshalled: &ohmcjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     ohmcjson.Int(1),
				Comment:     nil,
				CommentTo:   nil,
			},
		},
		{
			name: "sendfrom optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSendFromCmd("from", "1Address", 0.5, ohmcjson.Int(6), nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6],"id":1}`,
			unmarshalled: &ohmcjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     ohmcjson.Int(6),
				Comment:     nil,
				CommentTo:   nil,
			},
		},
		{
			name: "sendfrom optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6, "comment")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSendFromCmd("from", "1Address", 0.5, ohmcjson.Int(6),
					ohmcjson.String("comment"), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6,"comment"],"id":1}`,
			unmarshalled: &ohmcjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     ohmcjson.Int(6),
				Comment:     ohmcjson.String("comment"),
				CommentTo:   nil,
			},
		},
		{
			name: "sendfrom optional3",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("sendfrom", "from", "1Address", 0.5, 6, "comment", "commentto")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSendFromCmd("from", "1Address", 0.5, ohmcjson.Int(6),
					ohmcjson.String("comment"), ohmcjson.String("commentto"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendfrom","params":["from","1Address",0.5,6,"comment","commentto"],"id":1}`,
			unmarshalled: &ohmcjson.SendFromCmd{
				FromAccount: "from",
				ToAddress:   "1Address",
				Amount:      0.5,
				MinConf:     ohmcjson.Int(6),
				Comment:     ohmcjson.String("comment"),
				CommentTo:   ohmcjson.String("commentto"),
			},
		},
		{
			name: "sendmany",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("sendmany", "from", `{"1Address":0.5}`)
			},
			staticCmd: func() interface{} {
				amounts := map[string]float64{"1Address": 0.5}
				return ohmcjson.NewSendManyCmd("from", amounts, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5}],"id":1}`,
			unmarshalled: &ohmcjson.SendManyCmd{
				FromAccount: "from",
				Amounts:     map[string]float64{"1Address": 0.5},
				MinConf:     ohmcjson.Int(1),
				Comment:     nil,
			},
		},
		{
			name: "sendmany optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("sendmany", "from", `{"1Address":0.5}`, 6)
			},
			staticCmd: func() interface{} {
				amounts := map[string]float64{"1Address": 0.5}
				return ohmcjson.NewSendManyCmd("from", amounts, ohmcjson.Int(6), nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5},6],"id":1}`,
			unmarshalled: &ohmcjson.SendManyCmd{
				FromAccount: "from",
				Amounts:     map[string]float64{"1Address": 0.5},
				MinConf:     ohmcjson.Int(6),
				Comment:     nil,
			},
		},
		{
			name: "sendmany optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("sendmany", "from", `{"1Address":0.5}`, 6, "comment")
			},
			staticCmd: func() interface{} {
				amounts := map[string]float64{"1Address": 0.5}
				return ohmcjson.NewSendManyCmd("from", amounts, ohmcjson.Int(6), ohmcjson.String("comment"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendmany","params":["from",{"1Address":0.5},6,"comment"],"id":1}`,
			unmarshalled: &ohmcjson.SendManyCmd{
				FromAccount: "from",
				Amounts:     map[string]float64{"1Address": 0.5},
				MinConf:     ohmcjson.Int(6),
				Comment:     ohmcjson.String("comment"),
			},
		},
		{
			name: "sendtoaddress",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("sendtoaddress", "1Address", 0.5)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSendToAddressCmd("1Address", 0.5, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendtoaddress","params":["1Address",0.5],"id":1}`,
			unmarshalled: &ohmcjson.SendToAddressCmd{
				Address:   "1Address",
				Amount:    0.5,
				Comment:   nil,
				CommentTo: nil,
			},
		},
		{
			name: "sendtoaddress optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("sendtoaddress", "1Address", 0.5, "comment", "commentto")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSendToAddressCmd("1Address", 0.5, ohmcjson.String("comment"),
					ohmcjson.String("commentto"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"sendtoaddress","params":["1Address",0.5,"comment","commentto"],"id":1}`,
			unmarshalled: &ohmcjson.SendToAddressCmd{
				Address:   "1Address",
				Amount:    0.5,
				Comment:   ohmcjson.String("comment"),
				CommentTo: ohmcjson.String("commentto"),
			},
		},
		{
			name: "setaccount",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("setaccount", "1Address", "acct")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSetAccountCmd("1Address", "acct")
			},
			marshalled: `{"jsonrpc":"1.0","method":"setaccount","params":["1Address","acct"],"id":1}`,
			unmarshalled: &ohmcjson.SetAccountCmd{
				Address: "1Address",
				Account: "acct",
			},
		},
		{
			name: "settxfee",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("settxfee", 0.0001)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSetTxFeeCmd(0.0001)
			},
			marshalled: `{"jsonrpc":"1.0","method":"settxfee","params":[0.0001],"id":1}`,
			unmarshalled: &ohmcjson.SetTxFeeCmd{
				Amount: 0.0001,
			},
		},
		{
			name: "signmessage",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("signmessage", "1Address", "message")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSignMessageCmd("1Address", "message")
			},
			marshalled: `{"jsonrpc":"1.0","method":"signmessage","params":["1Address","message"],"id":1}`,
			unmarshalled: &ohmcjson.SignMessageCmd{
				Address: "1Address",
				Message: "message",
			},
		},
		{
			name: "signrawtransaction",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("signrawtransaction", "001122")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewSignRawTransactionCmd("001122", nil, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122"],"id":1}`,
			unmarshalled: &ohmcjson.SignRawTransactionCmd{
				RawTx:    "001122",
				Inputs:   nil,
				PrivKeys: nil,
				Flags:    ohmcjson.String("ALL"),
			},
		},
		{
			name: "signrawtransaction optional1",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("signrawtransaction", "001122", `[{"txid":"123","vout":1,"scriptPubKey":"00","redeemScript":"01"}]`)
			},
			staticCmd: func() interface{} {
				txInputs := []ohmcjson.RawTxInput{
					{
						Txid:         "123",
						Vout:         1,
						ScriptPubKey: "00",
						RedeemScript: "01",
					},
				}

				return ohmcjson.NewSignRawTransactionCmd("001122", &txInputs, nil, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[{"txid":"123","vout":1,"scriptPubKey":"00","redeemScript":"01"}]],"id":1}`,
			unmarshalled: &ohmcjson.SignRawTransactionCmd{
				RawTx: "001122",
				Inputs: &[]ohmcjson.RawTxInput{
					{
						Txid:         "123",
						Vout:         1,
						ScriptPubKey: "00",
						RedeemScript: "01",
					},
				},
				PrivKeys: nil,
				Flags:    ohmcjson.String("ALL"),
			},
		},
		{
			name: "signrawtransaction optional2",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("signrawtransaction", "001122", `[]`, `["abc"]`)
			},
			staticCmd: func() interface{} {
				txInputs := []ohmcjson.RawTxInput{}
				privKeys := []string{"abc"}
				return ohmcjson.NewSignRawTransactionCmd("001122", &txInputs, &privKeys, nil)
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[],["abc"]],"id":1}`,
			unmarshalled: &ohmcjson.SignRawTransactionCmd{
				RawTx:    "001122",
				Inputs:   &[]ohmcjson.RawTxInput{},
				PrivKeys: &[]string{"abc"},
				Flags:    ohmcjson.String("ALL"),
			},
		},
		{
			name: "signrawtransaction optional3",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("signrawtransaction", "001122", `[]`, `[]`, "ALL")
			},
			staticCmd: func() interface{} {
				txInputs := []ohmcjson.RawTxInput{}
				privKeys := []string{}
				return ohmcjson.NewSignRawTransactionCmd("001122", &txInputs, &privKeys,
					ohmcjson.String("ALL"))
			},
			marshalled: `{"jsonrpc":"1.0","method":"signrawtransaction","params":["001122",[],[],"ALL"],"id":1}`,
			unmarshalled: &ohmcjson.SignRawTransactionCmd{
				RawTx:    "001122",
				Inputs:   &[]ohmcjson.RawTxInput{},
				PrivKeys: &[]string{},
				Flags:    ohmcjson.String("ALL"),
			},
		},
		{
			name: "walletlock",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("walletlock")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewWalletLockCmd()
			},
			marshalled:   `{"jsonrpc":"1.0","method":"walletlock","params":[],"id":1}`,
			unmarshalled: &ohmcjson.WalletLockCmd{},
		},
		{
			name: "walletpassphrase",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("walletpassphrase", "pass", 60)
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewWalletPassphraseCmd("pass", 60)
			},
			marshalled: `{"jsonrpc":"1.0","method":"walletpassphrase","params":["pass",60],"id":1}`,
			unmarshalled: &ohmcjson.WalletPassphraseCmd{
				Passphrase: "pass",
				Timeout:    60,
			},
		},
		{
			name: "walletpassphrasechange",
			newCmd: func() (interface{}, error) {
				return ohmcjson.NewCmd("walletpassphrasechange", "old", "new")
			},
			staticCmd: func() interface{} {
				return ohmcjson.NewWalletPassphraseChangeCmd("old", "new")
			},
			marshalled: `{"jsonrpc":"1.0","method":"walletpassphrasechange","params":["old","new"],"id":1}`,
			unmarshalled: &ohmcjson.WalletPassphraseChangeCmd{
				OldPassphrase: "old",
				NewPassphrase: "new",
			},
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Marshal the command as created by the new static command
		// creation function.
		marshalled, err := ohmcjson.MarshalCmd(testID, test.staticCmd())
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			continue
		}

		// Ensure the command is created without error via the generic
		// new command creation function.
		cmd, err := test.newCmd()
		if err != nil {
			t.Errorf("Test #%d (%s) unexpected NewCmd error: %v ",
				i, test.name, err)
		}

		// Marshal the command as created by the generic new command
		// creation function.
		marshalled, err = ohmcjson.MarshalCmd(testID, cmd)
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			continue
		}

		var request ohmcjson.Request
		if err := json.Unmarshal(marshalled, &request); err != nil {
			t.Errorf("Test #%d (%s) unexpected error while "+
				"unmarshalling JSON-RPC request: %v", i,
				test.name, err)
			continue
		}

		cmd, err = ohmcjson.UnmarshalCmd(&request)
		if err != nil {
			t.Errorf("UnmarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !reflect.DeepEqual(cmd, test.unmarshalled) {
			t.Errorf("Test #%d (%s) unexpected unmarshalled command "+
				"- got %s, want %s", i, test.name,
				fmt.Sprintf("(%T) %+[1]v", cmd),
				fmt.Sprintf("(%T) %+[1]v\n", test.unmarshalled))
			continue
		}
	}
}
