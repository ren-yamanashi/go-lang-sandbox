package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getLuckyNum() {
	fmt.Println("...")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	fmt.Printf("Today's your lucky number is %d!\n", num)
}

func _goroutine() {
	fmt.Println("what is today's lucky number?")

	var wg sync.WaitGroup
	wg.Add(1)
	
	// `wg.Done` でゴルーチンが終了した時に wg の内部カウンタの値をデクリメントする
	go func() {
		defer wg.Done()
		getLuckyNum()
	}()

	// `wg.Wait` で wg の内部カウンタの値が 0 になるまで待機する
	wg.Wait()
}
