// Formatter Package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 04. May 2018
// This version: 06. May 2018

/* The formatter produces a standardized formatted source
   listing, following the example of gofmt
*/

package formatter

import (
	"fmt"

	"goasm65816/token"
)

const (
	indent1 = "        "
	indent2 = indent1 + indent1
)

func Formatter(tl *[]token.Token) bool {

	ok := true

	for _, t := range *(tl) {

		switch t.Type {
		case token.T_eol:
			fmt.Print("\n")
			continue

		case token.T_localLabel:
			fmt.Print(t.Text, "\n")
			continue

		case token.T_label:
			fmt.Print(t.Text, "\n")
			continue

		case token.T_comment:
			fmt.Print(t.Text)
			continue

		case token.T_directive:
			fmt.Print(indent1)
			fmt.Print(t.Text, " ")
			continue

		case token.T_opcode0:
			fmt.Print(indent2)
			fmt.Print(t.Text, " ")
			continue
		}

		t.Print()
	}

	fmt.Print("\n")

	return ok
}
