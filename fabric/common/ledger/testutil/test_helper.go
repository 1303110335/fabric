package testutil

import (
	"testing"
	"github.com/hyperledger2/fabric/protos/common"
	pb "github.com/hyperledger2/fabric/protos/peer"
)

// ConstructBlock constructs a single block
func ConstructBlock(t *testing.T, blockNum uint64, previousHash []byte, simulationResults [][]byte, sign bool) *common.Block {
	envs := []*common.Envelope{}
	for i := 0; i < len(simulationResults); i++ {
		env, _, err := ConstructTransaction(t, simulationResults[i], "", sign)
		if err != nil {
			t.Fatalf("ConstructTestTransaction failed, err %s", err)
		}
		envs = append(envs, env)
	}
	return NewBlock(envs, blockNum, previousHash)
}

//construct a single block with random content
func ConstructTestBlock(t *testing.T, blockNum uint64, numTx int, txSize int) *common.Block {
	simulationResults := [][]byte{}
	for i := 0; i < numTx; i++ {
		simulationResults = append(simulationResults, ContructRandomBytes(t, txSize))
	}
	return ConstructBlock(t, blockNum, ContructRandomBytes(t, 32), simulationResults, false)
}

//ConstructTransaction constructs a transaction for testing
func ConstructTransaction(_ *testing.T, simulationResult []byte, txid string, sign bool) (*common.Envelope, string, error) {
	ccid := &pb.ChaincodeID{
		Name:    "foo",
		Version: "v1",
	}
	var txID string
	var txEnv *common.Envelope
	var err error
	if sign {

	}
}
