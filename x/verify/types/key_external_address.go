package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExternalAddressKeyPrefix is the prefix to retrieve all ExternalAddress
	ExternalAddressKeyPrefix = "ExternalAddress/value/"
)

// ExternalAddressKey returns the store key to retrieve a ExternalAddress from the index fields
func ExternalAddressKey(
	vidulum string,
) []byte {
	var key []byte

	vidulumBytes := []byte(vidulum)
	key = append(key, vidulumBytes...)
	key = append(key, []byte("/")...)

	return key
}
