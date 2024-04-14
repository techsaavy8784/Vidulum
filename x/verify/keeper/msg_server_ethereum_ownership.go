package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"vidulum/x/verify/types"
)

func (k msgServer) EthereumOwnership(goCtx context.Context, msg *types.MsgEthereumOwnership) (*types.MsgEthereumOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetExternalAddress(ctx, msg.Owner)
	if isFound && valFound.Ethereum != "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "you already own a Ethereum address")
	}

	sig := hexutil.MustDecode(msg.Signature)
	msgHash := accounts.TextHash([]byte(msg.Message))
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msgHash, sig)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to verify the ownership of ethereum address")
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	isVerified := strings.EqualFold(msg.Address, recoveredAddr.Hex())
	if isVerified {
		var externalAddress = types.ExternalAddress{
			Vidulum:  msg.Owner,
			Bitcoin:  valFound.Bitcoin,
			Ethereum: msg.Address,
			Solana:   valFound.Solana,
			Zcash:    valFound.Zcash,
		}

		k.SetExternalAddress(ctx, externalAddress)

		// Emit an event for the verified Bitcoin address
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				"ethereum address verified",
				sdk.NewAttribute("owner", msg.Owner),
				sdk.NewAttribute("ethereum_address", msg.Address),
			),
		)
		return &types.MsgEthereumOwnershipResponse{IsVerified: true}, nil
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "failed to verify the ownership of ethereum address")
	}
}
