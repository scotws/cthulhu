// Lister Package for the GoAsm65816 assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 02. May 2018
// This version: 02. May 2018

/* The lister produces two text files for the user: A cleanly formatted
   listing of source code, and a list of the labels used.
*/

package lister

import (
	"fmt"
)

func Lister() {
	fmt.Println("(At lister)")
}
