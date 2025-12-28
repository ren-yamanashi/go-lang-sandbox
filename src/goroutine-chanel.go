package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNumChanel(c chan<- int) {
	fmt.Println("...")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)

	// num をチャネルに送信
	c <- num
}

func _goroutineChanel() {
	fmt.Println("what is today's lucky number?")
	// チャネルの作成
	c := make(chan int)
	go getLuckyNumChanel(c)

	// チャネルから値を受信
	num := <-c
	fmt.Printf("Today's your lucky number is %d!\n", num)

	close(c)
}
