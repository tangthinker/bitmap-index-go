package bitop

func AND(source1, source2 []uint64) []uint64 {
	size := len(source1)
	if len(source2) < size {
		size = len(source2)
	}

	ret := make([]uint64, size)

	for i := 0; i < size; i++ {
		ret[i] = source1[i] & source2[i]
	}

	return ret
}

func OR(source1, source2 []uint64) []uint64 {
	size := len(source1)
	if len(source2) < size {
		size = len(source2)
	}

	ret := make([]uint64, size)

	for i := 0; i < size; i++ {
		ret[i] = source1[i] | source2[i]
	}

	if len(source1) > size {
		ret = append(ret, source1[size:]...)
	}

	if len(source2) > size {
		ret = append(ret, source2[size:]...)
	}

	return ret
}

func NOT(source []uint64) []uint64 {
	ret := make([]uint64, len(source))

	for i := 0; i < len(source); i++ {
		ret[i] = ^source[i]
	}

	return ret
}

func XOR(source1, source2 []uint64) []uint64 {
	size := len(source1)
	if len(source2) < size {
		size = len(source2)
	}

	ret := make([]uint64, size)

	for i := 0; i < size; i++ {
		ret[i] = source1[i] ^ source2[i]
	}

	if len(source1) > size {
		n := NOT(source1[size:])
		ret = append(ret, n...)
	}

	if len(source2) > size {
		n := NOT(source2[size:])
		ret = append(ret, n...)
	}

	return ret
}

func SetBit(source uint64, index int) uint64 {
	return source | (1 << index)
}

func ClearBit(source uint64, index int) uint64 {
	return source &^ (1 << index)
}

func TraverseBit(source uint64, callback func(index int)) {
	for i := 0; i < 64; i++ {
		if source&(1<<i) != 0 {
			callback(i)
		}
	}
}

func SetBit2Data(source []uint64, index int) []uint64 {
	curBitSize := len(source) * 64
	if index >= curBitSize {
		newSize := (index/64 + 1) * 64
		newSource := make([]uint64, newSize/64)
		copy(newSource, source)
		source = newSource
	}
	source[index/64] = SetBit(source[index/64], index%64)
	return source
}

func ClearBit2Data(source []uint64, index int) []uint64 {
	if index >= len(source)*64 {
		return source
	}
	source[index/64] = ClearBit(source[index/64], index%64)
	return source
}

func TraverseData(source []uint64, callback func(index int)) {
	for i := 0; i < len(source)*64; i++ {
		if source[i/64]&(1<<(i%64)) != 0 {
			callback(i)
		}
	}
}

func Uint64Arr2ByteArr(source []uint64) []byte {
	if source == nil || len(source) == 0 {
		return nil
	}

	ret := make([]byte, len(source)*8)

	for i := 0; i < len(source); i++ {
		ret[i*8] = byte(source[i] >> 56)
		ret[i*8+1] = byte(source[i] >> 48)
		ret[i*8+2] = byte(source[i] >> 40)
		ret[i*8+3] = byte(source[i] >> 32)
		ret[i*8+4] = byte(source[i] >> 24)
		ret[i*8+5] = byte(source[i] >> 16)
		ret[i*8+6] = byte(source[i] >> 8)
		ret[i*8+7] = byte(source[i])
	}

	return ret
}

func ByteArr2Uint64Arr(source []byte) []uint64 {
	if source == nil || len(source) == 0 {
		return nil
	}

	size := len(source) / 8
	if len(source)%8 != 0 {
		size++
	}

	ret := make([]uint64, size)

	if len(source)/8 == 0 {
		var num uint64
		for i := 0; i < len(source); i++ {
			num |= uint64(source[i]) << (56 - i*8)
		}
		ret[0] = num
		return ret
	}

	for i := 0; i < len(ret); i++ {
		ret[i] = uint64(source[i*8])<<56 |
			uint64(source[i*8+1])<<48 |
			uint64(source[i*8+2])<<40 |
			uint64(source[i*8+3])<<32 |
			uint64(source[i*8+4])<<24 |
			uint64(source[i*8+5])<<16 |
			uint64(source[i*8+6])<<8 |
			uint64(source[i*8+7])
	}

	return ret
}
