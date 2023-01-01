package main

import "fmt"

type WebSite struct {
	Name string
}

func main() {
	site := WebSite{Name: "liu.io"}
	fmt.Printf("site.Name: %v\n", site.Name)
	fmt.Printf("site.Name: %#v\n", site.Name)

	x := 100
	p := &x
	fmt.Printf("p: %#v\n", p)

}
