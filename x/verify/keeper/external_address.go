package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vidulum/x/verify/types"
)

// SetExternalAddress set a specific externalAddress in the store from its index
func (k Keeper) SetExternalAddress(ctx sdk.Context, externalAddress types.ExternalAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExternalAddressKeyPrefix))
	b := k.cdc.MustMarshal(&externalAddress)
	store.Set(types.ExternalAddressKey(
		externalAddress.Vidulum,
	), b)
}

// GetExternalAddress returns a externalAddress from its index
func (k Keeper) GetExternalAddress(
	ctx sdk.Context,
	vidulum string,

) (val types.ExternalAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExternalAddressKeyPrefix))

	b := store.Get(types.ExternalAddressKey(
		vidulum,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExternalAddress removes a externalAddress from the store
func (k Keeper) RemoveExternalAddress(
	ctx sdk.Context,
	vidulum string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExternalAddressKeyPrefix))
	store.Delete(types.ExternalAddressKey(
		vidulum,
	))
}

// GetAllExternalAddress returns all externalAddress
func (k Keeper) GetAllExternalAddress(ctx sdk.Context) (list []types.ExternalAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExternalAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExternalAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
