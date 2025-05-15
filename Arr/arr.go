package main

import "fmt"

func main() {
	// 배열 선언 및 초기화
	var arr1 [5]int // 정수형 배열 선언
	arr2 := [3]string{"안녕하세요", "세계", "고랭"} // 문자열 배열 선언 및 초기화

	// 배열 값 설정
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50

	// 배열 출력
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)

	// 배열 순회
	fmt.Println("arr1 순회:")
	for i, v := range arr1 {
		fmt.Printf("인덱스 %d: 값 %d\n", i, v)
	}
}