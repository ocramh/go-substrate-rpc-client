package types

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
)

// ContractCallRequest represents a request sent to a smart contract
type ContractCallRequest struct {
	Origin              string `json:"origin"`   // the caller account ID
	Dest                string `json:"dest"`     // the account ID of the contract to invoke
	Value               U64    `json:"value"`    // the value to be transferred as part of the call
	GasLimit            U64    `json:"gasLimit"` // maximum amount of gas to be used
	StorageDepositLimit *U64   `json:"storageDepositLimit"`
	InputData           string `json:"inputData"`
}

// ContractExecResult is the response received when querying a smart contract
// storage via RPC.
type ContractExecResult struct {
	GasConsumed U64
	GasRequired U64
	StorageDeposit
	DebugMessage string
	Result       ContractExecResultResult
}

type StorageDeposit struct{}

// ContractExecResultResult holds the storage data returned by the contract.
// If the call is successful the OK field is populated with hex encoded data
// that needs further decoding depending on the response event. If the call is
// not successful the Err field is populated with an error message.
type ContractExecResultResult struct {
	Ok   ExecPayload
	Err  ExecPayload
	Type string // possible values are 'Ok' | 'Err';
}

// ExecPayload parses the contract storage response. The Data field
// contains the hex encoded payload - which is further encoded with the parity
// SCALE codec.
type ExecPayload struct {
	Flags interface{} `json:"flags"`
	Data  string      `json:"data"`
}

//
type Item struct {
	Cid    string
	Name   string
	Owners []Owner
	Price  U64
}

type T struct {
	Cid    string
	Name   string
	Owners []Owner
	Price  U64
}

type Owner struct {
	Account   AccountID
	Name      string
	Ownership U32
}

type F struct {
	Account   AccountID
	Name      string
	Ownership U32
}

func (o *Owner) Decode(decoder scale.Decoder) error {
	var f F
	err := decoder.Decode(&f)
	if err != nil {
		fmt.Println(err)
		return err
	}

	o.Name = f.Name
	o.Account = f.Account
	o.Ownership = f.Ownership

	return nil
}

func (i *Item) Decode(decoder scale.Decoder) error {
	var t T
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(err)
		return err
	}

	i.Cid = t.Cid
	i.Name = t.Name
	i.Owners = t.Owners
	i.Price = t.Price

	return nil
}

type ContractExecResultErr struct {
	Flags interface{}
	Data  Bytes
}

type ContractReturnFlags struct {
	IsRevert bool
}
