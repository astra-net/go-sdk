package cmd

import (
	"github.com/astra-net/go-sdk/pkg/common"
	"github.com/astra-net/go-sdk/pkg/validation"
)

type Address struct {
	address string
}

func (address Address) String() string {
	return address.address
}

func (address *Address) Set(s string) error {
	err := validation.ValidateAddress(s)
	if err != nil {
		return err
	}
	address.address = s
	return nil
}

func (address Address) Type() string {
	return "address"
}

type chainIDWrapper struct {
	chainID *common.ChainID
}

func (chainIDWrapper chainIDWrapper) String() string {
	return chainIDWrapper.chainID.Name
}

func (chainIDWrapper *chainIDWrapper) Set(s string) error {
	chain, err := common.StringToChainID(s)
	chainIDWrapper.chainID = chain
	if err != nil {
		return err
	}
	return nil
}

func (chainIDWrapper chainIDWrapper) Type() string {
	return "chain-id"
}
