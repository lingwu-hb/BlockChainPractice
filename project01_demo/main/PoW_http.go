package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//区块的结构
type Block struct {
	PreHash    string //上一个区块的哈希值
	MerkleHash string //通过此区块中的数据生成的哈希值
	TimeStamp  string //时间戳
	BMP        int    //难度值
	Data       string //区块中的交易信息
	Nonce      int    //随机数s
}

var BlockChain []Block

//you need to realize pow and blockchain first
/*add your code here*/
var mutex = &sync.Mutex{}

func main() {
	//load .env file and listen port 9000
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	//设置一个通道变量，用于阻塞通道，当用户输入完成之后，再开启一个新的协程
	waitPut := make(chan int)
	/*add your code here*/
	go func() {

		//genesis block
		/*add your code here*/
		a := GenerateFirstBlock("hanbo->xiongyuhan 100yuan")
		mutex.Lock()
		//add the genesis block to blockchain
		BlockChain = append(BlockChain, a)
		for {
			var pos int
			fmt.Println("输入0以退出循环，输入非零输入交易信息")
			fmt.Scanln(&pos)
			if pos == 0 {
				break
			}
			var data string
			fmt.Println("请输入交易信息")

			inputReader := bufio.NewReader(os.Stdin)
			data, err := inputReader.ReadString('\n')
			if err != nil {
				log.Fatal("some error here!")
			}
			b := GenerateNextBlock(data, a, 5)
			BlockChain = append(BlockChain, b)
		}
		/*add your code here*/
		mutex.Unlock()
		//formatted output the genesis block
		/*add your code here*/
		//向主协程发送数据，告知主协程用户已经输入完毕
		waitPut <- 0
	}()
	<-waitPut
	//HTTP
	log.Fatal(httpStart()) //简单理解为将httpStart()里面的内容输出，然后程序会退出
}

//HTTP
func httpStart() error {
	//callback for processing Get request
	mux := makeMuxRouter()
	httpAddr := os.Getenv("PORT") //获得设置好的环境变量——端口号9000
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:    ":" + httpAddr,
		Handler: mux,
		//timeout settings
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		//Set the size of the data included in the request
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

//callback
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()                                //返回一个Router实例
	muxRouter.HandleFunc("/", handGetBlockchain).Methods("GET") //可以理解为为get方法创建一套规则，可以跳转到函数handGetBlockchain
	muxRouter.HandleFunc("/", handPOSTBlockchain).Methods("POST")
	return muxRouter
}

type PostData struct {
	Bmp  int    `json:"BMP"`
	Data string `json:"DATA"`
}

//HTTP POST
func handPOSTBlockchain(w http.ResponseWriter, r *http.Request) {
	//通过 Request 里面的主体内容，更新一个新的区块到 BlockChain 中
	var postdata PostData
	lent := r.ContentLength
	body := make([]byte, lent)
	r.Body.Read(body)
	json.Unmarshal(body, &postdata) //解码数据之后保存到 postdata 中
	// fmt.Fprintf(w, "%#v\n", string(body)) // for test
	// fmt.Fprintf(w, "%#v\n", postdata)
	temp := GenerateNextBlock(postdata.Data, BlockChain[len(BlockChain)-1], postdata.Bmp)
	BlockChain = append(BlockChain, temp)
	//request with json format
	bytes, err := json.MarshalIndent(BlockChain, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

//HTTP GET
func handGetBlockchain(w http.ResponseWriter, r *http.Request) {
	//request with json format
	bytes, err := json.MarshalIndent(BlockChain, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
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

	firstblock.BMP = 5 //创世区块难度值
	firstblock.Nonce = 0
	firstblock.TimeStamp = time.Now().Format("2006-01-02 15:04:05") //格式化
	return firstblock
}

//计算区块的哈希值
func GenerationHashValue(block Block) string {
	//请仿照比特币的计算方式计算哈希
	var hashdata string
	/*add your code here*/
	hashdata = block.PreHash + block.MerkleHash + block.TimeStamp + strconv.Itoa(block.BMP) + block.Data + strconv.Itoa(block.Nonce)
	var sha = sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

//通过PoW共识算法生成新区块
func GenerateNextBlock(data string, oldBlock Block, bmp int) Block {
	var newBlock Block
	/*add your code here*/
	newBlock.PreHash = GenerationHashValue(oldBlock) //先算出上个区块的哈希，然后再赋值
	newBlock.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	newBlock.BMP = bmp
	newBlock.Data = data
	newBlock.MerkleHash = pow(oldBlock.BMP, &newBlock)
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
		for i = 0; i < block.BMP; i++ {
			if hash[i] != '0' {
				break
			}
		}
		if i == block.BMP {
			return hash
		}
		block.Nonce++
	}
}
