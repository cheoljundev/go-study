package main

import (
	"fmt"
	"sync"
	"time"
)

func workerStep15(job int, limiter chan struct{}, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// 세마포어(리미터) 자리 차지
	limiter <- struct{}{}

	fmt.Printf("작업 #%d 시작\n", job)
	time.Sleep(1 * time.Second)
	results <- fmt.Sprintf("작업 #%d 완료!", job)

	// 자리 반환
	<-limiter
}

func main() {
	const totalJobs = 10
	const maxConcurrency = 3 //동시에 실행될 수 있는 고루틴 수

	results := make(chan string, totalJobs)
	limiter := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup

	for j := 1; j <= totalJobs; j++ {
		wg.Add(1)
		go workerStep15(j, limiter, results, &wg)
	}

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("모든 작업 완료 (동시 제한 적용)")
}
