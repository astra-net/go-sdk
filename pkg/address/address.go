package address

import (
	ethCommon "github.com/ethereum/go-ethereum/common"
)

const (
	// HashLength is the expected length of the hash
	HashLength = ethCommon.HashLength
	// AddressLength is the expected length of the address
	AddressLength = ethCommon.AddressLength
)

type T = ethCommon.Address

// ParseAddr parses the given address
func Parse(s string) T {
	// The result can be 0x00...00 if the passing param is not a correct address.
	return ethCommon.HexToAddress(s)
}
