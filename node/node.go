// Node types for the AST of the Cthulhu Assembler
// Scot W. Stevenson
// First version 07. May 2018
// This version 11. May 2018

// Because we have a simple assembler and are not going to use obscene amounts
// of data, we can get away with a homogenous Abstract Syntax Tree (AST) with
// normalized children instead of having various specialized nodes. This also
// means we don't have to define interfaces, but can just used methods to the
// node type

package node

import (
	"cthulhu/token"
)

// Homogeneous node stucture. Not all of these are used for every type of token,
// the price we pay for a single homogenous node type.
type Node struct {
	token.Token         // embedding adds Type, Text, Line, Index
	Kids        []*Node // for children nodes
	Value       int     // for numbers of all sorts
	Code        []byte  // The final byte stream that is added at the end
}

// Add creates a new subnode on an existing node. This is just a nicer way of
// saying append()
func (n *Node) Add(k *Node) {
	n.Kids = append(n.Kids, k)
}

// Adopt creates a new node and adds to the node that is passed along. This is
// just a nice way of combining Create and Add
func (n *Node) Adopt(k *Node, t *token.Token) {
	nn := Create(*t)
	n.Kids = append(n.Kids, &nn)
}

// Creates returns a new node when given a token.
func Create(t token.Token) Node {
	return Node{Token: t}
}
