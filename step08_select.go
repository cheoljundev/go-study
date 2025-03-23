package main

import (
	"fmt"
	"time"
)

func sendAfter(ch chan string, msg string, delay time.Duration) {
	time.Sleep(delay)
	ch <- msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendAfter(ch1, "첫 번째 메시지", 2*time.Second)
	go sendAfter(ch2, "두 번째 메시지", 1*time.Second)

	//select 문으로 먼저 도착한 메시지를 받음
	select {
	case msg := <-ch1:
		fmt.Println("ch1에서 수신:", msg)
	case msg := <-ch2:
		fmt.Println("ch2에서 수신:", msg)
	}
}
