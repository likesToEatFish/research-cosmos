package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// ScheduleForkUpgrade executes any necessary fork logic for based upon the current
// block height and chain ID (mainnet or testnet). It sets an upgrade plan once
// the chain reaches the pre-defined upgrade height.
//
// CONTRACT: for this logic to work properly it is required to:
//
//  1. Release a non-breaking patch version so that the chain can set the scheduled upgrade plan at upgrade-height.
//  2. Release the software defined in the upgrade-info
func (app *RealioNetwork) ScheduleForkUpgrade(ctx sdk.Context) {
	if ctx.BlockHeight() == 5989487 {
		// remove duplicate UnbondingQueueKey
		removeDuplicateValueUnbondingQueueKey(app, ctx)
	}
	// NOTE: there are no testnet forks for the existing versions
	// if !types.IsMainnet(ctx.ChainID()) {
	//	return
	//}
	//
	// upgradePlan := upgradetypes.Plan{
	//	Height: ctx.BlockHeight(),
	//}
	//
	//// handle mainnet forks with their corresponding upgrade name and info
	// switch ctx.BlockHeight() {
	// case v2.MainnetUpgradeHeight:
	//	upgradePlan.Name = v2.UpgradeName
	//	upgradePlan.Info = v2.UpgradeInfo
	//default:
	//	// No-op
	//	return
	//}
	//
	//// schedule the upgrade plan to the current block height, effectively performing
	//// a hard fork that uses the upgrade handler to manage the migration.
	// if err := app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradePlan); err != nil {
	//	panic(
	//		fmt.Errorf(
	//			"failed to schedule upgrade %s during BeginBlock at height %d: %w",
	//			upgradePlan.Name, ctx.BlockHeight(), err,
	//		),
	//	)
	//}
}

func removeDuplicateValueUnbondingQueueKey(app *RealioNetwork, ctx sdk.Context) {
	// Get Staking keeper, codec and staking store
	sk := app.StakingKeeper
	cdc := app.AppCodec()
	store := ctx.KVStore(app.keys[stakingtypes.ModuleName])

	// remove duplicate UnbondingQueueKey
	ubdTime := sk.UnbondingTime(ctx)
	currTime := ctx.BlockTime()

	unbondingTimesliceIterator := sk.UBDQueueIterator(ctx, currTime.Add(ubdTime)) // make sure to iterate all queue
	defer unbondingTimesliceIterator.Close()

	for ; unbondingTimesliceIterator.Valid(); unbondingTimesliceIterator.Next() {
		timeslice := stakingtypes.DVPairs{}
		value := unbondingTimesliceIterator.Value()
		cdc.MustUnmarshal(value, &timeslice)

		dvPairs := removeDuplicates(timeslice.Pairs)
		bz := cdc.MustMarshal(&stakingtypes.DVPairs{Pairs: dvPairs})

		store.Set(unbondingTimesliceIterator.Key(), bz)
	}
}

func removeDuplicates(dvPairs []stakingtypes.DVPair) []stakingtypes.DVPair {
	var list []stakingtypes.DVPair
	for _, item := range dvPairs {
		if contains(list, item) == false {
			list = append(list, item)
		}
	}
	return list
}

func contains(s []stakingtypes.DVPair, e stakingtypes.DVPair) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
