/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"github.com/hyperledger2/fabric/common/crypto"
	"github.com/hyperledger2/fabric/protos/common"
	"github.com/hyperledger2/fabric/protos/peer"
	"github.com/hyperledger2/fabric/bccsp/factory"
	"github.com/hyperledger2/fabric/bccsp"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
)


// CreateChaincodeProposal creates a proposal from given input.
// It returns the proposal and the transaction id associated to the proposal
func CreateChaincodeProposal(typ common.HeaderType, chainID string, cis *peer.ChaincodeInvocationSpec, creator []byte) (*peer.Proposal, string, error) {
	return CreateChaincodeProposalWithTransient(typ, chainID, cis, creator, nil)
}

// CreateChaincodeProposalWithTransient creates a proposal from given input
// It returns the proposal and the transaction id associated to the proposal
func CreateChaincodeProposalWithTransient(typ common.HeaderType, chainID string, cis *peer.ChaincodeInvocationSpec, creator []byte, transientMap map[string][]byte) (*peer.Proposal, string, error) {
	// generate a random nonce
	nonce, err := crypto.GetRandomNonce()
	if err != nil {
		return nil, "", err
	}

	// compute txid
	txid, err := ComputeProposalTxID(nonce, creator)
	if err != nil {
		return nil, "", err
	}

	return CreateChaincodeProposalWithTxIDNonceAndTransient(txid, typ, chainID, cis, nonce, creator, transientMap)
}

func CreateChaincodeProposalWithTxIDNonceAndTransient(txid string, typ common.HeaderType, chainID string, cis *peer.ChaincodeInvocationSpec, nonce, creator []byte, transientMap map[string][]byte) (*peer.Proposal, string, error) {
	ccHdrExt := &peer.ChaincodeHeaderExtension{ChaincodeId: cis.ChaincodeSpec.ChaincodeId}
	ccHdrExtBytes, err := proto.Marshal(ccHdrExt)
	if err != nil {
		return nil, "", err
	}

	cisBytes, err := proto.Marshal(cis)
	if err != nil {
		return nil, "", err
	}

	ccPropPayload := peer.ChaincodeProposalPayload{Input: cisBytes, TransientMap: transientMap}
	ccPropPayloadBytes, err := proto.Marshal(ccPropPayload)

}


// ComputeProposalTxID computes TxID as the Hash computed
// over the concatenation of nonce and creator.
func ComputeProposalTxID(nonce, creator []byte) (string, error) {
	// TODO: Get the Hash function to be used from channel configuration
	digest, err := factory.GetDefault().Hash(
		append(nonce, creator...),
		&bccsp.SHA256Opts{})
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(digest), nil
}