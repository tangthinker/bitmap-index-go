package bitmap

import (
	"fmt"
	"testing"
)

func TestToBitmap(t *testing.T) {

	bm := NewBitmap()

	bm.SetBit(33)

	bm.SetBit(99)

	bm.SetBit(63)

	bm.SetBit(30043)

	fmt.Println("bm.data: ", bm.Data)

	fmt.Println("===============")
	bm.Traverse(func(index int) {
		fmt.Println(index)
	})
	fmt.Println("===============")

	bmStr := bm.String()

	fmt.Println("bm string: ", bm.String())

	bm2 := ToBitmap(bmStr)
	fmt.Println("===============")
	bm2.Traverse(func(index int) {
		fmt.Println(index)
	})
	fmt.Println("===============")

	fmt.Println("bm2.data: ", bm2.Data)

}
