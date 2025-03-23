package main

import (
	"fmt"
	"sync"
	"time"
)

func workerStep12(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("작업자 #%d: 작업 대기 중...\n", id)
	for job := range jobs {
		fmt.Printf("작업자 #%d: 작업 #%d 수신\n", id, job)
		fmt.Printf("작업자 #%d: 작업 #%d 처리 중...\n", id, job)
		time.Sleep(1 * time.Second)
		fmt.Printf("작업자 #%d: 작업 #%d 완료!\n", id, job)
	}
}

func main() {
	const workerCount = 3
	const jobCount = 5

	jobs := make(chan int, jobCount)
	var wg sync.WaitGroup

	//작업자 고루틴 실행
	//고루틴이 경쟁적으로 채널에서 데이터를 가져가게 된다.
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go workerStep12(w, jobs, &wg)
	}

	time.Sleep(100 * time.Millisecond) // 💡 고루틴이 실행될 틈을 줌

	//작업 전송
	for j := 1; j <= jobCount; j++ {
		jobs <- j
		fmt.Printf("작업 #%d 전송 완료\n", j)
	}

	close(jobs) //더 이상 작업 없음

	wg.Wait()
	fmt.Println("모든 작업 완료!")
}
