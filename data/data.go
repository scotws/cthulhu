// Contains basic data structures used by other routines
// A lot of these are used as sets since Go (golang) doesn't
// provide that data structure by default
// First version: 04. May 2018 (May the Force be with you!)
// This version: 06. May 2018

package data

type Opcode struct {
	SAN       string // lowercase SAN mnemonic
	WDC       string // lowercase WDC mnemonic
	Length    int    // number of bytes including opcode itself
	Operands  int    // number of operands
	Value     byte   // opcode
	Embiggens bool   // if the opcode changes length with 8/16 switch
}

var Directives = map[string](bool){
	".mpu": true, ".origin": true, ".equ": true, ".byte": true,
	".word": true, ".native": true, ".emulated": true, ".end": true,
	".a8": true, ".a16": true, ".xy8": true, ".xy16": true,
	".axy8": true, ".axy16": true, ".scope": true, ".scend": true,
	".macro": true, ".macend": true, ".lsb": true, ".msb": true,
	".bank": true, ".advance": true, ".skip": true,
	".assert": true, ".ram": true, ".rom": true, ".notation": true,
}
