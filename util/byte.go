package util

import "encoding/binary"

// UInt32ToBytes - converts a given integer to a [4]byte
func UInt32ToBytes(i uint32) []byte {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, i)

	return data
}
