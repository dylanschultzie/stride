package keeper

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/gogoproto/proto"

	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	"github.com/Stride-Labs/stride/v16/utils"
	epochstypes "github.com/Stride-Labs/stride/v16/x/epochs/types"
	icqtypes "github.com/Stride-Labs/stride/v16/x/interchainquery/types"
	"github.com/Stride-Labs/stride/v16/x/stakeibc/types"
)

// The goal of this code is to allow certain reward token types to be automatically traded into other types
// This happens before the rest of the staking, allocation, distribution etc. would continue as normal

// Reward tokens are any special denoms which are paid out in the withdrawal address
// Most host zones inflate their tokens and the newly minted host denom is what appears in the withdrawal ICA
// This code allows for chains to use foreign denoms as revenue, which can be traded to any other denom first

// 0. Before the normal epochly base denom withdrawal address check in the normal staking flow
//		by doing this check before the normal staking flow, we could trade tokens which are stake-able instead
// 1. Epochly check the reward denom balance in the withdrawal address
//       on callback, send all this reward denom from withdrawl ICA to trade ICA on the trade zone (OSMOSIS)
// 2. Epochly check the reward denom balance in trade ICA
//		on callback, trade all reward denom for target output denom defined by pool and routes in params
// 3. Epochly check the target denom balance in trade ICA
//		on callback, transfer these target denoms from trade ICA to withdrawal ICA on original host zone
// Normal staking flow continues from there. So if the target denom is the original host zone base denom
// as will often be the case, then these tokens will follow the normal staking and distribution flow.

// msgs with packet forwarding memos can unwind through the reward zone and chain two transfer hops without callbacks

// ICA tx will kick off transfering the reward tokens from the hostZone withdrawl ICA to the tradeZone trade ICA
// This will be two hops to unwind the ibc denom through the rewardZone using pfm in the transfer memo if possible
func (k Keeper) TransferRewardTokensHostToTrade(ctx sdk.Context, amount sdk.Int, route types.TradeRoute) error {
	// Timeout for ica tx and the transfer msgs is at end of epoch
	strideEpochTracker, found := k.GetEpochTracker(ctx, epochstypes.STRIDE_EPOCH)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, epochstypes.STRIDE_EPOCH)
	}
	timeout := uint64(strideEpochTracker.NextEpochStartTime)

	startingDenom := route.RewardDenomOnHostZone
	sendTokens := sdk.NewCoin(startingDenom, amount)

	withdrawlIcaAddress := route.HostToRewardHop.FromAccount.Address
	unwindIcaAddress := route.HostToRewardHop.ToAccount.Address
	tradeIcaAddress := route.RewardToTradeHop.ToAccount.Address

	// Import and use pfm data structures instead of this JSON if we can determine consistent module version...
	memoJSON := fmt.Sprintf(`"forward": {"receiver": "%s","port": "%s","channel":"%s","timeout":"10m","retries": 2}`,
		tradeIcaAddress, transfertypes.PortID, route.RewardToTradeHop.TransferChannelId)
	// This transfer channel id is a channel on the reward Zone for transfers to the trade zone
	// so remember, this transfer channel id will not exist on hostZone or stride

	var msgs []proto.Message
	msgs = append(msgs, &transfertypes.MsgTransfer{
		SourcePort:       transfertypes.PortID,
		SourceChannel:    route.HostToRewardHop.TransferChannelId, // channel on hostZone for transfers to rewardZone
		Token:            sendTokens,
		Sender:           withdrawlIcaAddress,
		Receiver:         unwindIcaAddress, // could be "pfm" or a real address depending on version
		TimeoutTimestamp: timeout,
		Memo:             memoJSON,
	})

	hostZoneId := route.HostToRewardHop.FromAccount.ChainId
	rewardZoneId := route.HostToRewardHop.ToAccount.ChainId
	tradeZoneId := route.RewardToTradeHop.ToAccount.ChainId
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZoneId,
		"Preparing MsgTransfer of %+v from %s to %s to %s", sendTokens, hostZoneId, rewardZoneId, tradeZoneId))

	// Send the ICA tx to kick off transfer from hostZone through rewardZone to the tradeZone (no callbacks)
	hostZoneConnectionId := route.HostToRewardHop.FromAccount.ConnectionId
	err := k.SubmitICATxWithoutCallback(ctx, hostZoneConnectionId, types.ICAAccountType_WITHDRAWAL, msgs, timeout)
	if err != nil {
		return errorsmod.Wrapf(types.ErrICATxFailed, "Failed to submit ICA tx, Messages: %+v, err: %s", msgs, err.Error())
	}

	return nil
}

// ICA tx to kick off transfering the converted tokens back from tradeZone to the hostZone withdrawal ICA
func (k Keeper) TransferConvertedTokensTradeToHost(ctx sdk.Context, amount sdk.Int, route types.TradeRoute) error {
	// Timeout for ica tx and the transfer msgs is at end of epoch
	strideEpochTracker, found := k.GetEpochTracker(ctx, epochstypes.STRIDE_EPOCH)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, epochstypes.STRIDE_EPOCH)
	}
	timeout := uint64(strideEpochTracker.NextEpochStartTime)

	convertedDenom := route.TargetDenomOnTradeZone
	sendTokens := sdk.NewCoin(convertedDenom, amount)

	tradeIcaAddress := route.TradeToHostHop.FromAccount.Address
	withdrawlIcaAddress := route.TradeToHostHop.ToAccount.Address

	var msgs []proto.Message
	msgs = append(msgs, &transfertypes.MsgTransfer{
		SourcePort:       transfertypes.PortID,
		SourceChannel:    route.TradeToHostHop.TransferChannelId, // channel on tradeZone for transfers to hostZone
		Token:            sendTokens,
		Sender:           tradeIcaAddress,
		Receiver:         withdrawlIcaAddress,
		TimeoutTimestamp: timeout,
		Memo:             "",
	})

	hostZoneId := route.TradeToHostHop.ToAccount.ChainId
	tradeZoneId := route.TradeToHostHop.FromAccount.ChainId
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZoneId,
		"Preparing MsgTransfer of %+v from %s to %s", sendTokens, tradeZoneId, hostZoneId))

	// Send the ICA tx to kick off transfer from hostZone through rewardZone to the tradeZone (no callbacks)
	tradeZoneConnectionId := route.TradeToHostHop.FromAccount.ConnectionId
	err := k.SubmitICATxWithoutCallback(ctx, tradeZoneConnectionId, types.ICAAccountType_TRADE, msgs, timeout)
	if err != nil {
		return errorsmod.Wrapf(types.ErrICATxFailed, "Failed to submit ICA tx, Messages: %+v, err: %s", msgs, err.Error())
	}

	return nil
}

// Trade reward tokens in the Trade ICA for the target output token type using ICA remote tx on trade zone
// The amount represents the total amount of the reward token in the trade ICA found by the calling ICQ
// Depending on min and max swap amounts set in the route, it is possible not the full amount given will swap
func (k Keeper) TradeRewardTokens(ctx sdk.Context, amount sdk.Int, route types.TradeRoute) error {
	// If the min swap amount was not set it would be ZeroInt, if positive we need to compare to the amount given
	//  then if the min swap amount is greater than the current amount, do nothing this epoch to avoid small swaps
	if route.MinSwapAmount.IsPositive() && route.MinSwapAmount.GT(amount) {
		return nil
	}

	// If the max swap amount was not set it would be ZeroInt, if positive we need to compare to the amount given
	//  then if max swap amount is LTE to amount full swap is possible so amount is fine, otherwise set amount to max
	if route.MaxSwapAmount.IsPositive() && route.MaxSwapAmount.GT(amount) {
		amount = route.MaxSwapAmount
	}

	// See if pool swap spot price has been set to a valid ratio (string representing a float like "10.203")
	// If there is a valid spot price, use it to set a floor for the acceptable minimum output tokens
	// 5% slippage is allowed so multiply by 0.95 the expected spot price to get the target minimum
	// minOut is the minimum number of route.TargetDenomOnTradeZone we must receive or the swap will fail
	minOut := sdk.ZeroInt()
	if route.SpotPrice != "" {
		if spotPrice, err := sdk.NewDecFromStr(route.SpotPrice); err == nil {
			slippageRatio := sdk.NewDecWithPrec(95, 2) // 0.95 to allow 5% loss to slippage
			inputToOutputRatio := slippageRatio.Mul(spotPrice)
			minOut = sdk.NewDecFromInt(amount).Mul(inputToOutputRatio).TruncateInt()
		} else {
			k.Logger(ctx).Error("Couldn't parse the spot price %s as a Decimal", route.SpotPrice)
			// Don't allow a missing spot price to stop the swap, use zeroInt as the lower bound for now
		}
	}

	tradeIcaAccount := route.RewardToTradeHop.ToAccount
	tradeTokens := sdk.NewCoin(route.RewardDenomOnTradeZone, amount)

	// Prepare Osmosis GAMM module MsgSwapExactAmountIn from the trade account to perform the trade
	// If we want to generalize in the future, write swap message generation funcs for each DEX type,
	// decide which msg generation function to call based on check of which tradeZone was passed in
	var msgs []proto.Message
	if amount.GT(sdk.ZeroInt()) {
		var routes []types.SwapAmountInRoute
		routes = append(routes, types.SwapAmountInRoute{
			PoolId:        route.PoolId,
			TokenOutDenom: route.TargetDenomOnTradeZone,
		})
		msgs = append(msgs, &types.MsgSwapExactAmountIn{
			Sender:            tradeIcaAccount.Address,
			Routes:            routes,
			TokenIn:           tradeTokens,
			TokenOutMinAmount: minOut,
		})
		k.Logger(ctx).Info(utils.LogWithHostZone(tradeIcaAccount.ChainId,
			"Preparing MsgSwapExactAmountIn of %+v from the trade account", tradeTokens))
	}

	strideEpochTracker, found := k.GetEpochTracker(ctx, epochstypes.STRIDE_EPOCH)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, epochstypes.STRIDE_EPOCH)
	}
	timeout := uint64(strideEpochTracker.NextEpochStartTime)

	// Send the ICA tx to perform the swap on the tradeZone
	err := k.SubmitICATxWithoutCallback(ctx, tradeIcaAccount.ConnectionId, types.ICAAccountType_TRADE, msgs, timeout)
	if err != nil {
		return errorsmod.Wrapf(types.ErrICATxFailed, "Failed to submit ICA tx for the swap, Messages: %v, err: %s", msgs, err.Error())
	}

	return nil
}

// ICQ calls for remote ICA balances
// There is a single trade zone (hardcoded as Osmosis for now but maybe additional DEXes allowed in the future)
// We have to initialize a single hostZone object for the trade zone once in initialization and then it can be used in all these calls

// Kick off ICQ for the reward denom balance in the withdrawal address
func (k Keeper) WithdrawalRewardBalanceQuery(ctx sdk.Context, route types.TradeRoute) error {
	withdrawalAccount := route.HostToRewardHop.FromAccount
	k.Logger(ctx).Info(utils.LogWithHostZone(withdrawalAccount.ChainId, "Submitting ICQ for reward denom in withdrawal account"))

	// Encode the withdrawal account address for the query request
	// The query request consists of the withdrawal account address and reward denom
	_, withdrawalAddressBz, err := bech32.DecodeAndConvert(withdrawalAccount.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid withdrawal account address, could not decode (%s)", err.Error())
	}
	queryData := append(bankTypes.CreateAccountBalancesPrefix(withdrawalAddressBz), []byte(route.RewardDenomOnHostZone)...)

	// Timeout query at end of epoch
	strideEpochTracker, found := k.GetEpochTracker(ctx, epochstypes.STRIDE_EPOCH)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, epochstypes.STRIDE_EPOCH)
	}
	timeout := time.Unix(0, int64(strideEpochTracker.NextEpochStartTime))
	timeoutDuration := timeout.Sub(ctx.BlockTime())

	// The only callback data we need is the trade route
	callbackData := route
	callbackDataBz, err := proto.Marshal(&callbackData)
	if err != nil {
		return errorsmod.Wrapf(err, "unable to marshal TradeRoute callback data")
	}

	// Submit the ICQ for the withdrawal account balance
	query := icqtypes.Query{
		ChainId:         withdrawalAccount.ChainId,
		ConnectionId:    withdrawalAccount.ConnectionId,
		QueryType:       icqtypes.BANK_STORE_QUERY_WITH_PROOF,
		RequestData:     queryData,
		CallbackModule:  types.ModuleName,
		CallbackId:      ICQCallbackID_WithdrawalRewardBalance,
		CallbackData:    callbackDataBz,
		TimeoutDuration: timeoutDuration,
		TimeoutPolicy:   icqtypes.TimeoutPolicy_REJECT_QUERY_RESPONSE,
	}
	if err := k.InterchainQueryKeeper.SubmitICQRequest(ctx, query, false); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error querying for withdrawal reward denom balance, error: %s", err.Error()))
		return err
	}

	return nil
}

// Kick off ICQ for how many reward tokens are in the trade ICA associated with this host zone
func (k Keeper) TradeRewardBalanceQuery(ctx sdk.Context, route types.TradeRoute) error {
	tradeAccount := route.RewardToTradeHop.ToAccount
	k.Logger(ctx).Info(utils.LogWithHostZone(tradeAccount.ChainId, "Submitting ICQ for reward denom in trade ICA account"))

	// Encode the trade account address for the query request
	// The query request consists of the trade account address and reward denom
	// keep in mind this ICA address actually exists on trade zone but is associated with trades performed for host zone
	_, tradeAddressBz, err := bech32.DecodeAndConvert(tradeAccount.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid trade account address, could not decode (%s)", err.Error())
	}
	queryData := append(bankTypes.CreateAccountBalancesPrefix(tradeAddressBz), []byte(route.RewardDenomOnTradeZone)...)

	// Timeout query at end of epoch
	strideEpochTracker, found := k.GetEpochTracker(ctx, epochstypes.STRIDE_EPOCH)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, epochstypes.STRIDE_EPOCH)
	}
	timeout := time.Unix(0, int64(strideEpochTracker.NextEpochStartTime))
	timeoutDuration := timeout.Sub(ctx.BlockTime())

	// The only callback data we need is the trade route
	callbackData := route
	callbackDataBz, err := proto.Marshal(&callbackData)
	if err != nil {
		return errorsmod.Wrapf(err, "unable to marshal TradeRewardBalanceQuery callback data")
	}

	// Submit the ICQ for the withdrawal account balance
	query := icqtypes.Query{
		ChainId:         tradeAccount.ChainId,
		ConnectionId:    tradeAccount.ConnectionId, // query needs to go to the trade zone, not the host zone
		QueryType:       icqtypes.BANK_STORE_QUERY_WITH_PROOF,
		RequestData:     queryData,
		CallbackModule:  types.ModuleName,
		CallbackId:      ICQCallbackID_TradeRewardBalance,
		CallbackData:    callbackDataBz,
		TimeoutDuration: timeoutDuration,
		TimeoutPolicy:   icqtypes.TimeoutPolicy_REJECT_QUERY_RESPONSE,
	}
	if err := k.InterchainQueryKeeper.SubmitICQRequest(ctx, query, false); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error querying trade ICA for reward denom balance, error: %s", err.Error()))
		return err
	}

	return nil
}

// Kick off ICQ for how many converted tokens are in the trade ICA associated with this host zone
func (k Keeper) TradeConvertedBalanceQuery(ctx sdk.Context, route types.TradeRoute) error {
	tradeAccount := route.RewardToTradeHop.ToAccount
	k.Logger(ctx).Info(utils.LogWithHostZone(tradeAccount.ChainId, "Submitting ICQ for converted denom in trade ICA account"))

	// Encode the trade account address for the query request
	// The query request consists of the trade account address and converted denom
	// keep in mind this ICA address actually exists on trade zone but is associated with trades performed for host zone
	_, tradeAddressBz, err := bech32.DecodeAndConvert(tradeAccount.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid trade account address, could not decode (%s)", err.Error())
	}
	queryData := append(bankTypes.CreateAccountBalancesPrefix(tradeAddressBz), []byte(route.TargetDenomOnTradeZone)...)

	// Timeout query at end of epoch
	strideEpochTracker, found := k.GetEpochTracker(ctx, epochstypes.STRIDE_EPOCH)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, epochstypes.STRIDE_EPOCH)
	}
	timeout := time.Unix(0, int64(strideEpochTracker.NextEpochStartTime))
	timeoutDuration := timeout.Sub(ctx.BlockTime())

	// The only callback data we need is the trade route
	callbackData := route
	callbackDataBz, err := proto.Marshal(&callbackData)
	if err != nil {
		return errorsmod.Wrapf(err, "unable to marshal trade route as callback data")
	}

	// Submit the ICQ for the withdrawal account balance
	query := icqtypes.Query{
		ChainId:         tradeAccount.ChainId,
		ConnectionId:    tradeAccount.ConnectionId, // query needs to go to the trade zone, not the host zone
		QueryType:       icqtypes.BANK_STORE_QUERY_WITH_PROOF,
		RequestData:     queryData,
		CallbackModule:  types.ModuleName,
		CallbackId:      ICQCallbackID_TradeConvertedBalance,
		CallbackData:    callbackDataBz,
		TimeoutDuration: timeoutDuration,
		TimeoutPolicy:   icqtypes.TimeoutPolicy_REJECT_QUERY_RESPONSE,
	}
	if err := k.InterchainQueryKeeper.SubmitICQRequest(ctx, query, false); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error querying trade ICA for converted denom balance, error: %s", err.Error()))
		return err
	}

	return nil
}

// Kick off ICQ for the spot price on the pool given the input and output denoms implied by the given TradeRoute
//
//	the callback for this query is responsible for updating the returned spot price on the keeper data
func (k Keeper) PoolSpotPriceQuery(ctx sdk.Context, route types.TradeRoute) error {
	tradeAccount := route.RewardToTradeHop.ToAccount
	k.Logger(ctx).Info(utils.LogWithHostZone(tradeAccount.ChainId, "Submitting ICQ for spot price in this pool"))

	// Unlike balance queries, we can't do a normal ICQ which interacts with the foreign data store directly
	// because on Osmosis the spot price for a pool is not stored down to their keeper data, it is computed each request
	// Therefore we have to perform a query against their actual service, through the Osmosis spot price query route

	// The stride interchainquery module likely can't actually handle this type of query so this likely won't work yet...

	// Encode the osmosis spot price query request for the specific pool and denoms
	spotPriceRequest := types.QuerySpotPriceRequest{
		PoolId:          route.PoolId,
		BaseAssetDenom:  route.RewardDenomOnTradeZone,
		QuoteAssetDenom: route.TargetDenomOnTradeZone,
	}
	queryData := k.cdc.MustMarshal(&spotPriceRequest)

	// Timeout query at end of epoch
	strideEpochTracker, found := k.GetEpochTracker(ctx, epochstypes.STRIDE_EPOCH)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, epochstypes.STRIDE_EPOCH)
	}
	timeout := time.Unix(0, int64(strideEpochTracker.NextEpochStartTime))
	timeoutDuration := timeout.Sub(ctx.BlockTime())

	// The only callback data we will need is the trade route
	callbackData := route
	callbackDataBz, err := proto.Marshal(&callbackData)
	if err != nil {
		return errorsmod.Wrapf(err, "unable to marshal TradeRewardBalanceQuery callback data")
	}

	// Submit the ICQ for the trade pool spot price query
	query := icqtypes.Query{
		ChainId:         tradeAccount.ChainId,
		ConnectionId:    tradeAccount.ConnectionId, // query needs to go to the trade zone, not the host zone
		QueryType:       icqtypes.BANK_STORE_QUERY_WITH_PROOF,
		RequestData:     queryData,
		CallbackModule:  types.ModuleName,
		CallbackId:      ICQCallbackID_PoolSpotPrice,
		CallbackData:    callbackDataBz,
		TimeoutDuration: timeoutDuration,
		TimeoutPolicy:   icqtypes.TimeoutPolicy_REJECT_QUERY_RESPONSE,
	}
	if err := k.InterchainQueryKeeper.SubmitICQRequest(ctx, query, false); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error querying pool spot price, error: %s", err.Error()))
		return err
	}

	return nil
}

// The current design assumes foreign reward tokens start and end in the hostZone withdrawal address
// Step 1: transfer reward tokens to trade chain
// Step 2: perform the swap with as many reward tokens as possible
// Step 3: return the swapped tokens to the withdrawal ICA on hostZone
// Independently there is an ICQ to get the swap price and update it in the keeper state

// Because the swaps have limits on how many tokens can be used to avoid slippage,
// the swaps and price checks happen on a faster (hourly) cadence than the transfers (stride epochly)

// Helper function to be run stride epochly, kicks off queries on specific denoms on route
func (k Keeper) TransferAllRewardTokens(ctx sdk.Context) {
	for _, route := range k.GetAllTradeRoutes(ctx) {
		// Step 1: ICQ reward balance on hostZone, transfer funds with unwinding to trade chain
		k.WithdrawalRewardBalanceQuery(ctx, route)
		// Step 3: ICQ converted tokens in trade ICA, transfer funds back to hostZone withdrawal ICA
		k.TradeConvertedBalanceQuery(ctx, route)
	}
}

// Helper function to be run hourly, kicks off query which will kick off actual swaps to happen
func (k Keeper) SwapAllRewardTokens(ctx sdk.Context) {
	for _, route := range k.GetAllTradeRoutes(ctx) {
		// Step 2: ICQ reward balance in trade ICA, swap tokens according to limiting rules
		k.TradeRewardBalanceQuery(ctx, route)
	}
}

// Helper function to be run hourly, kicks off query to get and update the swap price in keeper data
func (k Keeper) UpdateAllSwapPrices(ctx sdk.Context) {
	for _, route := range k.GetAllTradeRoutes(ctx) {
		// ICQ swap price for the specific pair on this route and update keeper on callback
		k.PoolSpotPriceQuery(ctx, route)
	}
}