package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"vidulum/x/verify/types"
)

func (k msgServer) EthereumOwnership(goCtx context.Context, msg *types.MsgEthereumOwnership) (*types.MsgEthereumOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEthereumOwnershipResponse{}, nil
}
