package keeper_test

import (
	"encoding/json"
	"time"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	transfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v5/testing"

	icqtypes "github.com/Stride-Labs/stride/v8/x/interchainquery/types"

	"github.com/Stride-Labs/stride/v8/x/stakeibc/keeper"
	"github.com/Stride-Labs/stride/v8/x/stakeibc/types"
)

var (
	ValAddress        = "cosmosvaloper1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrdt795p"
	LSMTokenBaseDenom = ValAddress + "/32"
)

type LSMLiquidStakeTestCase struct {
	hostZone            types.HostZone
	liquidStakerAddress sdk.AccAddress
	moduleAddress       sdk.AccAddress
	initialBalance      sdkmath.Int
	lsmTokenIBCDenom    string
	validMsg            *types.MsgLSMLiquidStake
}

func (s *KeeperTestSuite) SetupTestLSMLiquidStake() LSMLiquidStakeTestCase {
	s.CreateTransferChannel(HostChainId)

	initialBalance := sdkmath.NewInt(3_000_000)
	stakeAmount := sdkmath.NewInt(1_000_000)
	userAddress := s.TestAccs[0]
	moduleAddress := types.NewZoneAddress(HostChainId)

	// Need valid IBC denom here to test parsing
	sourcePrefix := transfertypes.GetDenomPrefix(transfertypes.PortID, ibctesting.FirstChannelID)
	prefixedDenom := sourcePrefix + LSMTokenBaseDenom
	lsmTokenDenomTrace := transfertypes.ParseDenomTrace(prefixedDenom)
	s.App.TransferKeeper.SetDenomTrace(s.Ctx, lsmTokenDenomTrace)
	lsmTokenIBCDenom := lsmTokenDenomTrace.IBCDenom()

	// Fund the user's account with the LSM token
	s.FundAccount(userAddress, sdk.NewCoin(lsmTokenIBCDenom, initialBalance))

	// Add the query exchange rate interval
	params := types.DefaultParams()
	params.ValidatorExchangeRateQueryInterval = 10_000_000
	s.App.StakeibcKeeper.SetParams(s.Ctx, params)

	// Add the host zone with a valid zone address as the LSM custodian
	hostZone := types.HostZone{
		ChainId:           HostChainId,
		HostDenom:         Atom,
		RedemptionRate:    sdk.NewDec(1.0),
		Address:           moduleAddress.String(),
		TransferChannelId: ibctesting.FirstChannelID,
		ConnectionId:      ibctesting.FirstConnectionID,
		Validators: []*types.Validator{{
			Address:                          ValAddress,
			ProgressTowardsExchangeRateQuery: sdkmath.NewInt(8_000_000),
		}},
		DelegationAccount: &types.ICAAccount{
			Address: "cosmos_DELEGATION",
		},
	}
	s.App.StakeibcKeeper.SetHostZone(s.Ctx, hostZone)

	return LSMLiquidStakeTestCase{
		hostZone:            hostZone,
		liquidStakerAddress: userAddress,
		moduleAddress:       moduleAddress,
		initialBalance:      initialBalance,
		lsmTokenIBCDenom:    lsmTokenIBCDenom,
		validMsg: &types.MsgLSMLiquidStake{
			Creator:          userAddress.String(),
			LsmTokenIbcDenom: lsmTokenIBCDenom,
			Amount:           stakeAmount,
		},
	}
}

func (s *KeeperTestSuite) TestLSMLiquidStake_Successful_NoExchangeRateQuery() {
	tc := s.SetupTestLSMLiquidStake()

	// Get the sequence number before the IBC Transfer is submitted to confirm it incremented
	startSequence, found := s.App.IBCKeeper.ChannelKeeper.GetNextSequenceSend(s.Ctx, transfertypes.PortID, ibctesting.FirstChannelID)
	s.Require().True(found, "sequence number not found before lsm liquid stake")

	// Call LSM Liquid stake with a valid message
	msgResponse, err := s.GetMsgServer().LSMLiquidStake(sdk.WrapSDKContext(s.Ctx), tc.validMsg)
	s.Require().NoError(err, "no error expected when calling lsm liquid stake")
	s.Require().True(msgResponse.TransactionComplete, "transaction should be complete")

	// Confirm the LSM token was sent to the protocol
	userLsmBalance := s.App.BankKeeper.GetBalance(s.Ctx, tc.liquidStakerAddress, tc.lsmTokenIBCDenom)
	s.Require().Equal(tc.initialBalance.Sub(tc.validMsg.Amount).Int64(), userLsmBalance.Amount.Int64(),
		"lsm token balance of user account")

	// Confirm stToken was sent to the user
	userStTokenBalance := s.App.BankKeeper.GetBalance(s.Ctx, tc.liquidStakerAddress, StAtom)
	s.Require().Equal(tc.validMsg.Amount.Int64(), userStTokenBalance.Amount.Int64(), "user stToken balance")

	// Confirm an LSMDeposit was created
	expectedDeposit := types.LSMTokenDeposit{
		ChainId:          HostChainId,
		Denom:            LSMTokenBaseDenom,
		ValidatorAddress: ValAddress,
		Amount:           tc.validMsg.Amount,
		Status:           types.TRANSFER_IN_PROGRESS,
	}
	actualDeposit, found := s.App.StakeibcKeeper.GetLSMTokenDeposit(s.Ctx, HostChainId, LSMTokenBaseDenom)
	s.Require().True(found, "lsm token deposit should have been found after LSM liquid stake")
	s.Require().Equal(expectedDeposit, actualDeposit)

	// Confirm IBC transfer was sent
	endSequence, found := s.App.IBCKeeper.ChannelKeeper.GetNextSequenceSend(s.Ctx, transfertypes.PortID, ibctesting.FirstChannelID)
	s.Require().True(found, "sequence number not found after lsm liquid stake")
	s.Require().Equal(startSequence+1, endSequence, "sequence number after IBC transfer")
}

func (s *KeeperTestSuite) TestLSMLiquidStake_Successful_WithExchangeRateQuery() {
	tc := s.SetupTestLSMLiquidStake()

	// Increase the liquid stake size so that it breaks the query checkpoint
	tc.validMsg.Amount = sdk.NewInt(3_000_000)

	// Call LSM Liquid stake
	msgResponse, err := s.GetMsgServer().LSMLiquidStake(sdk.WrapSDKContext(s.Ctx), tc.validMsg)
	s.Require().NoError(err, "no error expected when calling lsm liquid stake")
	s.Require().False(msgResponse.TransactionComplete, "transaction should still be pending")

	// Confirm the LSM token was escrowed
	userLsmBalance := s.App.BankKeeper.GetBalance(s.Ctx, tc.liquidStakerAddress, tc.lsmTokenIBCDenom)
	s.Require().Equal(tc.initialBalance.Sub(tc.validMsg.Amount).Int64(), userLsmBalance.Amount.Int64(),
		"lsm token balance of user account")

	moduleLsmBalance := s.App.BankKeeper.GetBalance(s.Ctx, tc.moduleAddress, tc.lsmTokenIBCDenom)
	s.Require().Equal(tc.validMsg.Amount.Int64(), moduleLsmBalance.Amount.Int64(),
		"lsm token balance of module account")

	// Confirm stToken was NOT sent to the user
	userStTokenBalance := s.App.BankKeeper.GetBalance(s.Ctx, tc.liquidStakerAddress, StAtom)
	s.Require().True(userStTokenBalance.Amount.IsZero(), "user stToken balance")

	// Confirm query was submitted
	allQueries := s.App.InterchainqueryKeeper.AllQueries(s.Ctx)
	s.Require().Len(allQueries, 1)

	// Confirm query metadata
	actualQuery := allQueries[0]
	s.Require().Equal(HostChainId, actualQuery.ChainId, "query chain-id")
	s.Require().Equal(ibctesting.FirstConnectionID, actualQuery.ConnectionId, "query connection-id")
	s.Require().Equal(icqtypes.STAKING_STORE_QUERY_WITH_PROOF, actualQuery.QueryType, "query types")

	s.Require().Equal(types.ModuleName, actualQuery.CallbackModule, "callback module")
	s.Require().Equal(keeper.ICQCallbackID_Validator, actualQuery.CallbackId, "callback-id")

	expectedTimeout := uint64(s.Ctx.BlockTime().UnixNano() + (time.Minute * 15).Nanoseconds())
	s.Require().Equal(expectedTimeout, actualQuery.Timeout, "callback module")

	// Confirm query callback data
	s.Require().True(len(actualQuery.CallbackData) > 0, "callback data exists")

	expectedLSMTokenDeposit := types.LSMTokenDeposit{
		ChainId:          HostChainId,
		Denom:            LSMTokenBaseDenom,
		ValidatorAddress: ValAddress,
		Amount:           tc.validMsg.Amount,
		Status:           types.DEPOSIT_PENDING,
	}

	var actualCallbackData keeper.LSMLiquidStake
	err = json.Unmarshal(actualQuery.CallbackData, &actualCallbackData)
	s.Require().NoError(err, "no error expected when unmarshalling query callback data")

	s.Require().Equal(sdk.NewCoin(tc.lsmTokenIBCDenom, tc.validMsg.Amount), actualCallbackData.LSMIBCToken, "callback data - lsm token")
	s.Require().Equal(sdk.NewCoin(StAtom, tc.validMsg.Amount), actualCallbackData.StToken, "callback data - stToken")

	s.Require().Equal(tc.liquidStakerAddress, actualCallbackData.Staker, "callback data - staker")
	s.Require().Equal(HostChainId, actualCallbackData.HostZone.ChainId, "callback data - host zone")
	s.Require().Equal(ValAddress, actualCallbackData.Validator.Address, "callback data - validator")

	s.Require().Equal(expectedLSMTokenDeposit, actualCallbackData.Deposit, "callback data - deposit")

}

func (s *KeeperTestSuite) TestLSMLiquidStakeFailed_NotIBCDenom() {
	tc := s.SetupTestLSMLiquidStake()

	// Change the message so that the denom is not an IBC token
	invalidMsg := tc.validMsg
	invalidMsg.LsmTokenIbcDenom = "fake_ibc_denom"

	_, err := s.GetMsgServer().LSMLiquidStake(sdk.WrapSDKContext(s.Ctx), invalidMsg)
	s.Require().ErrorContains(err, "lsm token is not an IBC token (fake_ibc_denom)")
}

func (s *KeeperTestSuite) TestLSMLiquidStakeFailed_HostZoneNotFound() {
	tc := s.SetupTestLSMLiquidStake()

	// Change the message so that the denom is an IBC denom from a channel that is not supported
	sourcePrefix := transfertypes.GetDenomPrefix(transfertypes.PortID, "channel-1")
	prefixedDenom := sourcePrefix + LSMTokenBaseDenom
	lsmTokenDenomTrace := transfertypes.ParseDenomTrace(prefixedDenom)
	s.App.TransferKeeper.SetDenomTrace(s.Ctx, lsmTokenDenomTrace)

	invalidMsg := tc.validMsg
	invalidMsg.LsmTokenIbcDenom = lsmTokenDenomTrace.IBCDenom()

	_, err := s.GetMsgServer().LSMLiquidStake(sdk.WrapSDKContext(s.Ctx), invalidMsg)
	s.Require().ErrorContains(err, "transfer channel-id from LSM token (channel-1) does not match any registered host zone")
}

func (s *KeeperTestSuite) TestLSMLiquidStakeFailed_ValidatorNotFound() {
	tc := s.SetupTestLSMLiquidStake()

	// Change the message so that the base denom is from a non-existent validator
	sourcePrefix := transfertypes.GetDenomPrefix(transfertypes.PortID, ibctesting.FirstChannelID)
	prefixedDenom := sourcePrefix + "cosmosvaloperXXX/42"
	lsmTokenDenomTrace := transfertypes.ParseDenomTrace(prefixedDenom)
	s.App.TransferKeeper.SetDenomTrace(s.Ctx, lsmTokenDenomTrace)

	invalidMsg := tc.validMsg
	invalidMsg.LsmTokenIbcDenom = lsmTokenDenomTrace.IBCDenom()

	_, err := s.GetMsgServer().LSMLiquidStake(sdk.WrapSDKContext(s.Ctx), invalidMsg)
	s.Require().ErrorContains(err, "validator (cosmosvaloperXXX) is not registered in the Stride validator set")
}

func (s *KeeperTestSuite) TestLSMLiquidStakeFailed_InvalidHostZoneAddress() {
	tc := s.SetupTestLSMLiquidStake()

	// Remove the host zones address from the store
	invalidHostZone := tc.hostZone
	invalidHostZone.Address = ""
	s.App.StakeibcKeeper.SetHostZone(s.Ctx, invalidHostZone)

	_, err := s.GetMsgServer().LSMLiquidStake(sdk.WrapSDKContext(s.Ctx), tc.validMsg)
	s.Require().ErrorContains(err, "host zone address is invalid")
}

func (s *KeeperTestSuite) TestLSMLiquidStakeFailed_InsufficientBalance() {
	tc := s.SetupTestLSMLiquidStake()

	// Send out all the user's coins so that they have an insufficient balance of LSM tokens
	initialBalanceCoin := sdk.NewCoins(sdk.NewCoin(tc.lsmTokenIBCDenom, tc.initialBalance))
	s.App.BankKeeper.SendCoins(s.Ctx, tc.liquidStakerAddress, s.TestAccs[1], initialBalanceCoin)

	_, err := s.GetMsgServer().LSMLiquidStake(sdk.WrapSDKContext(s.Ctx), tc.validMsg)
	s.Require().ErrorContains(err, "insufficient funds")
}