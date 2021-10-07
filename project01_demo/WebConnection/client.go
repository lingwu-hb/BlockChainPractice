package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct { //结构体里面的内容就是rpc通信过程中传递的内容！
	Width, Height int
}

//rpc client
func main() {
	rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	//set result -> 0
	ret := 0
	err2 := rp.Call("Rect.Area", Params{20, 30}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("area: ", ret)
	err3 := rp.Call("Rect.Perimeter", Params{20, 30}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("Perimeter: ", ret)
}
