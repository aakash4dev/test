package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"test/x/test/types"
)

func (k Keeper) ListofAll(goCtx context.Context, req *types.QueryAllListofRequest) (*types.QueryAllListofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var listofs []types.Listof
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	listofStore := prefix.NewStore(store, types.KeyPrefix(types.ListofKey))

	pageRes, err := query.Paginate(listofStore, req.Pagination, func(key []byte, value []byte) error {
		var listof types.Listof
		if err := k.cdc.Unmarshal(value, &listof); err != nil {
			return err
		}

		listofs = append(listofs, listof)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllListofResponse{Listof: listofs, Pagination: pageRes}, nil
}

func (k Keeper) Listof(goCtx context.Context, req *types.QueryGetListofRequest) (*types.QueryGetListofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	listof, found := k.GetListof(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetListofResponse{Listof: listof}, nil
}
