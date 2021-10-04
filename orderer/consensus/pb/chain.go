package pb

import (
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/orderer/consensus"
)

type Chain struct {
	node *node

	support  consensus.ConsenterSupport
	metadata *common.Metadata
	exitChan chan struct{}
}

func NewChain(support consensus.ConsenterSupport, metadata *common.Metadata) (*Chain, error) {

	newChain := Chain{
		support:  support,
		metadata: metadata,
	}
	return &newChain, nil
}

func (pbc *Chain) Start() {
	go pbc.main()
}

func (pbc *Chain) Halt() {

}

func (pbc *Chain) Order(env *common.Envelope, configSeq uint64) error {

	return nil
}

func (pbc *Chain) Configure(env *common.Envelope, configSeq uint64) error {

	return nil
}

func (pbc *Chain) WaitReady() error {

	return nil
}

func (pbc *Chain) Errored() <-chan struct{} {

	return pbc.exitChan
}

func (pbc *Chain) main() {

}
