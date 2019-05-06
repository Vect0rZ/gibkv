package util

import "hash/fnv"

// Hash - hashes a string to a Unsigned 32 bit integer
func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))

	return h.Sum32()
}
