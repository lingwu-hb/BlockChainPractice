package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

var users []User

//用户结构体
type User struct {
	coinNumber int64
	address    string
	_durtime   time.Time
	weight     int64 //用户权益值
}

// ComputeWeight 用户计算用户的权益值，并将其赋值为Weight
func (a User) ComputeWeight() {
	tmp := time.Now().Sub(a._durtime).Microseconds() //得到持币时间（用毫秒表示）
	a.weight = tmp * a.coinNumber
}

// Initialize 实现User的初始化
func (a User) Initialize() {
	//用户可以在这里指定初始信息
}

func ReturnIndex(users []User) int64 {
	for _, user := range users {
		user.ComputeWeight() //为每一个用户计算权益值
	}
	//随机数算法返回幸运的打包者
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

//创世区块
func GenerateFirstBlock(data string) Block {
	var firstblock Block
	/*add your code here*/
	firstblock.PreHash = "0"
	firstblock.Data = data

	var sha = sha256.New()
	sha.Write([]byte(firstblock.Data))
	hashed := sha.Sum(nil)
	firstblock.MerkleHash = hex.EncodeToString(hashed)

	firstblock.TimeStamp = time.Now().Format("2006-01-02 15:04:05") //格式化
	return firstblock
}

//计算区块的哈希值
func GenerationHashValue(block Block) string {
	//请仿照比特币的计算方式计算哈希
	var hashdata string
	/*add your code here*/
	hashdata = block.PreHash + block.MerkleHash + block.TimeStamp + strconv.Itoa(block.Diff) + block.Data + strconv.Itoa(block.Nonce)
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

	hashdata := string(newBlock.Index) + newBlock.Validator + newBlock.PreHash + newBlock.TimeStamp + newBlock.Data
	var sha = sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)
	newBlock.MerkleHash = hex.EncodeToString(hashed)
}

func Print_(block Block) {
	fmt.Println(BlockList)
}

var BlockList []Block

//main函数
func main() {
	//初始化每一个用户的前两个数据元素

}
