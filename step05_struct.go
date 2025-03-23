package main

import "fmt"

//구조체 정의

type User struct {
	Name string
	Age  int
}

func main() {
	//구조체 변수 생성
	u := User{
		Name: "Cheoljun",
		Age:  31,
	}

	//출력
	fmt.Printf("이름: %s\n", u.Name)
	fmt.Printf("나이: %d\n", u.Age)
}
