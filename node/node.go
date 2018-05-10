// Node data for goasm65816
// Scot W. Stevenson
// First version 07. May 2018
// This version 09. May 2018

package node

import (
	"goasm65816/token"
)

// Homogeneous node stucture. We might have to add stuff to this later
// for the further passes
type Node struct {
	token.Token         // adds Type, Text, Line, Index
	Kids        []*Node // for children nodes
}

// addKid creates a new subnode on an existing node. This is just a nicer way of
// saying append()
func (n *Node) Add(k *Node) {
	n.Kids = append(n.Kids, k)
}

// Creates returns a new node when given a token.
func Create(t token.Token) Node {
	return Node{t, nil}
}
