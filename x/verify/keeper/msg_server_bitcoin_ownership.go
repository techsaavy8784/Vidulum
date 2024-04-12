package keeper

import (
	"context"

	"github.com/libsv/go-bk/crypto"

	"vidulum/x/verify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/bitcoinsv/bsvd/chaincfg/chainhash"
	"github.com/bitcoinsv/bsvd/wire"
	"github.com/libsv/go-bk/bec"
	"github.com/libsv/go-bt/v2/bscript"
)

const (
	// hBSV is the magic header string required fore Bitcoin Signed Messages
	hBSV string = "Bitcoin Signed Message:\n"
)

func (k msgServer) BitcoinOwnership(goCtx context.Context, msg *types.MsgBitcoinOwnership) (*types.MsgBitcoinOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Reconstruct the pubkey
	publicKey, wasCompressed, err := k.PubKeyFromSignature(msg.Signature, msg.Message)
	if err != nil {
		return nil, err
	}

	// Get the address
	var bscriptAddress *bscript.Address
	if bscriptAddress, err = k.GetAddressFromPubKey(publicKey, wasCompressed); err != nil {
		return nil, err
	}

	// Return nil if addresses match.
	if bscriptAddress.AddressString == msg.Address {
		return nil, nil
	}

	var ExternalAddress = types.ExternalAddress{
		Vidulum: msg.Owner,
		Bitcoin: msg.Address,
	}
	k.SetExternalAddress(ctx, ExternalAddress)

	return &types.MsgBitcoinOwnershipResponse{}, nil
}

func (k Keeper) PubKeyFromSignature(sig, data string) (pubKey *bec.PublicKey, wasCompressed bool, err error) {

	var decodedSig []byte
	if decodedSig, err = base64.StdEncoding.DecodeString(sig); err != nil {
		return nil, false, err
	}

	// Validate the signature - this just shows that it was valid at all
	// we will compare it with the key next
	var buf bytes.Buffer
	if err = wire.WriteVarString(&buf, 0, hBSV); err != nil {
		return nil, false, err
	}
	if err = wire.WriteVarString(&buf, 0, data); err != nil {
		return nil, false, err
	}

	// Create the hash
	expectedMessageHash := chainhash.DoubleHashB(buf.Bytes())
	return bec.RecoverCompact(bec.S256(), decodedSig, expectedMessageHash)
}

// GetAddressFromPubKey gets a bscript.Address from a bec.PublicKey
func (k Keeper) GetAddressFromPubKey(publicKey *bec.PublicKey, compressed bool) (*bscript.Address, error) {
	if publicKey == nil {
		return nil, fmt.Errorf("publicKey cannot be nil")
	} else if publicKey.X == nil {
		return nil, fmt.Errorf("publicKey.X cannot be nil")
	}

	if !compressed {
		// go-bt/v2/bscript does not have a function that exports the uncompressed address
		// https://github.com/libsv/go-bt/blob/master/bscript/address.go#L98
		hash := crypto.Hash160(publicKey.SerialiseUncompressed())
		bb := make([]byte, 1)
		//nolint: makezero // we need to set up the array with 1
		bb = append(bb, hash...)
		return &bscript.Address{
			AddressString: bscript.Base58EncodeMissingChecksum(bb),
			PublicKeyHash: hex.EncodeToString(hash),
		}, nil
	}

	return bscript.NewAddressFromPublicKey(publicKey, true)
}
