package nodes


type NodeManager interface {
	sendVote()(bool,error)
}




