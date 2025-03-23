package main

import "fmt"

// 보내기 전용 채널
func sendMessage(ch chan<- string, msg string) {
	ch <- msg
}

// 받기 전용 채널
func receiveMessage(ch <-chan string) {
	msg := <-ch
	fmt.Println("받은 메시지:", msg)
}

func main() {
	ch := make(chan string)

	//고루틴으로 메시지 보내기
	go sendMessage(ch, "고루틴에서 전송한 메시지")

	//메시지 받기
	receiveMessage(ch)
}
