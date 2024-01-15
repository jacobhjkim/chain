package keeper_test

import (
	"fmt"
	"math/big"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/settlus/chain/contracts"
	"github.com/settlus/chain/testutil/sample"
	utiltx "github.com/settlus/chain/testutil/tx"
	settlustypes "github.com/settlus/chain/types"
	"github.com/settlus/chain/x/settlement/types"
)

func (suite *SettlementTestSuite) TestSettle_Settle() {
	_, err := suite.msgServer.DepositToTreasury(suite.ctx, &types.MsgDepositToTreasury{
		TenantId: 1,
		Amount:   sdk.NewCoin("uusdc", math.NewInt(1000)),
		Sender:   suite.appAdmin.String(),
	})
	suite.NoError(err)

	requestId := "request-1"
	recipient := sdk.MustAccAddressFromBech32(sample.AccAddress())
	_, err = suite.keeper.CreateUTXR(suite.ctx, 1, &types.UTXR{
		RequestId:   requestId,
		Recipient:   settlustypes.NewHexAddressString(recipient),
		Amount:      sdk.NewCoin("uusdc", math.NewInt(10)),
		PayoutBlock: 10,
	})
	suite.NoError(err)

	utxr := suite.keeper.GetUTXRByRequestId(suite.ctx, 1, requestId)
	suite.NotNil(utxr)
	suite.EqualValues(10, utxr.Amount.Amount.Int64())

	initialBalance := s.app.BankKeeper.GetBalance(suite.ctx, recipient, "uusdc")
	suite.EqualValues(0, initialBalance.Amount.Int64())

	suite.keeper.Settle(suite.ctx.WithBlockHeight(11), 1)

	utxr = suite.keeper.GetUTXRByRequestId(suite.ctx.WithBlockHeight(12), 1, requestId)
	suite.Nil(utxr)

	balance := s.app.Erc20Keeper.BalanceOf(suite.ctx, contracts.ERC20Contract.ABI, suite.erc20Address, common.BytesToAddress(recipient.Bytes()))
	suite.EqualValues(big.NewInt(10), balance)
}

func (suite *SettlementTestSuite) TestSettle_Settle_InsufficientTreasuryBalance() {
	// give only 50
	_, err := suite.msgServer.DepositToTreasury(suite.ctx, &types.MsgDepositToTreasury{
		TenantId: 1,
		Amount:   sdk.NewCoin("uusdc", math.NewInt(50)),
		Sender:   suite.appAdmin.String(),
	})
	suite.NoError(err)

	recipient := sdk.AccAddress(utiltx.GenerateAddress().Bytes())

	// total amount to settle is 100
	for i := 0; i < 10; i++ {
		res, err := suite.keeper.CreateUTXR(suite.ctx, 1, &types.UTXR{
			RequestId:   fmt.Sprintf("request-%d", i),
			Recipient:   settlustypes.NewHexAddressString(recipient),
			Amount:      sdk.NewCoin("uusdc", math.NewInt(10)),
			PayoutBlock: uint64(suite.ctx.BlockHeight()) + 10,
		})

		suite.NoError(err)
		suite.NotNil(res)
	}

	suite.keeper.Settle(suite.ctx.WithBlockHeight(100), 1)

	// first utxr should be settled and deleted
	utxr := suite.keeper.GetUTXRByRequestId(suite.ctx, 1, "request-1")
	suite.Nil(utxr)

	// this should not
	utxr = suite.keeper.GetUTXRByRequestId(suite.ctx, 1, "request-9")
	suite.NotNil(utxr)

	// treasury balance should be 0
	treasury := types.GetTenantTreasuryAccount(1)
	coins := suite.app.BankKeeper.SpendableCoins(suite.ctx, treasury)
	suite.EqualValues(0, len(coins))

	// recipient balance should be 50
	balance := s.app.Erc20Keeper.BalanceOf(suite.ctx, contracts.ERC20Contract.ABI, suite.erc20Address, common.BytesToAddress(recipient.Bytes()))
	suite.EqualValues(big.NewInt(50), balance)
}

func (suite *SettlementTestSuite) TestSettle_Settle_TopUpTreasuryBalance() {
	_, err := suite.msgServer.DepositToTreasury(suite.ctx, &types.MsgDepositToTreasury{
		Sender:   suite.appAdmin.String(),
		TenantId: 1,
		Amount:   sdk.NewCoin("uusdc", math.NewInt(50)),
	})
	suite.NoError(err)

	recipient := sdk.AccAddress(utiltx.GenerateAddress().Bytes())

	_, err = suite.keeper.CreateUTXR(suite.ctx, 1, &types.UTXR{
		RequestId:   "request-1",
		Recipient:   settlustypes.NewHexAddressString(recipient),
		Amount:      sdk.NewCoin("uusdc", math.NewInt(100)),
		PayoutBlock: uint64(suite.ctx.BlockHeight()) + 10,
	})
	suite.NoError(err)

	utxr := suite.keeper.GetUTXRByRequestId(suite.ctx, 1, "request-1")
	suite.NotNil(utxr)

	suite.keeper.Settle(suite.ctx.WithBlockHeight(100), 1)

	utxr = suite.keeper.GetUTXRByRequestId(suite.ctx, 1, "request-1")
	suite.NotNil(utxr)

	_, err = suite.msgServer.DepositToTreasury(suite.ctx, &types.MsgDepositToTreasury{
		Sender:   suite.appAdmin.String(),
		TenantId: 1,
		Amount:   sdk.NewCoin("uusdc", math.NewInt(50)),
	})
	suite.NoError(err)

	suite.keeper.Settle(suite.ctx.WithBlockHeight(100), 1)

	utxr = suite.keeper.GetUTXRByRequestId(suite.ctx, 1, "request-1")
	suite.Nil(utxr)

	balance := s.app.Erc20Keeper.BalanceOf(suite.ctx, contracts.ERC20Contract.ABI, suite.erc20Address, common.BytesToAddress(recipient.Bytes()))
	suite.EqualValues(big.NewInt(100), balance)
}
