package pb

import "github.com/hyperledger/fabric/common/flogging"

type node struct {
	chainID string
	nodeID  string

	logger *flogging.FabricLogger
}
