package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3}
	for _, v := range a {
		fmt.Printf("v: %v\n", v)
	}
}
