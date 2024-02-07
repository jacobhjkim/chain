package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"

	"github.com/settlus/chain/x/nobleforwarding/types"
)

const transferPort = "transfer"

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// RegisterAccount implements the Msg/RegisterAccount interface
func (m msgServer) RegisterAccount(goCtx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	packet := types.ForwardingAccountPacket{Recipient: msg.Sender}
	if err := packet.ValidateBasic(); err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidPacket, "failed to validate packet: %s", err)
	}
	packetBytes, err := packet.GetBytes()
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidPacket, "failed to marshal packet: %s", err)
	}

	channelCap, ok := m.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(transferPort, msg.ChannelId))
	if !ok {
		m.Logger(ctx).Error("failed to get channel capability")
		return nil, errorsmod.Wrapf(types.ErrInvalidChannelCapability, "failed to get channel capability for port %s and channel %s", transferPort, msg.ChannelId)
	}

	if _, err := m.channelKeeper.SendPacket(
		ctx,
		channelCap,
		transferPort,
		msg.ChannelId,
		clienttypes.ZeroHeight(),
		transfertypes.DefaultRelativePacketTimeoutTimestamp,
		packetBytes,
	); err != nil {
		return nil, errorsmod.Wrapf(types.ErrFailedToSendPacket, "failed to send packet: %s", err)
	}

	return &types.MsgRegisterAccountResponse{}, nil
}
