package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

func (whois Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, whois.Owner, whois.Value, whois.Price))
}

var miniPrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

func NewWhois() Whois {
	return Whois{Price: miniPrice}
}
