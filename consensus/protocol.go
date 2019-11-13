// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package consensus implements different Ethereum consensus engines.
package consensus

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// Constants to match up protocol versions and messages
const (
	Eth62 = 62
	Eth63 = 63
)

var (
	EthProtocol = Protocol{
		Name:     "eth",
		Versions: []uint{Eth62, Eth63},
		Lengths:  []uint64{17, 8},
	}
)

// Protocol defines the protocol of the consensus
type Protocol struct {
	// Official short name of the protocol used during capability negotiation.
	Name string
	// Supported versions of the eth protocol (first is primary).
	Versions []uint
	// Number of implemented message corresponding to different protocol versions.
	Lengths []uint64
	// Whether this should be the primary form of communication between nodes that support this protocol.
	Primary bool
}

// Broadcaster defines the interface to enqueue blocks to fetcher, find peer
type Broadcaster interface {
	// Enqueue add a block into fetcher queue
	Enqueue(id string, block *types.Block)
	// FindPeers retrives peers by addresses
	FindPeers(targets map[enode.ID]bool, label string) map[enode.ID]Peer
	// GetNodeKey retrieves the node's private key
	GetNodeKey() *ecdsa.PrivateKey
}

// Server defines the interface for a p2p.server to get the local node's enode and to addlabel/removelabel for static/trusted peers
type P2PServer interface {
	// Gets this node's enode
	Self() *enode.Node
	// AddPeerLabel will add a peer label to the p2p server instance
	AddPeerLabel(node *enode.Node, label string)
	// RemovePeerLabel will remove a peer label from the p2p server instance
	RemovePeerLabel(node *enode.Node, label string)
	// AddTrustedPeerLabel will add a trusted peer label to the p2p server instance
	AddTrustedPeerLabel(node *enode.Node, label string)
	// RemoveTrustedPeerLabel will remove a trusted peer label from the p2p server instance
	RemoveTrustedPeerLabel(node *enode.Node, label string)
}

// Peer defines the interface for a p2p.peer
type Peer interface {
	// Send sends the message to this peer
	Send(msgcode uint64, data interface{}) error
	// Returns the peer's enode
	Node() *enode.Node
}
