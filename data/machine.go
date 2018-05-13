// High-level structure of the code once we're past the parser

package data

import "cthulhu/node"

type Machine struct {
	MPU    string     // MPU as given by the user on the command line
	Origin int        // Start address for compilation as given in the source code
	Code   []byte     // Finished compiled data
	AST    *node.Node // The Abstract Syntax Tree (AST)
}
