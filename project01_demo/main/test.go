package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hanbpo")
	a := time.Now()
	var c float64
	c = 25.6
	b := time.Now().Sub(a).Microseconds()
	d := float64(b) * c
	fmt.Println(b)
	fmt.Println(d)
}
