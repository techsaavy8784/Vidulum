package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEthereumOwnership = "ethereum_ownership"

var _ sdk.Msg = &MsgEthereumOwnership{}

func NewMsgEthereumOwnership(owner string, address string, signature string, message string) *MsgEthereumOwnership {
	return &MsgEthereumOwnership{
		Owner:     owner,
		Address:   address,
		Signature: signature,
		Message:   message,
	}
}

func (msg *MsgEthereumOwnership) Route() string {
	return RouterKey
}

func (msg *MsgEthereumOwnership) Type() string {
	return TypeMsgEthereumOwnership
}

func (msg *MsgEthereumOwnership) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgEthereumOwnership) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEthereumOwnership) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
