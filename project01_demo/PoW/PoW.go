package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

//区块的结构
type Block struct {
	PreHash    string //上一个区块的哈希值
	MerkleHash string //通过此区块中的数据生成的哈希值
	TimeStamp  string //时间戳
	Diff       int    //难度值
	Data       string //区块中的交易信息
	Nonce      int    //随机数
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

	firstblock.Diff = 0
	firstblock.Nonce = 0
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

//通过PoW共识算法生成新区块
func GenerateNextBlock(data string, oldBlock Block) Block {
	var newBlock Block
	/*add your code here*/
	newBlock.PreHash = GenerationHashValue(oldBlock) //先算出上个区块的哈希，然后再赋值
	newBlock.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	newBlock.Diff = oldBlock.Diff
	newBlock.Data = data
	newBlock.MerkleHash = pow(oldBlock.Diff, &newBlock)
	return newBlock
}

//实现简易PoW算法
func pow(diff int, block *Block) string {
	for {
		var hash string
		//再循环内写出判断挖矿是否成功的依据
		/*add your code here*/
		hash = GenerationHashValue(*block)
		//不断尝试，直到满足难度值为止
		var i int
		for i = 0; i < block.Diff; i++ {
			if hash[i] != '0' {
				break
			}
		}
		if i == block.Diff {
			return hash
		}
		block.Nonce++
	}
}

func Print_(block Block) {
	fmt.Println(ListOfBlock)
}

var ListOfBlock []Block //some error here, wait to deal

//main函数
func main() {
	//生成创世区块
	myblock := GenerateFirstBlock("😓 -> 🐻 5元")
	ListOfBlock = append(ListOfBlock, myblock)
	//如何利用前后两个区块之间的联系来达到访问的区别
	Print_(myblock)
	ListOfBlock = append(ListOfBlock, GenerateNextBlock("🐻 -> 😓 10元", ListOfBlock[0]))
	Print_(myblock)
}
