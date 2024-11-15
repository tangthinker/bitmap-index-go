package bitop

import (
	"fmt"
	"testing"
)

func TestInt64Arr2ByteArr(t *testing.T) {

	var num uint64 = 1

	byteArr := Uint64Arr2ByteArr([]uint64{num})

	t.Log(byteArr)

}

func TestSetBit(t *testing.T) {

	var num uint64

	//num = SetBit(num, 0)
	//num = SetBit(num, 1)
	//num = SetBit(num, 2)
	//num = SetBit(num, 3)
	// 1111 = 15

	//num = SetBit(num, 0)
	// 1

	num = SetBit(num, 63)

	num = SetBit(num, 32)

	t.Log(num)

	fmt.Println("========================")
	TraverseBit(num, func(index int) {
		fmt.Println(index)
	})

}

func TestSetBit2Data(t *testing.T) {

	var data []uint64

	SetBit2Data(data, 0)

	SetBit2Data(data, 99)

	t.Log(data)

	fmt.Println("========================")
	TraverseData(data, func(index int) {
		fmt.Println(index)
	})

}
