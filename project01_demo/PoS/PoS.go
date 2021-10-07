package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

var users []User
var userNum int

//用户结构体
type User struct {
	coinNumber int64
	address    string //用户的地址，用于身份验证
	_durtime   time.Time
	weight     int64 //用户权益值
}

// ComputeWeight 用户计算用户的权益值，并将其赋值为Weight
func (a User) ComputeWeight() {
	tmp := time.Since(a._durtime).Microseconds() //得到持币时间（用毫秒表示）
	a.weight = tmp * a.coinNumber
}

// Initialize 实现User的初始化
func Initialize() {
	//用户可以在这里指定初始信息
	fmt.Println("请输入用户个数")
	var n int
	fmt.Scanln(&n)
	i := 1
	for ; i <= n; i++ {
		var (
			coinNumber int64
			address    string
		)
		fmt.Println("请输入该用户的持币数，以及用户地址（随便一个字符串就行）")
		fmt.Scanln(&coinNumber, &address)
		var temp User
		temp.coinNumber = coinNumber
		temp.address = address
		temp._durtime = time.Now()
		temp.ComputeWeight()
		users = append(users, temp)
	}
	userNum = i - 1
}

// ReturnIndex 利用随机数返回一个幸运矿工的序号
func ReturnIndex() int64 {
	for _, user := range users {
		user.ComputeWeight() //为每一个用户计算权益值
	}
	//随机数算法返回幸运的打包者
	//计算总权重
	var totalWeight int64
	for i := 0; i < userNum; i++ {
		totalWeight += users[i].weight
	}
	var portion []float64 = make([]float64, userNum+1) //记录每个用户的比例（portion）
	for i := 1; i <= userNum; i++ {
		// 将每个用户所占的比例计算出来，然后再加上之前的用户的比例（最终所有用户的比例会占满[0, 1]）
		portion[i] = float64(users[i-1].weight)/float64(totalWeight) + portion[i-1]
	}
	rand.Seed(time.Now().Unix())
	var x = rand.Float64()
	for i := 1; i <= userNum; i++ {
		if x > portion[i-1] && x < portion[i] {
			return int64(i)
		}
	}
	return 0
}

//区块的结构
type Block struct {
	Index      int    //每个区块在序列中的位置
	Validator  string //每个区块从属于那个矿工（存储矿工的地址）
	PreHash    string //上一个区块的哈希值
	MerkleHash string //通过此区块中的数据生成的哈希值
	TimeStamp  string //时间戳
	Data       string //区块中的交易信息
}

var BlockChain []Block

//创世区块
func GenerateFirstBlock(data string, cut int, validator string) {
	var firstblock Block
	/*add your code here*/
	firstblock.Index = cut
	firstblock.Validator = validator
	firstblock.PreHash = ""                                         //默认初始值，为空
	firstblock.TimeStamp = time.Now().Format("2006-01-02 15:04:05") //格式化
	firstblock.Data = data
	firstblock.MerkleHash = GenerationHashValue(firstblock)
	BlockChain = append(BlockChain, firstblock)
}

//计算区块的哈希值
func GenerationHashValue(block Block) string {
	//请仿照比特币的计算方式计算哈希
	var hashdata string
	/*add your code here*/
	hashdata = string(block.Index) + block.Validator + block.PreHash + block.TimeStamp + block.Data
	var sha = sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

//生成新区块——validator由main函数传递
func GenerateNextBlock(data string, oldBlock Block, cut int, validator string) {
	var newBlock Block
	/*add your code here*/
	newBlock.Index = cut
	newBlock.Validator = validator
	newBlock.PreHash = GenerationHashValue(oldBlock) //先算出上个区块的哈希，然后再赋值
	newBlock.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	newBlock.Data = data
	newBlock.MerkleHash = GenerationHashValue(newBlock)
	BlockChain = append(BlockChain, newBlock)
}

func Print_() {
	//fmt.Println(users)
	fmt.Println(BlockChain)
}

//main函数
func main() {
	//初始化每一个用户的前两个数据元素
	Initialize()
	//Print_()
	cut := 0
	validator := ReturnIndex()
	GenerateFirstBlock("hanbo->xiongyuhan 100bitcoin", cut, users[validator].address)
	for {
		var pos int
		fmt.Println("输入0以退出循环")
		fmt.Scanln(&pos)
		if pos == 0 {
			break
		}
		var data string
		fmt.Println("请输入交易信息")
		fmt.Scanln(&data)
		cut++
		validator := ReturnIndex()
		GenerateNextBlock(data, BlockChain[cut-1], cut, users[validator].address)
	}
	Print_()
}
