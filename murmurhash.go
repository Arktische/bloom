package bloom

import "math/bits"

const (
	murmur128m0 = 0x87c37b91114253d5
	murmur128m1 = 0x4cf5ad432745937f
)

// fmix64 finalization mix, force all bits to avalanche
func fmix64(k uint64) uint64 {
	k ^= k >> 33
	k *= 0xff51afd7ed558ccd
	k ^= k >> 33
	k *= 0xc4ceb9fe1a85ec53
	k ^= k >> 33
	return k
}

// murmurhash128
func murmurhash128(k0 uint64, k1 uint64, seed uint) (uint64, uint64) {
	h0, h1 := uint64(seed), uint64(seed)

	k0 *= murmur128m0
	k0 = bits.RotateLeft64(k0, 31)
	k0 *= murmur128m1
	h0 ^= k0

	h0 = bits.RotateLeft64(h0, 27)
	h0 += h1
	h0 = h0*5 + 0x52dce729

	k1 *= murmur128m1
	k1 = bits.RotateLeft64(k1, 33)
	k1 *= murmur128m0
	h1 ^= k1

	h1 = bits.RotateLeft64(h1, 31)
	h1 += h0
	h1 = h1*5 + 0x38495ab5

	h0 ^= 16
	h1 ^= 16

	h0 += h1
	h1 += h0

	h0 = fmix64(h0)
	h1 = fmix64(h1)

	h0 += h1
	h1 += h0
	return h0, h1
}

const (
	murmur32m0 = 0xcc9e2d51
	murmur32m1 = 0x1b873593
)

func murmurhash32(k uint32, seed uint) uint32 {
	h := uint32(seed)
	k *= murmur32m0
	k = bits.RotateLeft32(k, 15)
	k *= murmur32m1
	h ^= k
	h = bits.RotateLeft32(h, 13)
	h = 5*h + 0xe6546b64
	h ^= 32

	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}
