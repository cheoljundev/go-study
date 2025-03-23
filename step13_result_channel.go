package main

import (
	"fmt"
	"sync"
	"time"
)

func workerStep13(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("작업자 #%d: 작업 대기 중...\n", id)
	for job := range jobs {
		fmt.Printf("작업자 #%d: 작업 #%d 처리 중...\n", id, job)
		time.Sleep(1 * time.Second)
		result := fmt.Sprintf("작업자 #%d 작업 %d 완료!", id, job) //문자열을 포맷팅해서 리턴
		results <- result
	}
}

func main() {
	const workerCount = 3
	const jobCount = 5

	jobs := make(chan int, jobCount)
	results := make(chan string, jobCount)
	var wg sync.WaitGroup

	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go workerStep13(w, jobs, results, &wg)
	}

	time.Sleep(100 * time.Millisecond) // 💡 고루틴이 실행될 틈을 줌

	for j := 1; j <= jobCount; j++ {
		jobs <- j
		fmt.Printf("작업 #%d 전송 완료\n", j)
	}
	close(jobs)

	//결과 출력
	for i := 1; i <= jobCount; i++ {
		result := <-results
		fmt.Println(result)
	}

	wg.Wait()
	fmt.Println("모든 작업 완료")
}
