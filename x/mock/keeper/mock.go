package keeper

import (
	"context"

	"cosmossdk.io/math"
	// "github.com/likesToEatFish/research-cosmos/x/mock/types"
)

func (k Keeper) AddNewSymbolToBandOracleRequest(ctx context.Context, symbol string, oracleScriptId int64) error {
	_, ok := k.price[symbol]

	if !ok {
		k.SetPricex(ctx, symbol, math.LegacyOneDec())
	}
	return nil
}

func (k Keeper) GetPrice(ctx context.Context, base, quote string) *math.LegacyDec {
	base_price, ok := k.price[base]

	if !ok {
		// panic("call SetPrice " + base)
		k.SetPricex(ctx, base, math.LegacyOneDec())
	}

	quote_price, ok := k.price[quote]

	if !ok {
		// panic("call SetPrice " + quote)
		k.SetPricex(ctx, quote, math.LegacyOneDec())
	}
	base_price, _ = k.price[base]
	quote_price, _ = k.price[quote]
	multiplier := base_price.Quo(quote_price)
	return &multiplier
}

func (k Keeper) SetPricex(ctx context.Context, denom string, price math.LegacyDec) {
	k.price[denom] = price
}
