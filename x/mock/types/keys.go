package types

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/math"
)

const (
	// ModuleName defines the module name
	ModuleName = "mock"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mock"
)

// ParamsKey is the prefix to retrieve all Params
var ParamsKey = collections.NewPrefix("p_mock")

func NewSetPrice(denom, authority string, price math.LegacyDec) MsgSetPrice {
	return MsgSetPrice{
		Denom:     denom,
		Price:     price,
		Authority: authority,
	}
}

func NewQueryPrice(denom, denom2 string) QueryPriceRequest {
	return QueryPriceRequest{
		Denom1: denom,
		Denom2: denom2,
	}
}
