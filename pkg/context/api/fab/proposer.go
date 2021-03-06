/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fab

import (
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
)

// ProposalProcessor simulates transaction proposal, so that a client can submit the result for ordering.
type ProposalProcessor interface {
	ProcessTransactionProposal(ProcessProposalRequest) (*TransactionProposalResponse, error)
}

// ProposalSender provides the ability for a transaction proposal to be created and sent.
type ProposalSender interface {
	CreateTransactionHeader() (TransactionHeader, error)
	SendTransactionProposal(*TransactionProposal, []ProposalProcessor) ([]*TransactionProposalResponse, error)
}

// TransactionID provides the identifier of a Fabric transaction proposal.
type TransactionID string

// EmptyTransactionID represents a non-existing transaction (usually due to error).
const EmptyTransactionID = TransactionID("")

// SystemChannel is the Fabric channel for managaing resources.
const SystemChannel = ""

// TransactionHeader provides a handle to transaction metadata.
type TransactionHeader interface {
	TransactionID() TransactionID
	Creator() []byte
	Nonce() []byte
	ChannelID() string
}

// ChaincodeInvokeRequest contains the parameters for sending a transaction proposal.
type ChaincodeInvokeRequest struct {
	Targets      []ProposalProcessor // Deprecated: this parameter is ignored in the new codes and will be removed shortly.
	ChaincodeID  string
	TransientMap map[string][]byte
	Fcn          string
	Args         [][]byte
}

// TransactionProposal contains a marashalled transaction proposal.
type TransactionProposal struct {
	TxnID TransactionID
	*pb.Proposal
}

// ProcessProposalRequest requests simulation of a proposed transaction from transaction processors.
type ProcessProposalRequest struct {
	SignedProposal *pb.SignedProposal
}

// TransactionProposalResponse respresents the result of transaction proposal processing.
type TransactionProposalResponse struct {
	Endorser string
	Status   int32
	*pb.ProposalResponse
}
