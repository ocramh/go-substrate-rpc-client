package contracts

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// Contract is an interface which defines functionalities used for interacting
// with smart contracts.
type Contract interface {
	Call(result interface{}, r types.ContractCallRequest) error
}

type contract struct {
	client client.Client
}

func NewContract(c client.Client) Contract {
	return contract{client: c}
}

// Call executes a contract request and marshals the response into result.
// The result argument must be a pointer.
func (c contract) Call(result interface{}, callRequest types.ContractCallRequest) error {
	r := types.ContractExecResult{}

	err := c.client.Call(&r, "contracts_call", callRequest)
	if err != nil {
		return err
	}

	fmt.Println("call ok", r.Result.Ok.Flags)
	fmt.Println("call err", r.Result.Err.Flags)

	return types.DecodeFromHex(r.Result.Ok.Data, result)
}

type Message struct {
	Name     string `json:"label"`
	Payable  bool   `json:"payable"`
	Mutate   bool   `json:"mutate"`
	Selector string `json:"selector"`
}

// TODO
// set up Contract Transcode
// 1. read generated contract metadata.json
// 2. expose metadata messages names and args
