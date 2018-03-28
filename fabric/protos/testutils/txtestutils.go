package testutils

import (
	pb "github.com/hyperledger2/fabric/protos/peer"
	"github.com/hyperledger2/fabric/msp"
	"github.com/hyperledger2/fabric/protos/common"
	putils "github.com/hyperledger2/fabric/protos/utils"
)

var (
	signer msp.SigningIdentity
)

// ConstructSignedTxEnv constructs a transaction envelop for tests
func ConstructSingedTxEnv(chainID string, ccid *pb.ChaincodeID, pResponse *pb.Response, simulationResults []byte, txid string, events []byte, visibility []byte, signer msp.SigningIdentity) (*common.Envelope, string, error) {
	ss, err := signer.Serialize()
	if err != nil {
		return nil, "", err
	}

	var prop *pb.Proposal
	if txid == "" {
		prop, txid, err = putils.CreateChaincodeProposal()
	}
}

// ConstructSingedTxEnvWithDefaultSigner constructs a transaction envelop for tests with a default signer.
// This method helps other modules to construct a transaction with supplied parameters
func ConstructSingedTxEnvWithDefaultSigner(chainID string, ccid *pb.ChaincodeID, response *pb.Response, simulationResults []byte, txid string, events []byte, visibility []byte) (*common.Envelope, string, error) {
	return ConstructSingedTxEnv(chainID, ccid, response, simulationResults, txid, events, visibility, signer)
}