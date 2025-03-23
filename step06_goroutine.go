package main

import (
	"fmt"
	"time"
)

// ch라는 string 타입 채널을 매개변수로 받는다
func sayHello(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Hello from goroutine!"
}

func main() {
	//채널 생성
	messageChannel := make(chan string)

	//고루틴 실행(비동기)
	go sayHello(messageChannel)

	//채널로부터 메시지 수신(여기서 대기함)
	msg := <-messageChannel

	fmt.Println("메시지 받음:", msg)
}
