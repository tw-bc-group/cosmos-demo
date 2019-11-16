package nameservice

import sdk "github.com/cosmos/cosmos-sdk/types"

type Whois struct {
	Value string
	Owner sdk.AccAddress
	Price sdk.Coins
}

var miniPrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

func NewWhois() Whois {
	return Whois{Price: miniPrice}
}
