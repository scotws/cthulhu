// Analyzer package for the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 12. May 2018
// This version: 12. May 2018

// The analyzer is where the main processing happens. As the core of the back
// end part of the assembler, it is nicknamed "Azathoth, ruler of the Outer
// Gods"

package analyzer

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"cthulhu/node"
	"cthulhu/token"
)

// The analyzer walks the Abstract Syntax Tree (AST) created by the parser and
// modifies it in various ways
func Analyzer(ast *node.Node, debug bool) *node.Node {

	// FIRST PASS
	walk(ast, debug)

	// SECOND PASS

	fmt.Println()
	return ast
}

// cleanupNumbers takes a number string and removes the visual helper runes ':'
// and '.', returning a number string
func cleanupNumber(s string) string {
	s0 := strings.Replace(s, ":", "", -1)
	s1 := strings.Replace(s0, ".", "", -1)
	s2 := strings.TrimSpace(s1) // paranoid, should have been done by lexer
	return s2
}

// Walk is the main internal routine that visits every node and does something
// depending on type
func walk(n *node.Node, trace bool) {

	switch n.Type {

	case token.BIN_NUM:
		s0 := cleanupNumber(n.Text)
		v, err := strconv.ParseInt(s0, 2, 64)

		if err != nil {
			log.Fatalf("ANALYZER FATAL: (%d,%d): Can't convert string '%s' to number",
				n.Line, n.Index, s0)
		}
		n.Value = int(v)
		n.Type = token.DEC_NUM
		n.Modified = true

		if trace {
			fmt.Printf("ANALYZER (%d, %d): Processed decimal number '%s'\n",
				n.Line, n.Index, n.Text)
		}

	case token.DEC_NUM:
		v, err := strconv.Atoi(n.Text)
		if err != nil {
			log.Fatalf("ANALYZER FATAL: (%d,%d): Can't convert string '%s' to number",
				n.Line, n.Index, n.Text)
		}
		n.Value = int(v)
		n.Modified = true

		if trace {
			fmt.Printf("ANALYZER (%d, %d): Processed decimal number '%s'\n",
				n.Line, n.Index, n.Text)
		}

	case token.HEX_NUM:
		s0 := cleanupNumber(n.Text)
		v, err := strconv.ParseInt(s0, 16, 64)

		if err != nil {
			log.Fatalf("ANALYZER ERROR: (%d,%d): Can't convert string '%s' to number",
				n.Line, n.Index, s0)
		}
		n.Value = int(v)
		n.Type = token.DEC_NUM
		n.Modified = true

		if trace {
			fmt.Printf("ANALYZER (%d, %d): Processed decimal number '%s'\n",
				n.Line, n.Index, n.Text)
		}

	default: // Do nothing
	}

	// if this node doesn't have kids, we're done
	if len(n.Kids) == 0 {
		return
	}

	// We have kids, but we don't want those freeloaders like comments and
	// empty lines that just suck away our energy
	for i := 0; i < len(n.Kids); i++ {

		tt := n.Kids[i].Type

		switch tt {

		case token.EMPTY, token.COMMENT:
			n.Kids = n.Evict(i)

		// If we have a full-line comment, we have to remove the EOL as
		// well
		case token.COMMENT_LINE:
			n.Kids = n.Evict(i)
			n.Kids = n.Evict(i)
			i--
		}

		if trace {
			fmt.Printf("ANALYZER (%d, %d): Removed subnode of type '%s'\n",
				n.Line, n.Index, token.Name[tt])
		}
		n.Modified = true
	}

	// we've got kids, let's walk them recursively
	for _, k := range n.Kids {
		walk(k, trace)
	}

}
