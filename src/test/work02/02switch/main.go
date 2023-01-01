package main

import "fmt"

func main() {
	//switch 简化大量的判断
	var n = 5
	switch n {
	case 1:
		fmt.Println("1111")
	case 2:
		fmt.Println("222")
	case 3:
		fmt.Println("333")
	case 4:
		fmt.Println("444")
	case 5:
		fmt.Println("555")

	default:
		fmt.Println("default")
	}
}
