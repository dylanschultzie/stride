package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	recordstypes "github.com/Stride-Labs/stride/v9/x/records/types"
	"github.com/Stride-Labs/stride/v9/x/stakeibc/types"
)

// Emits a successful liquid stake event, and displays metadata such as the stToken amount
func EmitSuccessfulLiquidStakeEvent(ctx sdk.Context, msg *types.MsgLiquidStake, hostZone types.HostZone, stAmount sdkmath.Int) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLiquidStakeRequest,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyLiquidStaker, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyHostZone, hostZone.ChainId),
			sdk.NewAttribute(types.AttributeKeyNativeBaseDenom, msg.HostDenom),
			sdk.NewAttribute(types.AttributeKeyNativeIBCDenom, hostZone.IbcDenom),
			sdk.NewAttribute(types.AttributeKeyNativeAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyStTokenAmount, stAmount.String()),
		),
	)
}

// Emits a successful LSM liquid stake event, and displays metadata such as the stToken amount
func EmitSuccessfulLSMLiquidStakeEvent(ctx sdk.Context, hostZone types.HostZone, lsmTokenDeposit recordstypes.LSMTokenDeposit) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLSMLiquidStakeRequest,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyLiquidStaker, lsmTokenDeposit.StakerAddress),
			sdk.NewAttribute(types.AttributeKeyHostZone, hostZone.ChainId),
			sdk.NewAttribute(types.AttributeKeyNativeBaseDenom, hostZone.HostDenom),
			sdk.NewAttribute(types.AttributeKeyValidator, lsmTokenDeposit.ValidatorAddress),
			sdk.NewAttribute(types.AttributeKeyNativeIBCDenom, lsmTokenDeposit.IbcDenom),
			sdk.NewAttribute(types.AttributeKeyLSMTokenBaseDenom, lsmTokenDeposit.Denom),
			sdk.NewAttribute(types.AttributeKeyNativeAmount, lsmTokenDeposit.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyStTokenAmount, lsmTokenDeposit.StToken.Amount.String()),
		),
	)
}

// Emits a failed LSM liquid stake event, and displays the error
func EmitFailedLSMLiquidStakeEvent(ctx sdk.Context, hostZone types.HostZone, lsmTokenDeposit recordstypes.LSMTokenDeposit, errorMessage string) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLSMLiquidStakeFailed,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyLiquidStaker, lsmTokenDeposit.StakerAddress),
			sdk.NewAttribute(types.AttributeKeyHostZone, hostZone.ChainId),
			sdk.NewAttribute(types.AttributeKeyValidator, lsmTokenDeposit.ValidatorAddress),
			sdk.NewAttribute(types.AttributeKeyNativeBaseDenom, hostZone.HostDenom),
			sdk.NewAttribute(types.AttributeKeyNativeIBCDenom, lsmTokenDeposit.IbcDenom),
			sdk.NewAttribute(types.AttributeKeyLSMTokenBaseDenom, lsmTokenDeposit.Denom),
			sdk.NewAttribute(types.AttributeKeyNativeAmount, lsmTokenDeposit.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyStTokenAmount, lsmTokenDeposit.StToken.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyError, errorMessage),
		),
	)
}

// Emits an event if a validator's exchange rate changed
func EmitValidatorExchangeRateChangeEvent(
	ctx sdk.Context,
	chainId string,
	validatorAddress string,
	previousExchangeRate,
	currentExchangeRate sdk.Dec,
) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeValidatorExchangeRate,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyHostZone, chainId),
			sdk.NewAttribute(types.AttributeKeyValidator, validatorAddress),
			sdk.NewAttribute(types.AttributeKeyPreviousExchangeRate, previousExchangeRate.String()),
			sdk.NewAttribute(types.AttributeKeyCurrentExchangeRate, currentExchangeRate.String()),
		),
	)
}

// Emits an event if a validator was slashed
func EmitValidatorSlashEvent(
	ctx sdk.Context,
	hostZone types.HostZone,
	validatorAddress string,
	slashPercent sdk.Dec,
	slashAmount sdkmath.Int,
	currentDelegation sdkmath.Int,
) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeValidatorSlash,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyHostZone, hostZone.ChainId),
			sdk.NewAttribute(types.AttributeKeyNativeBaseDenom, hostZone.HostDenom),
			sdk.NewAttribute(types.AttributeKeyValidator, validatorAddress),
			sdk.NewAttribute(types.AttributeKeySlashPercent, slashPercent.String()),
			sdk.NewAttribute(types.AttributeKeySlashAmount, slashAmount.String()),
			sdk.NewAttribute(types.AttributeKeyCurrentDelegation, currentDelegation.String()),
		),
	)
}