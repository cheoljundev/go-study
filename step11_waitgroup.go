package main

import (
	"fmt"
	"sync"
	"time"
)

// wg *sync.WaitGroup: WaitGroup 포인터를 받아서, 작업이 끝났을 때 Done() 호출하게 함
func workerStep11(id int, wg *sync.WaitGroup) {
	/**
	* 이 함수가 끝날 때 wg.Done()을 호출해서 작업 완료를 알림
	* wg.Add(1)로 더했던 카운트를 하나 줄여줌
	 */
	defer wg.Done() // 끝나면 -1

	fmt.Printf("작업자 #%d: 작업 시작\n", id)
	time.Sleep(time.Duration(id) * time.Second) //일부러 고루틴들이 다른 속도로 끝나도록 시뮬레이션
	fmt.Printf("작업자 #%d: 작업 완료\n", id)
}

func main() {
	var wg sync.WaitGroup //WaitGroup 생성, 여러 고루틴을 실행하고, 모두 끝날 때까지 기다릴 수 있는 도구

	for i := 1; i <= 3; i++ {
		wg.Add(1)               //작업 추가, 작업 하나 시작 예정이니, WaitGroup에 카운트를 1 추가함
		go workerStep11(i, &wg) //고루틴으로 workerStep11(i, &wg) 실행
	}

	wg.Wait() //모든 고루틴이 끝날때까지 대기
	fmt.Printf("모든 작업자 종료 완료!")
}
