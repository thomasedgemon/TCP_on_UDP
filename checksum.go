//implement checksum algo and apply to data chunk

// use CRC32
package main

import (
	"hash/crc32"
)

func MakeChecksum(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
