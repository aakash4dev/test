package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"test/x/test/types"
)

// GetListofCount get the total number of listof
func (k Keeper) GetListofCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ListofCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetListofCount set the total number of listof
func (k Keeper) SetListofCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ListofCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendListof appends a listof in the store with a new id and update the count
func (k Keeper) AppendListof(
	ctx sdk.Context,
	listof types.Listof,
) uint64 {
	// Create the listof
	count := k.GetListofCount(ctx)

	// Set the ID of the appended value
	listof.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ListofKey))
	appendedValue := k.cdc.MustMarshal(&listof)
	store.Set(GetListofIDBytes(listof.Id), appendedValue)

	// Update listof count
	k.SetListofCount(ctx, count+1)

	return count
}

// SetListof set a specific listof in the store
func (k Keeper) SetListof(ctx sdk.Context, listof types.Listof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ListofKey))
	b := k.cdc.MustMarshal(&listof)
	store.Set(GetListofIDBytes(listof.Id), b)
}

// GetListof returns a listof from its id
func (k Keeper) GetListof(ctx sdk.Context, id uint64) (val types.Listof, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ListofKey))
	b := store.Get(GetListofIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveListof removes a listof from the store
func (k Keeper) RemoveListof(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ListofKey))
	store.Delete(GetListofIDBytes(id))
}

// GetAllListof returns all listof
func (k Keeper) GetAllListof(ctx sdk.Context) (list []types.Listof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ListofKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Listof
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetListofIDBytes returns the byte representation of the ID
func GetListofIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetListofIDFromBytes returns ID in uint64 format from a byte array
func GetListofIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
