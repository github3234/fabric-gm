/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package server

import (
	"github.com/tw-bc-group/fabric-gm/core/peer"
	"github.com/tw-bc-group/fabric-gm/token/ledger"
	"github.com/pkg/errors"
)

// PeerLedgerManager implements the LedgerManager interface
// by using the peer infrastructure
type PeerLedgerManager struct {
}

func (*PeerLedgerManager) GetLedgerReader(channel string) (ledger.LedgerReader, error) {
	l := peer.Default.GetLedger(channel)
	if l == nil {
		return nil, errors.Errorf("ledger not found for channel %s", channel)
	}

	return l.NewQueryExecutor()
}
