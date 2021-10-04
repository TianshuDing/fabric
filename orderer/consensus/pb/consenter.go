package pb

import (
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/orderer/consensus"
)

var logger = flogging.MustGetLogger("orderer.consensus.pb")

type Consenter struct {
	// maybe have something here
}

func New() *Consenter {

	return &Consenter{}
}

func (pbc *Consenter) HandleChain(support consensus.ConsenterSupport, metadata *cb.Metadata) (consensus.Chain, error) {
	logger.Info("Use of the pb orderer is not perfect and remains only for use in test environments.")
	return NewChain(support, metadata)
}
