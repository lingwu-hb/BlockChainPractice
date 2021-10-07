package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//Go suports three levels of RPC: TCP、HTTP、JSONRPC
//The first letter must be capitalized
type Params struct {
	Width, Height int
}
type Rect struct{}

//compute area
func (r *Rect) Area(p Params, ret *int) error { //注册一个满足rpc规则的方法
	*ret = p.Width * p.Height
	return nil
}

//compute perimeter
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}
func main() {
	//register service with http
	rect := new(Rect)
	rpc.Register(rect)
	rpc.HandleHTTP()
	//monitoring service
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
