package main

import "fmt"

func main() {
	age := 20

	//조건문
	if age >= 18 {
		fmt.Println("성인입니다.")
	} else {
		fmt.Println("미성년자입니다.")
	}

	//반복문
	fmt.Println("1부터 5까지 출력:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
