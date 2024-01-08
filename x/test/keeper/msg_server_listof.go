package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"test/x/test/types"
)

func (k msgServer) CreateListof(goCtx context.Context, msg *types.MsgCreateListof) (*types.MsgCreateListofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var listof = types.Listof{
		Creator:   msg.Creator,
		Something: msg.Something,
	}

	id := k.AppendListof(
		ctx,
		listof,
	)

	return &types.MsgCreateListofResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateListof(goCtx context.Context, msg *types.MsgUpdateListof) (*types.MsgUpdateListofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var listof = types.Listof{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Something: msg.Something,
	}

	// Checks that the element exists
	val, found := k.GetListof(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetListof(ctx, listof)

	return &types.MsgUpdateListofResponse{}, nil
}

func (k msgServer) DeleteListof(goCtx context.Context, msg *types.MsgDeleteListof) (*types.MsgDeleteListofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetListof(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveListof(ctx, msg.Id)

	return &types.MsgDeleteListofResponse{}, nil
}
