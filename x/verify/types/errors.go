package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/verify module sentinel errors
var (
	ErrSample            = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrReconstructPubKey = sdkerrors.Register(ModuleName, 2, "failed to reconstruct public key")
	ErrGetBitcoinAddress = sdkerrors.Register(ModuleName, 3, "failed to get bitcoin address")
	ErrRecoveredAddress  = sdkerrors.Register(ModuleName, 4, "failed to recover ECDSA of address")
	ErrDecodeAddress     = sdkerrors.Register(ModuleName, 5, "failed to decode address")
	ErrDecodeSignature   = sdkerrors.Register(ModuleName, 5, "failed to decode signature")
	ErrVerifyAddress     = sdkerrors.Register(ModuleName, 6, "failed to verify address")
)
