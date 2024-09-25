package keeper

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

func (k *Keeper) CreateNewVault(
	ctx context.Context,
	denom string,
	owner sdk.AccAddress,
	collateral sdk.Coin,
	mint sdk.Coin,
) error {
	vm, err := k.GetVaultManager(ctx, denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", denom)
	}

	params := k.GetParams(ctx)
	vmParams := vm.Params

	// Check if expect min less than MinInitialDebt
	if mint.Amount.LT(params.MinInitialDebt) {
		return fmt.Errorf("initial mint should be greater than min. Got %v, expected %v", mint, params.MinInitialDebt)
	}

	// Calculate collateral ratio
	price := k.oracleKeeper.GetPrice(ctx, denom)
	// TODO: recalculate with denom decimal?
	collateralValue := math.LegacyNewDecFromInt(collateral.Amount).Mul(price)
	ratio := collateralValue.QuoInt(mint.Amount)

	if ratio.LT(vmParams.MinCollateralRatio) {
		return fmt.Errorf("collateral ratio invalid, got %d, min %d", ratio, vmParams.MinCollateralRatio)
	}

	feeAmount := math.LegacyNewDecFromInt(mint.Amount).Mul(params.MintingFee).TruncateInt()
	feeCoins := sdk.NewCoins(sdk.NewCoin(mint.Denom, feeAmount))
	mintedCoins := feeCoins.Add(mint)

	vaultId, vaultAddress := k.GetVaultIdAndAddress(ctx)

	// Lock collateral asset
	err = k.bankKeeper.SendCoins(ctx, owner, vaultAddress, sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	// Mint and transfer to user and reserve
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, mintedCoins)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, feeCoins)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, sdk.NewCoins(mint))
	if err != nil {
		return err
	}

	// Set vault
	vault := types.Vault{
		Id:               vaultId,
		Owner:            owner.String(),
		Debt:             mintedCoins[0],
		CollateralLocked: collateral,
		Status:           types.ACTIVE,
	}
	err = k.SetVault(ctx, vault)
	if err != nil {
		return err
	}
	// Update vault manager
	vm.MintAvailable = vm.MintAvailable.Sub(mintedCoins[0].Amount)
	return k.VaultsManager.Set(ctx, denom, vm)
}

func (k *Keeper) MintCoin(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	mint sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}
	vm, err := k.GetVaultManager(ctx, vault.CollateralLocked.Denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", vault.CollateralLocked.Denom)
	}

	params := k.GetParams(ctx)

	lockedCoin := vault.CollateralLocked
	price := k.oracleKeeper.GetPrice(ctx, lockedCoin.Denom)
	lockedValue := math.LegacyNewDecFromInt(lockedCoin.Amount).Mul(price)

	feeAmount := math.LegacyNewDecFromInt(mint.Amount).Mul(params.MintingFee).TruncateInt()
	feeCoins := sdk.NewCoins(sdk.NewCoin(mint.Denom, feeAmount))
	mintedAmount := feeAmount.Add(mint.Amount)
	mintedCoins := feeCoins.Add(mint)

	// calculate ratio
	ratio := lockedValue.Quo(math.LegacyNewDecFromInt(vault.Debt.Amount.Add(mintedAmount)))
	if ratio.LT(vm.Params.MinCollateralRatio) {
		return fmt.Errorf("collateral ratio invalid, got %d, min %d", ratio, vm.Params.MinCollateralRatio)
	}

	// Mint and transfer to user and reserve
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, mintedCoins)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, feeCoins)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(vault.Owner), sdk.NewCoins(mint))
	if err != nil {
		return err
	}

	// Update vault debt
	vault.Debt = vault.Debt.Add(sdk.NewCoin(vault.Debt.Denom, mintedAmount))
	err = k.SetVault(ctx, vault)
	if err != nil {
		return err
	}

	// Update vault manager
	vm.MintAvailable = vm.MintAvailable.Sub(mintedCoins[0].Amount)
	return k.VaultsManager.Set(ctx, vault.CollateralLocked.Denom, vm)

}

func (k *Keeper) RepayDebt(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	mint sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}
	vm, err := k.GetVaultManager(ctx, vault.CollateralLocked.Denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", vault.CollateralLocked.Denom)
	}

	burnAmount := mint
	if vault.Debt.IsLT(burnAmount) {
		burnAmount = vault.Debt
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(burnAmount))
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(burnAmount))
	if err != nil {
		return err
	}

	// Update vault debt
	vault.Debt = vault.Debt.Sub(burnAmount)
	err = k.SetVault(ctx, vault)
	if err != nil {
		return err
	}

	vm.MintAvailable = vm.MintAvailable.Add(burnAmount.Amount)
	return k.VaultsManager.Set(ctx, vm.Denom, vm)
}

func (k *Keeper) DepositToVault(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	collateral sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}

	// Lock collateral asset
	err = k.bankKeeper.SendCoins(ctx, sender, sdk.MustAccAddressFromBech32(vault.Address), sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	// Update vault
	vault.CollateralLocked = vault.CollateralLocked.Add(collateral)
	return k.SetVault(ctx, vault)
}

func (k *Keeper) WithdrawFromVault(
	ctx context.Context,
	vaultId uint64,
	sender sdk.AccAddress,
	collateral sdk.Coin,
) error {
	vault, err := k.GetVault(ctx, vaultId)
	if err != nil {
		return err
	}

	if vault.CollateralLocked.Amount.LT(collateral.Amount) {
		return fmt.Errorf("%d exeed locked amount: %d", collateral.Amount, vault.CollateralLocked.Amount)
	}

	vm, err := k.GetVaultManager(ctx, vault.CollateralLocked.Denom)
	if err != nil {
		return fmt.Errorf("%s was not actived", vault.CollateralLocked.Denom)
	}

	newLock := vault.CollateralLocked.Sub(collateral)
	price := k.oracleKeeper.GetPrice(ctx, collateral.Denom)
	newLockValue := math.LegacyNewDecFromInt(newLock.Amount).Mul(price)
	ratio := newLockValue.Quo(math.LegacyNewDecFromInt(vault.Debt.Amount))

	if ratio.LT(vm.Params.MinCollateralRatio) {
		return fmt.Errorf("ratio less than min ratio. Got: %d, min: %d", ratio, vm.Params.MinCollateralRatio)
	}

	err = k.bankKeeper.SendCoins(ctx, sdk.MustAccAddressFromBech32(vault.Address), sender, sdk.NewCoins(collateral))
	if err != nil {
		return err
	}

	// Update vault
	vault.CollateralLocked = vault.CollateralLocked.Sub(collateral)
	return k.SetVault(ctx, vault)
}

func (k *Keeper) UpdateVaultsDebt(
	ctx context.Context,
) error {
	params := k.GetParams(ctx)
	fee := params.StabilityFee

	return k.Vaults.Walk(ctx, nil, func(id uint64, vault types.Vault) (bool, error) {
		var err error
		if vault.Status == 0 {
			debtAmount := vault.Debt.Amount
			newDebtAmount := math.LegacyNewDecFromInt(debtAmount).Add(math.LegacyNewDecFromInt(debtAmount).Mul(fee)).TruncateInt()
			vault.Debt.Amount = newDebtAmount
			err = k.Vaults.Set(ctx, id, vault)
		}

		return false, err
	})
}

func (k *Keeper) ShouldLiquidate(
	ctx context.Context,
	vault types.Vault,
	price math.LegacyDec,
	liquidationRatio math.LegacyDec,
) (bool, error) {
	// Only liquidate OPEN vault
	if vault.Status != 0 {
		return false, nil
	}

	collateralValue := math.LegacyNewDecFromInt(vault.CollateralLocked.Amount).Mul(price)
	ratio := collateralValue.Quo(math.LegacyNewDecFromInt(vault.Debt.Amount))

	if ratio.LTE(liquidationRatio) {
		return true, nil
	}
	return false, nil
}

func (k *Keeper) GetLiquidations(
	ctx context.Context,
) ([]*types.Liquidation, error) {

	liquidationRatios := make(map[string]math.LegacyDec)
	prices := make(map[string]math.LegacyDec)
	liquidations := make(map[string]*types.Liquidation)

	err := k.VaultsManager.Walk(ctx, nil, func(key string, vm types.VaultMamager) (bool, error) {
		price := k.oracleKeeper.GetPrice(ctx, vm.Denom)
		prices[vm.Denom] = price
		liquidationRatios[vm.Denom] = vm.Params.LiquidationRatio
		liquidations[vm.Denom] = types.NewEmptyLiquidation(vm.Denom)

		return false, nil
	})
	if err != nil {
		return nil, err
	}

	err = k.Vaults.Walk(ctx, nil, func(id uint64, vault types.Vault) (bool, error) {
		denom := vault.CollateralLocked.Denom
		shouldLiquidate, err := k.ShouldLiquidate(ctx, vault, prices[denom], liquidationRatios[denom])
		if shouldLiquidate && err == nil {
			liquidations[denom].LiquidatingVaults = append(liquidations[denom].LiquidatingVaults, &vault)
			liquidations[denom].VaultLiquidationStatus[id] = &types.VaultLiquidationStatus{}

			vault.Status = types.LIQUIDATING
			vault.LiquidationPrice = prices[denom]
			err := k.SetVault(ctx, vault)
			if err != nil {
				return true, err
			}
		}

		return false, nil
	})
	if err != nil {
		return nil, err
	}

	var result []*types.Liquidation
	for _, liquidation := range liquidations {
		if len(liquidation.LiquidatingVaults) != 0 {
			result = append(result, liquidation)
		}
	}

	return result, nil
}

// TODO: Separate this func
func (k *Keeper) Liquidate(
	ctx context.Context,
	liquidation types.Liquidation,
) error {
	params := k.GetParams(ctx)

	// Get total sold amount & collateral asset remain
	var (
		totalDebt, sold, totalCollateralRemain sdk.Coin
	)

	for _, vault := range liquidation.LiquidatingVaults {
		totalDebt = totalDebt.Add(vault.Debt)
		// transfer all remain collateral locked in vault to vaults module for distributing.
		vaultAddr := sdk.MustAccAddressFromBech32(vault.Address)
		balances := k.bankKeeper.GetAllBalances(ctx, vaultAddr)
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, vaultAddr, types.ModuleName, balances)
		if err != nil {
			return err
		}
	}

	for _, status := range liquidation.VaultLiquidationStatus {
		sold = sold.Add(status.Sold)
		totalCollateralRemain = totalCollateralRemain.Add(status.RemainCollateral)
	}

	// Sold amount enough to cover debt
	if sold.Amount.GTE(totalDebt.Amount) {
		// Burn debt
		err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(totalDebt))
		if err != nil {
			return err
		}

		// If remain sold, send to reserve
		remain := sold.Sub(totalDebt)
		if remain.Amount.GT(math.ZeroInt()) {
			err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(remain))
			if err != nil {
				return err
			}
		}

		// Take the liquidation penalty and send back to vault owner
		if totalCollateralRemain.Amount.GT(math.ZeroInt()) {
			//TODO: decimal

			for _, vault := range liquidation.LiquidatingVaults {
				collateralRemain := liquidation.VaultLiquidationStatus[vault.Id].RemainCollateral
				if collateralRemain.Amount.Equal(math.ZeroInt()) {
					continue
				}
				penaltyAmount := math.LegacyNewDecFromInt(vault.Debt.Amount).Quo(vault.LiquidationPrice).Mul(params.LiquidationPenalty).TruncateInt()
				if penaltyAmount.GTE(collateralRemain.Amount) {
					err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(collateralRemain))
					if err != nil {
						return err
					}
				} else {
					err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(sdk.NewCoin(collateralRemain.Denom, penaltyAmount)))
					if err != nil {
						return err
					}
					err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(vault.Owner), sdk.NewCoins(sdk.NewCoin(collateralRemain.Denom, collateralRemain.Amount.Sub(penaltyAmount))))
					if err != nil {
						return err
					}
				}
			}
		}
	} else {
		// does not raise enough to cover nomUSD debt

		// Burn sold amount
		err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sold))
		if err != nil {
			return err
		}

		// No collateral remain
		if totalCollateralRemain.Amount.Equal(math.ZeroInt()) {
			//TODO: send shortfall to reserve
			return nil
		} else {
			// If there some collateral asset remain, try to reconstitue vault
			// Priority by collateral ratio at momment
			// So that mean we need less resource for high ratio vault

			ratios := make([]math.LegacyDec, 0)
			//TODO: Sort by CR in GetLiquidations could reduce calculate here
			for _, vault := range liquidation.LiquidatingVaults {
				penaltyAmount := math.LegacyNewDecFromInt(vault.Debt.Amount).Quo(vault.LiquidationPrice).Mul(params.LiquidationPenalty).TruncateInt()
				err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(sdk.NewCoin(liquidation.Denom, penaltyAmount)))
				if err != nil {
					return err
				}
				vault.CollateralLocked.Amount = vault.CollateralLocked.Amount.Sub(penaltyAmount)
				totalCollateralRemain.Amount = totalCollateralRemain.Amount.Sub(penaltyAmount)

				ratio := math.LegacyNewDecFromInt(vault.CollateralLocked.Amount).Mul(vault.LiquidationPrice).Quo(math.LegacyNewDecFromInt(vault.Debt.Amount))
				ratios = append(ratios, ratio)
			}

			// Sort the vaults by CR in descending order
			sort.Slice(liquidation.LiquidatingVaults, func(i, j int) bool {
				return ratios[i].GT(ratios[j])
			})

			// Try to reconstitue vaults
			totalRemainDebt := totalDebt.Sub(sold)
			for _, vault := range liquidation.LiquidatingVaults {
				// if remain debt & collateral can cover full vault
				// open again
				if vault.Debt.IsLTE(totalRemainDebt) && vault.CollateralLocked.IsLTE(totalCollateralRemain) {
					// Lock collateral to vault address
					err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(vault.Address), sdk.NewCoins(vault.CollateralLocked))
					if err != nil {
						return err
					}
					totalRemainDebt = totalRemainDebt.Sub(vault.Debt)
					totalCollateralRemain = totalCollateralRemain.Sub(vault.CollateralLocked)

					vault.Status = types.ACTIVE
					err = k.SetVault(ctx, *vault)
					if err != nil {
						return err
					}
				} else {
					vault.Status = types.LIQUIDATED
					err := k.SetVault(ctx, *vault)
					if err != nil {
						return err
					}
				}
			}

			// if remain collateral, send to reserve
			if totalCollateralRemain.Amount.GT(math.ZeroInt()) {
				err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.ReserveModuleName, sdk.NewCoins(totalCollateralRemain))
				if err != nil {
					return err
				}
			}

			// if remain debt, send shortfall
			// TODO: Shortfall

		}
	}
	return nil
}

func (k *Keeper) GetVault(
	ctx context.Context,
	id uint64,
) (types.Vault, error) {
	vault, err := k.Vaults.Get(ctx, id)
	if err != nil {
		return types.Vault{}, err
	}
	return vault, nil
}

func (k *Keeper) SetVault(
	ctx context.Context,
	vault types.Vault,
) error {
	return k.Vaults.Set(ctx, vault.Id, vault)
}

func (k *Keeper) GetVaultIdAndAddress(
	ctx context.Context,
) (uint64, sdk.AccAddress) {
	id, err := k.VaultsSequence.Next(ctx)
	if err != nil {
		return 0, sdk.AccAddress{}
	}
	address := address.Module(types.ModuleName, []byte(strconv.Itoa(int(id))))

	return id, address
}
