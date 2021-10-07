package main

import "fmt"

type CMState int

const ( //Go语言中枚举类型的声明
	Follower CMState = iota
	Candidate
	Leader
	Dead
)

//实现String()方法之后，就可以按照字符串的格式进行输出了（很方便）
func (s CMState) String() string {
	switch s {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	case Dead:
		return "Dead"
	default:
		panic("unreachable")
	}
}

func main() {
	fmt.Println(Dead)
}

//输出结果为：Dead
