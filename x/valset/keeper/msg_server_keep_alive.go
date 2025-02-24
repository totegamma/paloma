package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/palomachain/paloma/x/valset/types"
)

func (k msgServer) KeepAlive(goCtx context.Context, msg *types.MsgKeepAlive) (*types.MsgKeepAliveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	valAddr := sdk.ValAddress(creator.Bytes())
	err := k.Keeper.KeepValidatorAlive(ctx, valAddr, msg.PigeonVersion)
	if err != nil {
		return nil, err
	}

	return &types.MsgKeepAliveResponse{}, nil
}
