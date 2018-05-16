// Basic data structures for the Cthulhu Assembler
// A lot of these are used as sets since Go (golang) doesn't
// provide that data structure by default
// First version: 04. May 2018 (May the Force be with you!)
// This version: 11. May 2018

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

// List of all directives (actually used as a set)
var Directives = map[string](bool){
	".mpu": true, ".origin": true, ".equ": true, ".byte": true,
	".word": true, ".native": true, ".emulated": true, ".end": true,
	".a8": true, ".a16": true, ".xy8": true, ".xy16": true,
	".axy8": true, ".axy16": true, ".scope": true, ".scend": true,
	".macro": true, ".macend": true, ".lsb": true, ".msb": true,
	".bank": true, ".advance": true, ".skip": true,
	".assert": true, ".ram": true, ".rom": true,
	".swap": true, ".drop": true, ".dup": true, ".lshift": true,
	".rshift": true, ".not": true, ".here": true, ".include": true,
	"...": true,
}

// List of directives with Parameters
var DirectivesPara = map[string](bool){
	".mpu": true, ".origin": true, ".equ": true, ".byte": true,
	".word": true, ".macro": true, ".lsb": true, ".msb": true,
	".bank": true, ".advance": true, ".skip": true,
	".assert": true, ".ram": true, ".rom": true, ".include": true,
	".lshift": true, ".rshift": true, ".not": true,
}

// List of directives that are actually operators, that is, they don't start a
// line but are part of a math term
var Operators = map[string](bool){
	".lshift": true, ".rshift": true, ".lsb": true, ".msb": true,
	".bank": true, ".and": true, ".or": true, ".xor": true,
	".not": true, ".dup": true, ".swap": true, ".drop": true,
}
