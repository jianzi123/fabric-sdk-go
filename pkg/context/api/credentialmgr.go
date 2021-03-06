/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import "github.com/hyperledger/fabric-sdk-go/pkg/context/api/core"

// SigningIdentity is the identity object that encapsulates the user's private key for signing
// and the user's enrollment certificate (identity)
type SigningIdentity struct {
	MspID          string
	EnrollmentCert []byte
	PrivateKey     core.Key
}

// CredentialManager retrieves user's signing identity
type CredentialManager interface {
	GetSigningIdentity(name string) (*SigningIdentity, error)
}
