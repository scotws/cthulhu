// Formatter Package for the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 04. May 2018
// This version: 11. May 2018

// The formatter produces a cleanly formatted source file, following
// the example of gofmt

// TODO This version is based on the token list from the lexer. We need a
// version baed on the node tree (AST) from the parser

// TODO remove double empty lines
// TODO print .byte, .word, .long nicely

package formatter

import (
	"fmt"

	"cthulhu/token"
)

const (
	indent1 = "        "
	indent2 = indent1 + indent1
)

func Formatter(tl *[]token.Token) {

	for _, t := range *(tl) {

		switch t.Type {

		case token.EOL, token.EMPTY, token.EOF:
			fmt.Print("\n")
			continue

		case token.LOCAL_LABEL, token.LABEL:
			fmt.Print(t.Text, ":")
			continue

		case token.ANON_LABEL:
			fmt.Print(t.Text)
			continue

		case token.STRING:
			fmt.Printf("\"%s\"", t.Text)
			continue

		case token.BIN_NUM:
			fmt.Print("%", t.Text)
			continue

		case token.HEX_NUM:
			fmt.Print("$", t.Text)
			continue

		case token.COMMENT, token.COMMENT_LINE:
			fmt.Print(t.Text)
			continue

		case token.DIREC, token.DIREC_PARA:
			fmt.Print(indent1)
			fmt.Print(t.Text, " ")
			continue

		case token.OPC_0, token.OPC_1, token.OPC_2:
			fmt.Print(indent2)
			fmt.Print(t.Text, " ")
			continue

		default:
			fmt.Print(t.Text, " ")
		}

	}

	fmt.Print("\n")
}
