// Token Package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 02. May 2018

package parser

import (
	"fmt"

	"goasm65816/token"
)

func Parser(tl *[]token.Token) bool {

	ok := true

	// TODO Test routines

	for _, t := range *(tl) {
		t.PrintLine()
	}

	fmt.Println("------------")

	for _, t := range *(tl) {
		t.Print()
	}

	return ok
}
