package pb

import (
	"encoding/json"
	"time"
)

// 共识节点间的消息结构体
type nodeMsg struct {
	// 目前分为startVote、priScore、candidate、newBlock、confirmBlock
	msgType string

	from string
	to   string

	data []byte

	sendTimestamp time.Time
}

type startVote struct {
}

func (sv *startVote) marshal() []byte {
	returnData, _ := json.Marshal(sv)
	return returnData
}

type priScore struct {
	nodeID string
	score  float64
}

func (ps *priScore) marshal() []byte {
	returnData, _ := json.Marshal(ps)
	return returnData
}

// 特别说明，candidate指的是发送该消息的节点认为可以成为leader的目标节点
type candidate struct {
}

func (cd *candidate) marshal() []byte {
	returnData, _ := json.Marshal(cd)
	return returnData
}

type newBlock struct {
}

type confirmBlock struct {
}

func (cfb *confirmBlock) marshal() []byte {
	returnData, _ := json.Marshal(cfb)
	return returnData
}
