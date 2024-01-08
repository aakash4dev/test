package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateListof = "create_listof"
	TypeMsgUpdateListof = "update_listof"
	TypeMsgDeleteListof = "delete_listof"
)

var _ sdk.Msg = &MsgCreateListof{}

func NewMsgCreateListof(creator string, something string) *MsgCreateListof {
	return &MsgCreateListof{
		Creator:   creator,
		Something: something,
	}
}

func (msg *MsgCreateListof) Route() string {
	return RouterKey
}

func (msg *MsgCreateListof) Type() string {
	return TypeMsgCreateListof
}

func (msg *MsgCreateListof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateListof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateListof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateListof{}

func NewMsgUpdateListof(creator string, id uint64, something string) *MsgUpdateListof {
	return &MsgUpdateListof{
		Id:        id,
		Creator:   creator,
		Something: something,
	}
}

func (msg *MsgUpdateListof) Route() string {
	return RouterKey
}

func (msg *MsgUpdateListof) Type() string {
	return TypeMsgUpdateListof
}

func (msg *MsgUpdateListof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateListof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateListof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteListof{}

func NewMsgDeleteListof(creator string, id uint64) *MsgDeleteListof {
	return &MsgDeleteListof{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteListof) Route() string {
	return RouterKey
}

func (msg *MsgDeleteListof) Type() string {
	return TypeMsgDeleteListof
}

func (msg *MsgDeleteListof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteListof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteListof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
