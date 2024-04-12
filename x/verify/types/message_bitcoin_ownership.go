package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBitcoinOwnership = "bitcoin_ownership"

var _ sdk.Msg = &MsgBitcoinOwnership{}

func NewMsgBitcoinOwnership(owner string, address string, signature string, message string) *MsgBitcoinOwnership {
	return &MsgBitcoinOwnership{
		Owner:     owner,
		Address:   address,
		Signature: signature,
		Message:   message,
	}
}

func (msg *MsgBitcoinOwnership) Route() string {
	return RouterKey
}

func (msg *MsgBitcoinOwnership) Type() string {
	return TypeMsgBitcoinOwnership
}

func (msg *MsgBitcoinOwnership) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgBitcoinOwnership) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBitcoinOwnership) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
