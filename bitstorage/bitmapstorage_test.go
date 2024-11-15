package bitstorage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestRedisBitStorage(t *testing.T) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "",
	})

	bitmapStorage := NewRedisBitStorage(redisClient)

	qKey := "q-test"
	wKey := "w-test"

	if err := bitmapStorage.SetBits(context.Background(), qKey, 0, 1, 2, 3, 1000, 10000, 1000000, 999, 999999, 1000000); err != nil {
		t.Errorf("set bits error: %v", err)
		return
	}

	if err := bitmapStorage.SetBits(context.Background(), wKey, 0, 1, 3, 1000, 10000, 1000000, 999, 1000000); err != nil {
		t.Errorf("set bits error: %v", err)
		return
	}

	fmt.Println("=======================")
	err := bitmapStorage.Traverse(context.Background(), qKey, func(index int) {
		fmt.Println(index)
	})
	if err != nil {
		t.Errorf("traverse error: %v", err)
		return
	}
	fmt.Println("=======================")

	qBm, err := bitmapStorage.Bitmap(context.Background(), qKey)
	if err != nil {
		t.Errorf("get bitmap error: %v", err)
		return
	}

	wBm, err := bitmapStorage.Bitmap(context.Background(), wKey)
	if err != nil {
		t.Errorf("get bitmap error: %v", err)
		return
	}

	andBm := qBm.And(wBm)
	fmt.Println("==================================")
	andBm.Traverse(func(index int) {
		fmt.Println(index)
	})
	fmt.Println("==================================")

}
