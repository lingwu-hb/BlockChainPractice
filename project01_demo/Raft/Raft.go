// Just need to complete the election of three Nodes
// Not consider log replication

package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"time"
)

type State int

//枚举类型——状态类型
const (
	Follower State = iota
	Candidate
	Leader
	Dead
)

//将枚举类型设置字符串
func (s State) String() string {
	switch s {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	default:
		panic("unreachable")
	}
}

type ConsensusModule struct {
	id                int       //该节点的编号
	voteFor           int       //投票给谁
	currentTerm       int       //当前任期
	state             State     //发出拉票的结点的状态
	electionResetTime time.Time //根据该时间来重启选举过程
}

var peers []ConsensusModule //计算机集群上的所有计算机的编号

func (n *ConsensusModule) RequestVote(args RequestVoteArgs, reply *RequestVoteReply) error {
	//如果收到的请求的任期大于自己的，那么自己直接变成Follower
	if args.Term > n.currentTerm {
		n.becomeFollower(args.Term)
	}

	if n.currentTerm == args.Term &&
		(n.voteFor == -1 || n.voteFor == args.CandidateId) {
		reply.VoteGranted = true
		n.voteFor = args.CandidateId
		n.electionResetTime = time.Now()
	} else {
		reply.VoteGranted = false
	}
	reply.Term = n.currentTerm
	return nil
}

func NewNode(Id int) *ConsensusModule {
	node := new(ConsensusModule)
	node.id = Id
	node.voteFor = -1
	node.state = Follower
	node.electionResetTime = time.Now()
	return node
}

//根据结点的选举时间来返回一个随机时间
func (n *ConsensusModule) runElectionTimeOut() time.Duration {
	return time.Duration(150+rand.Intn(150)) * time.Millisecond
}

//用于通信的结构体
type RequestVoteArgs struct {
	Term        int
	CandidateId int
}

type RequestVoteReply struct {
	Term        int
	VoteGranted bool
}

func (n *ConsensusModule) startElection() {
	//n这个服务器的共识模块需要开始进入选举模式
	n.state = Candidate
	n.currentTerm += 1
	savedCurrentTerm := n.currentTerm //当前任期更新为n的任期
	n.electionResetTime = time.Now()
	n.voteFor = n.id //为自己投票

	votesReceived := 1

	for i := 1; i >= len(peers); i++ {
		go func() {
			//设置参与选举的通信rpc参数
			args := RequestVoteArgs{
				Term:        savedCurrentTerm,
				CandidateId: n.id,
			}
			var reply RequestVoteReply

			rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
			if err != nil {
				log.Fatal(err)
			}
			//按照rpc通信规则进行通信，并请求处理，处理后的值返回到reply结构体中
			err2 := rp.Call("ConsensusModule.RequestVote", args, &reply)
			if err2 != nil {
				log.Fatal(err2)
			}

			//如果返回的任期比现在的任期还新，说明该结点没资格参与选举
			if reply.Term > savedCurrentTerm {
				//该节点根据返回的最新任期改变状态为追随者
				n.becomeFollower(reply.Term)
				return
			}
			//而如果返回的任期与现任期相同，而且同意投票，则该节点加一票
			if reply.Term == savedCurrentTerm && reply.VoteGranted {
				votesReceived += 1
				//获取票数超过半数，则当选为Leader
				if votesReceived*2 > len(peers)+1 {
					n.startLeader()
					return
				}
			}
		}()
	}

	go n.runElectionTimer()
}

func (n *ConsensusModule) becomeFollower(Term int) {
	n.state = Follower
	n.currentTerm = Term
	n.voteFor = -1
	n.electionResetTime = time.Now()
	//重新进入计时器倒计时阶段
	go n.runElectionTimer()
}

func (n *ConsensusModule) startLeader() {
	n.state = Leader
	//后续日志复制部分内容
}

//开启一个选举周期
func (n *ConsensusModule) runElectionTimer() {
	timeoutDuration := n.runElectionTimeOut()
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for {
		<-ticker.C
		//如果我们本身就是Leader，那么直接退出即可
		if n.state == Leader {
			return
		}
		//如果触发了超时，那么就要开始进入选举过程
		if elapsed := time.Since(n.electionResetTime); elapsed >= timeoutDuration {
			n.startElection()
			return
		}
	}
}

func PeersListen() {
	cm := new(ConsensusModule)
	rpc.Register(cm)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//initialize peerIds
	var NodeNum int = 3
	for i := 1; i <= NodeNum; i++ {
		peers = append(peers, *NewNode(i))
	}
	//让所有服务器都监听对应端口号
	go PeersListen()
	for _, peer := range peers {
		//所有节点同时开始选举周期
		go peer.runElectionTimer()
	}
}
