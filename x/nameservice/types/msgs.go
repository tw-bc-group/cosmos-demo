package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

func (msg MsgSetName) Route() string {
	panic("implement me")
}

func (msg MsgSetName) Type() string {
	panic("implement me")
}

func (msg MsgSetName) ValidateBasic() sdk.Error {
	panic("implement me")
}

func (msg MsgSetName) GetSignBytes() []byte {
	panic("implement me")
}

func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	panic("implement me")
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

var _ sdk.Msg = new(MsgSetName)
