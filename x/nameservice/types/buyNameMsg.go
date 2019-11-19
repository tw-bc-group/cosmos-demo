package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

const RouterKey = StoreKey

func (msg MsgBuyName) Route() string { return RouterKey }

func (msg MsgBuyName) Type() string {
	return "buy_name"
}

func (msg MsgBuyName) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

func (msg MsgBuyName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

var _ sdk.Msg = new(MsgBuyName)
