package wei
import (
	"wei"
	"fmt"
	"time"
	"math/rand"
)

type Role int

// 节点状态
const (
	FOLLOWER Role = iota
	CANDIDATE
	LEADER
)

type Raft struct {
	currentTerm int32
	rpc         RpcServer
	events chan *CommandEvent
	role Role
}

type CommandEvent struct{
	target interface{}
	returnValue interface{}
	eventType CommandType
	err error
}


type Server interface {
	Start() error

	Stop() error
}


func NewRaft() (*Raft, error) {

	raft := &Raft{currentTerm:0, rpc:NewRpcServer() }

	return raft, nil
}


func (r *Raft) Start() error {




}

func (r *Raft) loopLeader()error{




}

func (r *Raft) loopFollower()error{

	t := time.NewTimer(time.Second * randomTime())

	for {
		select {
		case event := <-r.events:
		case <-t.C:
			t.Reset(time.Second*randomTime())
			r.sendVote<-&VoteRequest{}
		}

	}


}
func (r *Raft) loopCandidate() error {
	t := time.NewTimer(time.Second * randomTime())

	for {
		select {
		case vote := <-r.events:
			if r.currentTerm>vote.term{
				r.sendVoteAck<-&VoteResponse{success:false,term:r.currentTerm}
			} else {
				r.sendVoteAck<-&VoteResponse{success:true,term:r.currentTerm}
				r.role = FOLLOWER
				return nil
			}
		case <-t.C:
			t.Reset(time.Second*randomTime())
			r.sendVote<-&VoteRequest{}
		}

	}
}



func randomTime() int32 {

	return rand.Int31n(5)

}








