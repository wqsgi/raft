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
	currentTerm   int32
	rpc           RpcServer
	receiveEvents chan *CommandEvent
	sendEvents    chan *CommandEvent
	role          Role
	stop chan bool

}

type CommandEvent struct {
	target    interface{}
	eventType CommandType
	err       error
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

func (r *Raft) loopLeader() error {




}

func (r *Raft) loopFollower() error {
	t := time.NewTimer(time.Second * randomTime())

	for {
		select {
		case event := <-r.receiveEvents:

			switch event.eventType{
			case APPEND_LOG:

			}
		case <-t.C:
			t.Reset(time.Second*randomTime())
			r.receiveEvents <- &CommandEvent{&VoteRequest{}, eventType:VOTE_REP}
		}

	}

}
func (r *Raft) loopCandidate() error {
	t := time.NewTimer(time.Second * randomTime())

	for {
		select {
		case event := <-r.receiveEvents:

			switch event.eventType{
			case VOTE:
				var vote VoteRequest = event.target;
				if r.currentTerm>vote.term {
					r.sendEvents <- &CommandEvent{target:&VoteResponse{success:false, term:r.currentTerm}, eventType:VOTE_REP}
				} else {
					r.receiveEvents <- &CommandEvent{&VoteResponse{success:true, term:r.currentTerm}, eventType:VOTE_REP}
					r.role = FOLLOWER
					return nil
				}
			}
		case <-t.C:

			// 如果长时间得不到应答。
			t.Reset(time.Second*randomTime())
			r.receiveEvents <- &CommandEvent{&VoteRequest{}, eventType:VOTE_REP}
		}

	}
}

func (r *Raft) processVote(req *VoteRequest) {

}



func randomTime() int32 {

	return rand.Int31n(5)

}








