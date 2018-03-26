package testutils

import (
	pb "github.com/hyperledger2/fabric/protos/peer"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger2/fabric/protos/common"
)
// ConstructSignedTxEnv constructs a transaction envelop for tests
func ConstructSignedTxEnv(chainID string, ccid *pb.ChaincodeID, pResponse *pb.Response, simulationResults []byte, txid string, events []byte, visibility []byte, signer msp.SigningIdentity) (*common.Envelope, string error) {

}