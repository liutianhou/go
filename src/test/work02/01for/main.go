package main

import "fmt"

func main() {

	//当我i==5时跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)

	}
	fmt.Println("over")
	// 当我 i=5 跳出此次循环 继续下次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)

	}
	fmt.Println("over")

}
