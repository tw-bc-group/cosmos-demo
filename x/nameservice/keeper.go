package nameservice

import (
	"github.com/arthaszeng/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Sets the entire Whois metadata struct for a name
func (keeper Keeper) SetWhois(ctx sdk.Context, name string, whois types.Whois) {
	if whois.Owner.Empty() {
		return
	}
	store := ctx.KVStore(keeper.storeKey)
	store.Set([]byte(name), keeper.cdc.MustMarshalBinaryBare(whois))
}

// Gets the entire Whois metadata struct for a name
func (keeper Keeper) GetWhois(ctx sdk.Context, name string) types.Whois {
	store := ctx.KVStore(keeper.storeKey)
	if !store.Has([]byte(name)) {
		return types.NewWhois()
	}
	bytesOfValue := store.Get([]byte(name))
	var whois types.Whois
	keeper.cdc.MustUnmarshalBinaryBare(bytesOfValue, &whois)
	return whois
}

// ResolveName - returns the string that the name resolves to
func (keeper Keeper) ResolveName(ctx sdk.Context, name string) string {
	return keeper.GetWhois(ctx, name).Value
}

// SetName - sets the value string that a name resolves to
func (keeper Keeper) SetName(ctx sdk.Context, name string, value string) {
	whois := keeper.GetWhois(ctx, name)
	whois.Value = value
	keeper.SetWhois(ctx, name, whois)
}

// HasOwner - returns whether or not the name already has an owner
func (keeper Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !keeper.GetWhois(ctx, name).Owner.Empty()
}

// GetOwner - get the current owner of a name
func (keeper Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return keeper.GetWhois(ctx, name).Owner
}

// SetOwner - sets the current owner of a name
func (keeper Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	whois := keeper.GetWhois(ctx, name)
	whois.Owner = owner
	keeper.SetWhois(ctx, name, whois)
}

// GetPrice - gets the current price of a name.  If price doesn't exist yet, set to 1nametoken.
func (keeper Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	return keeper.GetWhois(ctx, name).Price
}

// SetPrice - sets the current price of a name
func (keeper Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois := keeper.GetWhois(ctx, name)
	whois.Price = price
	keeper.SetWhois(ctx, name, whois)
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (keeper Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}
