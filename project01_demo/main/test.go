package main

import (
	"fmt"
)

var a []int = make([]int, 10)

func main() {
	for i := range a {
		fmt.Println(i)
	}
}
