// Node data for goasm65816
// Scot W. Stevenson
// First version 07. May 2018
// This version 07. May 2018

package node

import (
	"goasm65816/token"
)

type Node struct {
	token.Token         // adds Type, Text, Line, Index
	Kids        []*Node // for children nodes
	Bytes       []byte  // used later for code
}

// addKid creates a new subnode on an existing node
func (n *Node) Add(k *Node) {
	n.Kids = append(n.Kids, k)
}

// Generic walking function, depth first
func Walk(AST *Node, f func(*Node)) {

	f(AST)

	if len(AST.Kids) == 0 {
		return
	}

	for _, k := range AST.Kids {
		Walk(k, f)
	}
}
