package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSolanaOwnership = "solana_ownership"

var _ sdk.Msg = &MsgSolanaOwnership{}

func NewMsgSolanaOwnership(owner string, address string, signature string, message string) *MsgSolanaOwnership {
	return &MsgSolanaOwnership{
		Owner:     owner,
		Address:   address,
		Signature: signature,
		Message:   message,
	}
}

func (msg *MsgSolanaOwnership) Route() string {
	return RouterKey
}

func (msg *MsgSolanaOwnership) Type() string {
	return TypeMsgSolanaOwnership
}

func (msg *MsgSolanaOwnership) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgSolanaOwnership) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSolanaOwnership) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
