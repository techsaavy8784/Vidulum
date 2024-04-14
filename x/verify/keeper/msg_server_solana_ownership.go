package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"crypto/ed25519"
	"vidulum/x/verify/types"

	"github.com/mr-tron/base58"
)

func (k msgServer) SolanaOwnership(goCtx context.Context, msg *types.MsgSolanaOwnership) (*types.MsgSolanaOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetExternalAddress(ctx, msg.Owner)
	if isFound && valFound.Solana != "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "you already own Solana address")
	}

	// Decode Solana address from base58
	decodedAddress, err := base58.Decode(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrDecodeAddress, "failed to decode Solana address")
	}

	// Decode Solana signature from base58
	decodedSignature, err := base58.Decode(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrDecodeSignature, "failed to decode Solana sinature")
	}

	// Verify Solana signature
	isVerified := ed25519.Verify(decodedAddress, []byte(msg.Message), decodedSignature)
	if isVerified {
		var externalAddress = types.ExternalAddress{
			Vidulum:  msg.Owner,
			Bitcoin:  valFound.Bitcoin,
			Ethereum: valFound.Ethereum,
			Solana:   msg.Address,
			Zcash:    valFound.Zcash,
		}
		k.SetExternalAddress(ctx, externalAddress)

		ctx.EventManager().EmitEvent(sdk.NewEvent(
			"Solana address verified",
			sdk.NewAttribute("Owner", msg.Owner),
			sdk.NewAttribute("Solana address", msg.Address),
		))
		return &types.MsgSolanaOwnershipResponse{IsVerified: true}, nil
	} else {
		return nil, sdkerrors.Wrap(types.ErrVerifyAddress, "failed to verify Solana address")
	}
}
