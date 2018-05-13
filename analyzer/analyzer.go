// Analyzer package for the Cthulhu Assembler
// Scot W. Stevenson <scot.stevenson@gmail.com>
// First version: 12. May 2018
// This version: 13. May 2018

// The analyzer is where the main processing happens. As the core of the back
// end part of the assembler, it is nicknamed "Azathoth, ruler of the Outer
// Gods"

package analyzer

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"cthulhu/data"
	"cthulhu/node"
	"cthulhu/token"
)

// The analyzer walks the Abstract Syntax Tree (AST) created by the parser and
// modifies it in various ways
func Analyzer(m *data.Machine, trace bool) {

	// FIRST PASS
	walk(m.AST, m.MPU, trace)

	// SECOND PASS
	// TODO

	fmt.Println()
}

// Walk is the main internal routine that visits every node and does something
// depending on type. We break out what we do into little functions to allow
// easier testing and possibly concurrency once we know what we are doing.
func walk(n *node.Node, mpu string, trace bool) {

	var ok bool

	// Paranoid: If node is already completely processed, ignore
	if !n.Done {

		switch n.Type {

		case token.DIREC_PARA:

			switch n.Text {

			case ".mpu":
				// We should have exactly one parameter of the type string
				if len(n.Kids) != 1 {
					log.Fatalf("ANALYZER FATAL (%d,%d): MPU directive takes exactly one parameter, got %d",
						n.Line, n.Index, len(n.Kids))
				}

				k := n.Kids[0]

				if k.Type != token.STRING {
					log.Fatalf("ANALYZER FATAL (%d,%d): MPU directive takes a STRING, got %s",
						k.Line, k.Index, token.Name[k.Type])
				}

				if k.Text != "65816" && k.Text != "65c02" && k.Text != "6502" {
					log.Fatalf("ANALYZER FATAL (%d,%d): MPU type '%s' not supported",
						k.Line, k.Index, k.Text)
				}

				// The main program has already made sure that the MPU the user
				// requested is legal
				// TODO check if MPU types match

			}

		// NUMBER CONVERSION: Convert the strings kept in node.Text and store
		// them as values in node.Value; change node.Type to token.DEC_NUM.
		// Remember that binary and hex numbers as strings can contain ":" and
		// "." to make reading easier

		// Paranoid: If for some reason we were given the generic number type (which
		// shouldn't happen) we first convert it to the special type
		case token.NUMBER:
			switch {

			case strings.HasPrefix(n.Text, "$"):
				n.Type = token.HEX_NUM
			case strings.HasPrefix(n.Text, "%"):
				n.Type = token.BIN_NUM
			default:
				n.Type = token.DEC_NUM
			}

			// Now that we know what type we are dealing with, continue to do the
			// actual conversion
			fallthrough

		case token.BIN_NUM:
			n.Value, ok = convertNum(n.Text, 2)
			if !ok {
				log.Fatalf("ANALYZER FATAL: (%d,%d): Can't convert binary number string '%s' to int",
					n.Line, n.Index, n.Text)
			}

			if trace {
				fmt.Printf("ANALYZER (%d, %d): Processed BIN_NUM %s, now %d\n",
					n.Line, n.Index, n.Text, n.Value)
			}

			n.Type = token.DEC_NUM

		// Decimal numbers don't contain ":" or "." so we don't break it out
		// into a separate function
		case token.DEC_NUM:
			v, err := strconv.Atoi(n.Text)
			if err != nil {
				log.Fatalf("ANALYZER FATAL: (%d,%d): Can't convert decimal number string '%s' to int",
					n.Line, n.Index, n.Text)
			}
			n.Value = int(v)

			if trace {
				fmt.Printf("ANALYZER (%d, %d): Processed DEC_NUM %s, now %d\n",
					n.Line, n.Index, n.Text, n.Value)
			}

		case token.HEX_NUM:
			n.Value, ok = convertNum(n.Text, 16)
			if !ok {
				log.Fatalf("ANALYZER FATAL: (%d,%d): Can't convert hex number string '%s' to int",
					n.Line, n.Index, n.Text)
			}

			if trace {
				fmt.Printf("ANALYZER (%d, %d): Processed HEX_NUM %s, now %d\n",
					n.Line, n.Index, n.Text, n.Value)
			}

			n.Type = token.DEC_NUM

		// STRING CONVERSION: Convert string from node.Text to a sequence of
		// bytes. Store them in node.Code. Mark node as done. Note that Go
		// converts strings as unicode, which not all assembler programs are
		// equipped to handle
		case token.STRING:
			n.Code = []byte(n.Text)
			n.Done = true

			if trace {
				fmt.Printf("ANALYZER (%d, %d): Processed STRING \"%s\", now %s\n",
					n.Line, n.Index, n.Text, node.FormatByteSlice(n.Code))
			}

		// OPCODE 0 CONVERSION: Opcodes with no operands
		case token.OPC_0:
			oc, ok := getOpcode(mpu, n.Text)
			if !ok {
				log.Fatalf("ANALYZER FATAL (%d, %d): Opcode '%s' unrecognized",
					n.Line, n.Index, n.Text)
			}

			n.Code = append(n.Code, oc)
			n.Done = true
		}

		// If this node doesn't have kids, we're done. This ends the
		// recursion
		if len(n.Kids) == 0 {
			return
		}

		// We have kids, but we kick out those deadbeats such as
		// comments and empty lines that just suck all our energy
		var newKids []*node.Node

		for i := 0; i < len(n.Kids); i++ {

			tt := n.Kids[i].Type

			if tt != token.EMPTY &&
				tt != token.COMMENT &&
				tt != token.EOL &&
				tt != token.COMMENT_LINE {
				newKids = append(newKids, n.Kids[i])
			}
		}

		n.Kids = newKids

		// We've got good kids now, let's walk them recursively
		for _, k := range n.Kids {
			walk(k, mpu, trace)
		}
	}
}

// INDIVIDUAL STEPS

// convertNum takes a number string that includes ":" and "." and a base int, and
// returns an int value as well as a code for success or failure. We
// artificially limit the bases to 2 and 16.
func convertNum(s string, base int) (int, bool) {

	if base != 2 && base != 16 {
		log.Fatalf("ANALYZER FATAL: Given base %d to convert, must be binary, hex or decimal",
			base)
	}

	ok := true

	s0 := strings.Replace(s, ":", "", -1)
	s1 := strings.Replace(s0, ".", "", -1)
	s2 := strings.TrimSpace(s1) // paranoid, should have been done by lexer

	v, err := strconv.ParseInt(s2, base, 64)
	if err != nil {
		ok = false
	}

	return int(v), ok
}

// getOpcode takes a Simpler Assembler Notation (SAN) mnemonic and returns the
// opcode value and a flag to single if it went okay
func getOpcode(mpu string, mn string) (byte, bool) {
	oc, ok := data.OpcodesSAN[mpu][mn]
	return oc.Value, ok
}
