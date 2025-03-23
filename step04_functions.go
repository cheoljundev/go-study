package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	//함수 호출
	result := add(3, 4)
	fmt.Printf("3 + 4 = %d\n", result)
}
