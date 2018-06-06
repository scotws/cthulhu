// Symbol Table Code for the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version 21. May 2018
// This version 21. May 2018

package analyzer

type Symbol struct {
	Value int    // added once defined
	File  string // where defined
	Line  int    // where defined
	Type  string // "label" or "symbol"
	Used  bool   // see if symbol defined but not used
}

var (
	SymbolTable = map[string]Symbol{}
)
