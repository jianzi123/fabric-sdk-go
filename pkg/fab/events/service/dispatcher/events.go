/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package dispatcher

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/context/api/fab"
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
)

// Event is an event that's sent to the dispatcher. This includes client registration
// requests or events that come from an event producer.
type Event interface{}

// RegisterEvent is the base for all registration events.
type RegisterEvent struct {
	RegCh chan<- fab.Registration
	ErrCh chan<- error
}

// StopEvent tells the dispatcher to stop processing
type StopEvent struct {
	RegCh chan<- error
}

// RegisterBlockEvent registers for block events
type RegisterBlockEvent struct {
	RegisterEvent
	Reg *BlockReg
}

// RegisterFilteredBlockEvent registers for filtered block events
type RegisterFilteredBlockEvent struct {
	RegisterEvent
	Reg *FilteredBlockReg
}

// RegisterChaincodeEvent registers for chaincode events
type RegisterChaincodeEvent struct {
	RegisterEvent
	Reg *ChaincodeReg
}

// RegisterTxStatusEvent registers for transaction status events
type RegisterTxStatusEvent struct {
	RegisterEvent
	Reg *TxStatusReg
}

// UnregisterEvent unregisters a registration
type UnregisterEvent struct {
	Reg fab.Registration
}

// NewRegisterBlockEvent creates a new RegisterBlockEvent
func NewRegisterBlockEvent(filter fab.BlockFilter, eventch chan<- *fab.BlockEvent, respch chan<- fab.Registration, errCh chan<- error) *RegisterBlockEvent {
	return &RegisterBlockEvent{
		Reg:           &BlockReg{Filter: filter, Eventch: eventch},
		RegisterEvent: NewRegisterEvent(respch, errCh),
	}
}

// NewRegisterFilteredBlockEvent creates a new RegisterFilterBlockEvent
func NewRegisterFilteredBlockEvent(eventch chan<- *fab.FilteredBlockEvent, respch chan<- fab.Registration, errCh chan<- error) *RegisterFilteredBlockEvent {
	return &RegisterFilteredBlockEvent{
		Reg:           &FilteredBlockReg{Eventch: eventch},
		RegisterEvent: NewRegisterEvent(respch, errCh),
	}
}

// NewUnregisterEvent creates a new UnregisterEvent
func NewUnregisterEvent(reg fab.Registration) *UnregisterEvent {
	return &UnregisterEvent{
		Reg: reg,
	}
}

// NewRegisterChaincodeEvent creates a new RegisterChaincodeEvent
func NewRegisterChaincodeEvent(ccID, eventFilter string, eventch chan<- *fab.CCEvent, respch chan<- fab.Registration, errCh chan<- error) *RegisterChaincodeEvent {
	return &RegisterChaincodeEvent{
		Reg: &ChaincodeReg{
			ChaincodeID: ccID,
			EventFilter: eventFilter,
			Eventch:     eventch,
		},
		RegisterEvent: NewRegisterEvent(respch, errCh),
	}
}

// NewRegisterTxStatusEvent creates a new RegisterTxStatusEvent
func NewRegisterTxStatusEvent(txID string, eventch chan<- *fab.TxStatusEvent, respch chan<- fab.Registration, errCh chan<- error) *RegisterTxStatusEvent {
	return &RegisterTxStatusEvent{
		Reg:           &TxStatusReg{TxID: txID, Eventch: eventch},
		RegisterEvent: NewRegisterEvent(respch, errCh),
	}
}

// NewRegisterEvent creates a new RgisterEvent
func NewRegisterEvent(respch chan<- fab.Registration, errCh chan<- error) RegisterEvent {
	return RegisterEvent{
		RegCh: respch,
		ErrCh: errCh,
	}
}

// NewChaincodeEvent creates a new ChaincodeEvent
func NewChaincodeEvent(chaincodeID, eventName, txID string) *fab.CCEvent {
	return &fab.CCEvent{
		ChaincodeID: chaincodeID,
		EventName:   eventName,
		TxID:        txID,
	}
}

// NewTxStatusEvent creates a new TxStatusEvent
func NewTxStatusEvent(txID string, txValidationCode pb.TxValidationCode) *fab.TxStatusEvent {
	return &fab.TxStatusEvent{
		TxID:             txID,
		TxValidationCode: txValidationCode,
	}
}

// NewStopEvent creates a new StopEvent
func NewStopEvent(respch chan<- error) *StopEvent {
	return &StopEvent{
		RegCh: respch,
	}
}
