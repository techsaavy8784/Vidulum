package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"vidulum/x/verify/types"
)

func (k Keeper) ExternalAddressAll(goCtx context.Context, req *types.QueryAllExternalAddressRequest) (*types.QueryAllExternalAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var externalAddresss []types.ExternalAddress
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	externalAddressStore := prefix.NewStore(store, types.KeyPrefix(types.ExternalAddressKeyPrefix))

	pageRes, err := query.Paginate(externalAddressStore, req.Pagination, func(key []byte, value []byte) error {
		var externalAddress types.ExternalAddress
		if err := k.cdc.Unmarshal(value, &externalAddress); err != nil {
			return err
		}

		externalAddresss = append(externalAddresss, externalAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExternalAddressResponse{ExternalAddress: externalAddresss, Pagination: pageRes}, nil
}

func (k Keeper) ExternalAddress(goCtx context.Context, req *types.QueryGetExternalAddressRequest) (*types.QueryGetExternalAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetExternalAddress(
		ctx,
		req.Vidulum,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetExternalAddressResponse{ExternalAddress: val}, nil
}
