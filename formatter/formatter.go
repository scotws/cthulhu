// Formatter Package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 04. May 2018
// This version: 10. May 2018

// The formatter produces a cleanly formatted source file, following
// the example of gofmt

package formatter

import (
	"fmt"

	"goasm65816/token"
)

const (
	indent1 = "        "
	indent2 = indent1 + indent1
)

func Formatter(tl *[]token.Token) {

	for _, t := range *(tl) {

		switch t.Type {
		case token.EOL:
			fmt.Print("\n")
			continue

		case token.LOCAL_LABEL:
			fmt.Print(t.Text, "\n")
			continue

		case token.LABEL:
			fmt.Print(t.Text, "\n")
			continue

		case token.COMMENT:
			fmt.Print(t.Text)
			continue

		case token.DIREC:
			fmt.Print(indent1)
			fmt.Print(t.Text, " ")
			continue

		case token.SAN_0:
			fmt.Print(indent2)
			fmt.Print(t.Text, " ")
			continue
		}

		t.Print()
	}

	fmt.Print("\n")
}
