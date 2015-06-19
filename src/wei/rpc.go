package wei


type CommandType int32

const(
	VOTE CommandType = iota
	VOTE_REP
	APPEND_LOG
	APPEND_LOG_REP
)


type RpcServer interface {
	Vote(req VoteRequest) (VoteResponse,error)
}


type Rpc struct {

}

type VoteRequest struct {

	term int32
	candidate string
	lastLogIndex int32
	lastTerm int32
}

type VoteResponse struct {
	success bool
	// 当前任期
	term int32
}

type AppendLog struct {

}

type AppendLogAck struct {

}
func NewRpcServer()*Rpc{
	return &Rpc{}
}



