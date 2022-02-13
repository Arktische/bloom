package bloom

import (
	"encoding/binary"
	"unsafe"
)

type BloomFilter struct {
	bitArray *BitArray
}

func (b *BloomFilter) Get() {

}

var put = []func(bf *BloomFilter, key []byte){
	putKey32bit, putKey32bit, putKey32bit, putKey32bit,
	putKey64bit, putKey64bit, putKey64bit, putKey64bit,
	putKey128bit, putKey128bit, putKey128bit, putKey128bit,
	putKey128bit, putKey128bit, putKey128bit, putKey128bit,
	putKeyLongerThan128bit,
}

func putKey64bit(bf *BloomFilter, key []byte) {

}

func putKey32bit(bf *BloomFilter, key []byte) {
	murmurhash32(*(*uint32)(unsafe.Pointer(&key[0])), 0xbc9f1d34)
}

func putKey128bit(bf *BloomFilter, key []byte) {
	murmurhash128(binary.LittleEndian.Uint64(key[:8]),
		binary.LittleEndian.Uint64(key[8:]), 0xbc9f1d34)
}

func putKeyLongerThan128bit(bf *BloomFilter, key []byte) {
}

func (b *BloomFilter) Put(key []byte) {
	put[len(key)](b, key)
}

func (b *BloomFilter) PutUint32(key uint32) {

}

func (b *BloomFilter) PutUin64(key uint64) {

}
