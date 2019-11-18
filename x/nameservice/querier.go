package nameservice

import (
	"github.com/arthaszeng/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
	QueryNames   = "names"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], keeper)
		case QueryWhois:
			return queryWhois(ctx, path[1:], keeper)
		case QueryNames:
			return queryNames(ctx, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown name service query endpoint.")
		}
	}
}

func queryNames(ctx sdk.Context, keeper Keeper) ([]byte, sdk.Error) {
	var namesList types.QueryResNames

	iterator := keeper.GetNamesIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		namesList = append(namesList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryWhois(ctx sdk.Context, path []string, keeper Keeper) ([]byte, sdk.Error) {
	whois := keeper.GetWhois(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, whois)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryResolve(ctx sdk.Context, path []string, keeper Keeper) ([]byte, sdk.Error) {
	value := keeper.ResolveName(ctx, path[0])

	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: value})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
