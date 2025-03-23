package main

import (
	"fmt"
	"time"
)

func sendLate(ch chan string) {
	//일부러 3초 후에 메시지를 보냄
	time.Sleep(3 * time.Second)
	ch <- "응답 도착!"
}

func main() {
	ch := make(chan string)

	go sendLate(ch)

	select {
	case msg := <-ch:
		fmt.Println("수신된 메시지:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("타임 아웃 : 2초 내 응답 없음")
	}
}
