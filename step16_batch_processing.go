package main

import (
	"fmt"
	"sync"
	"time"
)

func workerStep16(batch []int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("배치 처리 시작: %v\n", batch)
	time.Sleep(2 * time.Second) // 작업 시간 시뮬레이션
	results <- fmt.Sprintf("배치 완료 : %v", batch)
}

func main() {
	totalJobs := 20
	batchSize := 5

	jobs := make([]int, totalJobs)
	for i := 0; i < totalJobs; i++ {
		jobs[i] = i + 1
	}

	batchCount := (totalJobs + batchSize - 1) / batchSize //배치사이즈 올림계산
	results := make(chan string, batchCount)
	var wg sync.WaitGroup

	//배치 단위로 작업 나누기
	for i := 0; i < totalJobs; i += batchSize {
		end := i + batchSize
		if end > totalJobs {
			end = totalJobs
		}

		batch := jobs[i:end]
		wg.Add(1)
		go workerStep16(batch, results, &wg)
	}

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("모든 배치 작업 완료")
}
