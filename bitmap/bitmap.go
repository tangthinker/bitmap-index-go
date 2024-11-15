package bitmap

import (
	"fmt"
	"github.com/tangthinker/bitmap-index-go/bitop"
	"strconv"
	"strings"
)

type Bitmap struct {
	Data []uint64
}

func NewBitmap() *Bitmap {
	return &Bitmap{
		Data: make([]uint64, 0),
	}
}

func (b *Bitmap) String() string {
	var builder strings.Builder
	builder.Grow(len(b.Data) * 16)
	for _, d := range b.Data {
		strD := strconv.FormatUint(d, 16)
		strD = fmt.Sprintf("%016s", strD)
		builder.WriteString(strD)
	}
	return builder.String()
}

func ToBitmap(data string) *Bitmap {
	var bitmap Bitmap
	bitmap.Data = make([]uint64, len(data)/16)

	for i := 0; i < len(bitmap.Data); i++ {
		value, _ := strconv.ParseUint(data[i*16:i*16+16], 16, 64)
		bitmap.Data[i] = value
	}

	return &bitmap
}

func (b *Bitmap) SetBit(index int) {
	b.Data = bitop.SetBit2Data(b.Data, index)
}

func (b *Bitmap) SetBits(index ...int) {
	for _, idx := range index {
		b.SetBit(idx)
	}
}

func (b *Bitmap) ClearBit(index int) {
	b.Data = bitop.ClearBit2Data(b.Data, index)
}

func (b *Bitmap) ClearBits(index ...int) {
	for _, idx := range index {
		b.ClearBit(idx)
	}
}

func (b *Bitmap) Traverse(callback func(index int)) {
	bitop.TraverseData(b.Data, callback)
}

func (b *Bitmap) TargetIds() []int {
	var ids []int
	b.Traverse(func(index int) {
		ids = append(ids, index)
	})
	return ids
}

func (b *Bitmap) Or(b2 *Bitmap) *Bitmap {
	bm := NewBitmap()
	bm.Data = bitop.OR(b.Data, b2.Data)
	return bm
}

func (b *Bitmap) And(b2 *Bitmap) *Bitmap {
	bm := NewBitmap()
	bm.Data = bitop.AND(b.Data, b2.Data)
	return bm
}

func (b *Bitmap) Xor(b2 *Bitmap) *Bitmap {
	bm := NewBitmap()
	bm.Data = bitop.XOR(b.Data, b2.Data)
	return bm
}

func (b *Bitmap) Not() *Bitmap {
	bm := NewBitmap()
	bm.Data = bitop.NOT(b.Data)
	return bm
}
