package bloom

type BitArray struct {
	bitSize     uint64
	bitCapacity uint64
	segment     []uint64
}

func div64Ceil(nbit uint64) uint64 {
	nseg := nbit >> 6
	resnbit := nbit - nseg<<6
	nseg += resnbit & 1
	return nseg
}

func mod64(i uint64) uint64 {
	return i - (i>>6)<<6
}

func newBitArray(nbit uint64) BitArray {
	return BitArray{
		bitCapacity: nbit,
		segment:     make([]uint64, div64Ceil(nbit)),
	}
}

func (b *BitArray) Size() uint64 {
	return b.bitSize
}

func (b *BitArray) Capacity() uint64 {
	return b.bitCapacity
}

func (b *BitArray) Set(index uint64) {
	segidx := index >> 6
	shift := mod64(index)
	mask := uint64(1) << shift
	b.segment[segidx] |= mask
	b.bitSize += (mask & b.segment[segidx]) >> shift
}

func (b *BitArray) Get(index uint64) bool {
	return (b.segment[index>>6] & (uint64(1) << mod64(index))) != 0
}
