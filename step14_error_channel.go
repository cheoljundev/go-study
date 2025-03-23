package main

import (
	"fmt"
	"sync"
	"time"
)

func workerStep14(id int, jobs <-chan int, results chan<- string, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("작업자 #%d: 작업 대기 중...\n", id)
	for job := range jobs {
		fmt.Printf("작업자 #%d: #%d작업 처리 중...\n", id, job)
		time.Sleep(1 * time.Second)

		//에러 시뮬레이션: 짝수 작업은 실패
		if job%2 == 0 {
			err := fmt.Errorf("작업자 #%d: 작업 %d 실패 (에러 발생)", id, job)
			errors <- err
			continue
		}

		//정상 처리
		result := fmt.Sprintf("작업자 #%d 작업 #%d 완료!", id, job)
		results <- result
	}
}

func main() {
	const workerCount = 3
	const jobCount = 5

	jobs := make(chan int, jobCount)
	results := make(chan string, jobCount)
	errors := make(chan error, jobCount)

	var wg sync.WaitGroup

	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go workerStep14(w, jobs, results, errors, &wg)
	}

	time.Sleep(100 * time.Millisecond) // 💡 고루틴이 실행될 틈을 줌

	for j := 1; j <= jobCount; j++ {
		jobs <- j
		fmt.Printf("작업 #%d 전송 완료\n", j)
	}
	close(jobs)

	//결과 수신
	for i := 1; i <= jobCount; i++ {
		select {
		case res := <-results:
			fmt.Println(res)
		case err := <-errors:
			fmt.Println("에러 발생:", err)
		}
	}

	wg.Wait()
	fmt.Println("모든 작업 완료")
}
