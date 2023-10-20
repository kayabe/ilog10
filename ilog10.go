package ilog10

import (
	"math/bits"
)

func FastUint64(v uint64) (r uint) {
	lz := uint(bits.LeadingZeros64(v)) & 0x3f // &63 to eliminate bounds checking
	r = (0x3f - lz) * 3
	r += (uint(v-lookup64[lz])>>0x3f)*3 ^ 3
	r *= 6554
	r >>= 16
	return
}

func FastUint32(v uint32) (r uint) {
	lz := uint(bits.LeadingZeros32(v)) & 0x1f // &31 to eliminate bounds checking
	r = (0x1f - lz) * 3
	r += (uint(v-lookup32[lz])>>0x1f)*3 ^ 3
	r *= 6554
	r >>= 16
	return
}
