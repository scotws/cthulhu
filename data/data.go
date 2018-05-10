// Basic data structures for the Cthulhu Assembler
// A lot of these are used as sets since Go (golang) doesn't
// provide that data structure by default
// First version: 04. May 2018 (May the Force be with you!)
// This version: 09. May 2018

package data

type Opcode struct {
	SAN       string // lowercase SAN mnemonic
	WDC       string // lowercase WDC mnemonic
	Length    int    // number of bytes including opcode itself
	Operands  int    // number of operands
	Value     byte   // opcode
	Embiggens bool   // if the opcode changes length with 8/16 switch
}

var OpcodesSAN = map[string]map[string](Opcode){
	"6502":  Opcodes6502,
	"65c02": Opcodes65c02,
	"65816": Opcodes65816,
}

var OpcodesWDC = map[string]map[string](bool){
	"6502":  MneWDC6502,
	"65c02": MneWDC65c02,
	"65816": MneWDC65816,
}

// List of all directives (actually used as a set)
var Directives = map[string](bool){
	".mpu": true, ".origin": true, ".equ": true, ".byte": true,
	".word": true, ".native": true, ".emulated": true, ".end": true,
	".a8": true, ".a16": true, ".xy8": true, ".xy16": true,
	".axy8": true, ".axy16": true, ".scope": true, ".scend": true,
	".macro": true, ".macend": true, ".lsb": true, ".msb": true,
	".bank": true, ".advance": true, ".skip": true,
	".assert": true, ".ram": true, ".rom": true, ".notation": true,
}

// List of directives with Parameters
var DirectivesPara = map[string](bool){
	".mpu": true, ".origin": true, ".equ": true, ".byte": true,
	".word": true, ".macro": true, ".lsb": true, ".msb": true,
	".bank": true, ".advance": true, ".skip": true,
	".assert": true, ".ram": true, ".rom": true, ".notation": true,
}
