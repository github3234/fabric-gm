/*
Copyright IBM Corp, SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"

	"github.com/tw-bc-group/fabric-gm/core/handlers/auth"
	"github.com/tw-bc-group/fabric-gm/protos/peer"
)

// NewFilter creates a new Filter
func NewFilter() auth.Filter {
	return &filter{}
}

type filter struct {
	next peer.EndorserServer
}

// Init initializes the Filter with the next EndorserServer
func (f *filter) Init(next peer.EndorserServer) {
	f.next = next
}

// ProcessProposal processes a signed proposal
func (f *filter) ProcessProposal(ctx context.Context, signedProp *peer.SignedProposal) (*peer.ProposalResponse, error) {
	return f.next.ProcessProposal(ctx, signedProp)
}

func main() {
}
