package main

import "fmt"

func main() {
	//버퍼 크기가 2인 채널 생성
	ch := make(chan string, 2)

	//고루틴 없이도 연속으로 보낼 수 있음
	ch <- "message 1"
	ch <- "message 2"

	fmt.Println("버퍼에 메시지 2개 저장 완료!")

	// 채넣에서 순차적으로 값 꺼내기
	msg1 := <-ch
	msg2 := <-ch

	fmt.Println("받은 메시지1:", msg1)
	fmt.Println("받은 메시지2:", msg2)
}
