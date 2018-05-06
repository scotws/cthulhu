// Opcode tables for the 65816
// First version: 06. May 2018
// This version: 06. May 2018

package data

// Map (actually set) of legal WDC mnemonics for the 65816
// TODO missing jml, pei, rtl, etc
var MnemonicsWDC65816 = map[string](bool){
	"adc": true, "and": true, "asl": true, "bcc": true, "bcs": true, "beq": true,
	"bit": true, "bmi": true, "bne": true, "bpl": true, "bra": true, "brk": true,
	"bvc": true, "bvs": true, "clc": true, "cld": true, "cli": true, "clv": true,
	"cmp": true, "cop": true, "cpx": true, "cpy": true, "dec": true, "dex": true,
	"dey": true, "eor": true, "inc": true, "inx": true, "iny": true, "jmp": true,
	"jsr": true, "lda": true, "ldx": true, "ldy": true, "lsr": true, "mvn": true,
	"mvp": true, "nop": true, "ora": true, "pha": true, "phb": true, "phd": true,
	"phe": true, "phk": true, "php": true, "phx": true, "phy": true, "pla": true,
	"plb": true, "pld": true, "plp": true, "plx": true, "ply": true, "rep": true,
	"rol": true, "ror": true, "rti": true, "rts": true, "sbc": true, "sec": true,
	"sed": true, "sei": true, "sep": true, "sta": true, "stp": true, "stx": true,
	"sty": true, "stz": true, "tax": true, "tay": true, "tcd": true, "tcs": true,
	"tdc": true, "trb": true, "tsb": true, "tsc": true, "tsx": true, "txa": true,
	"txs": true, "txy": true, "tya": true, "tyx": true, "wai": true, "wdm": true,
	"xba": true, "xce": true,
}

// Data bank of opcodes, with mnemonics (SAN and WDC), length in bytes, number
// of operands, the actual opcode, and a flag if the opcode is affected by
// switches of the register size
var Opcodes65816 = map[string](Opcode){
	"brk":      Opcode{"brk", "brk", 2, 1, 0x00, false}, // we require a signature byte
	"ora.dxi":  Opcode{"ora.dxi", "ora", 2, 1, 0x01, false},
	"cop":      Opcode{"cop", "cop", 2, 0, 0x02, false},
	"ora.s":    Opcode{"ora.s", "ora", 2, 1, 0x03, false},
	"tsb.d":    Opcode{"tsb.d", "tsb", 2, 1, 0x04, false},
	"ora.d":    Opcode{"ora.d", "ora", 2, 1, 0x05, false},
	"asl.d":    Opcode{"asl.d", "asl", 2, 1, 0x06, false},
	"ora.dil":  Opcode{"ora.dil", "ora", 2, 1, 0x07, false},
	"php":      Opcode{"php", "php", 1, 0, 0x08, false},
	"ora.#":    Opcode{"ora.#", "ora", 2, 1, 0x09, false},
	"asl.a":    Opcode{"asl.a", "asl", 1, 0, 0x0a, false},
	"phd":      Opcode{"phd", "phd", 1, 0, 0x0b, false},
	"tsb":      Opcode{"tsb", "tsb", 3, 0, 0x0c, false},
	"ora":      Opcode{"ora", "ora", 3, 1, 0x0d, false},
	"asl":      Opcode{"asl", "asl", 3, 1, 0x0e, false},
	"ora.l":    Opcode{"ora.l", "ora", 4, 1, 0x0f, false},
	"bpl":      Opcode{"bpl", "bpl", 2, 1, 0x10, false},
	"ora.diy":  Opcode{"ora.diy", "ora", 2, 1, 0x11, false},
	"ora.di":   Opcode{"ora.di", "ora", 2, 1, 0x12, false},
	"ora.siy":  Opcode{"ora.siy", "ora", 2, 1, 0x13, false},
	"trb.d":    Opcode{"trb.d", "trb", 2, 1, 0x14, false},
	"ora.dx":   Opcode{"ora.dx", "ora", 2, 1, 0x15, false},
	"asl.dx":   Opcode{"asl.dx", "asl", 2, 1, 0x16, false},
	"ora.dily": Opcode{"ora.dily", "ora", 2, 1, 0x17, false},
	"clc":      Opcode{"clc", "clc", 1, 0, 0x18, false},
	"ora.y":    Opcode{"ora.y", "ora", 3, 1, 0x19, false},
	"inc.a":    Opcode{"inc.a", "inc", 1, 0, 0x1a, false},
	"tcs":      Opcode{"tcs", "tcs", 1, 0, 0x1b, false},
	"trb":      Opcode{"trb", "trb", 3, 1, 0x1c, false},
	"ora.x":    Opcode{"ora.x", "ora", 3, 1, 0x1d, false},
	"asl.x":    Opcode{"asl.x", "asl", 3, 1, 0x1e, false},
	"ora.lx":   Opcode{"ora.lx", "ora", 4, 1, 0x1f, false},
	"jsr":      Opcode{"jsr", "jsr", 3, 1, 0x20, false},
	"and.dxi":  Opcode{"and.dxi", "and", 2, 1, 0x21, false},
	"jsr.l":    Opcode{"jsr.l", "jml", 4, 1, 0x22, false},
	"and.s":    Opcode{"and.s", "and", 2, 1, 0x23, false},
	"bit.d":    Opcode{"bit.d", "bit", 2, 1, 0x24, false},
	"and.d":    Opcode{"and.d", "and", 2, 1, 0x25, false},
	"rol.d":    Opcode{"rol.d", "rol", 2, 1, 0x26, false},
	"and.dil":  Opcode{"and.dil", "and", 2, 1, 0x27, false},
	"plp":      Opcode{"plp", "plp", 1, 0, 0x28, false},
	"and.#":    Opcode{"and.#", "and", 2, 1, 0x29, false},
	"rol.a":    Opcode{"rol.a", "rol", 1, 0, 0x2a, false},
	"pld":      Opcode{"pld", "pld", 1, 0, 0x2b, false},
	"bit":      Opcode{"bit", "bit", 3, 1, 0x2c, false},
	"and":      Opcode{"and", "and", 3, 1, 0x2d, false},
	"rol":      Opcode{"rol", "rol", 3, 1, 0x2e, false},
	"and.l":    Opcode{"and.l", "and", 4, 1, 0x2f, false},
	"bmi":      Opcode{"bmi", "bmi", 2, 1, 0x30, false},
	"and.diy":  Opcode{"and.diy", "and", 2, 1, 0x31, false},
	"and.di":   Opcode{"and.di", "and", 2, 1, 0x32, false},
	"and.siy":  Opcode{"and.siy", "and", 2, 1, 0x33, false},
	"bit.dx":   Opcode{"bit.dx", "bit", 2, 1, 0x34, false},
	"and.dx":   Opcode{"and.dx", "and", 2, 1, 0x35, false},
	"rol.dx":   Opcode{"rol.dx", "rol", 2, 1, 0x36, false},
	"and.dily": Opcode{"and.dily", "and", 2, 1, 0x37, false},
	"sec":      Opcode{"sec", "sec", 1, 0, 0x38, false},
	"and.y":    Opcode{"and.y", "and", 3, 1, 0x39, false},
	"dec.a":    Opcode{"dec.a", "dec", 1, 0, 0x3a, false},
	"tsc":      Opcode{"tsc", "tsc", 1, 0, 0x3b, false},
	"bit.x":    Opcode{"bit.x", "bit", 3, 1, 0x3c, false},
	"and.x":    Opcode{"and.x", "and", 3, 1, 0x3d, false},
	"rol.x":    Opcode{"rol.x", "rol", 3, 1, 0x3e, false},
	"and.lx":   Opcode{"and.lx", "and", 4, 1, 0x3f, false},
	"rti":      Opcode{"rti", "rti", 1, 0, 0x40, false},
	"eor.dxi":  Opcode{"eor.dxi", "eor", 2, 1, 0x41, false},
	"wdm":      Opcode{"wdm", "wdm", 2, 1, 0x42, false},
	"eor.s":    Opcode{"eor.s", "eor", 2, 1, 0x43, false},
	"mvp":      Opcode{"mvp", "mvp", 3, 2, 0x44, false}, // takes two operands
	"eor.d":    Opcode{"eor.d", "eor", 2, 1, 0x45, false},
	"lsr.d":    Opcode{"lsr.d", "lsr", 2, 1, 0x46, false},
	"eor.dil":  Opcode{"eor.dil", "eor", 2, 1, 0x47, false},
	"pha":      Opcode{"pha", "pha", 1, 0, 0x48, false},
	"eor.#":    Opcode{"eor.#", "eor", 2, 1, 0x49, false},
	"lsr.a":    Opcode{"lsr.a", "lsr", 1, 0, 0x4a, false},
	"phk":      Opcode{"phk", "phk", 1, 0, 0x4b, false},
	"jmp":      Opcode{"jmp", "jmp", 3, 1, 0x4c, false},
	"eor":      Opcode{"eor", "eor", 3, 1, 0x4d, false},
	"lsr":      Opcode{"lsr", "lsr", 3, 1, 0x4e, false},
	"eor.l":    Opcode{"eor.l", "eor", 4, 1, 0x4f, false},
	"bvc":      Opcode{"bvc", "bvc", 2, 1, 0x50, false},
	"eor.diy":  Opcode{"eor.diy", "eor", 2, 1, 0x51, false},
	"eor.di":   Opcode{"eor.di", "eor", 2, 1, 0x52, false},
	"eor.siy":  Opcode{"eor.siy", "eor", 2, 1, 0x53, false},
	"mvn":      Opcode{"mvn", "mvn", 3, 2, 0x54, false}, // takes two operands
	"eor.dx":   Opcode{"eor.dx", "eor", 2, 1, 0x55, false},
	"lsr.dx":   Opcode{"lsr.dx", "lsr", 2, 1, 0x56, false},
	"eor.dily": Opcode{"eor.dily", "eor", 2, 1, 0x57, false},
	"cli":      Opcode{"cli", "cli", 1, 0, 0x58, false},
	"eor.y":    Opcode{"eor.y", "eor", 3, 1, 0x59, false},
	"phy":      Opcode{"phy", "phy", 1, 0, 0x5a, false},
	"tcd":      Opcode{"tcd", "tcd", 1, 0, 0x5b, false},
	"jmp.l":    Opcode{"jmp.l", "jml", 4, 1, 0x5c, false},
	"eor.x":    Opcode{"eor.x", "eor", 3, 1, 0x5d, false},
	"lsr.x":    Opcode{"lsr.x", "lsr", 3, 1, 0x5e, false},
	"eor.lx":   Opcode{"eor.lx", "eor", 4, 1, 0x5f, false},
	"rts":      Opcode{"rts", "rts", 1, 0, 0x60, false},
	"adc.dxi":  Opcode{"adc.dxi", "adc", 2, 1, 0x61, false},
	"phe.r":    Opcode{"phe.r", "per", 2, 1, 0x62, false},
	"adc.s":    Opcode{"adc.s", "adc", 2, 1, 0x63, false},
	"stz.d":    Opcode{"stz.d", "stz", 2, 1, 0x64, false},
	"adc.d":    Opcode{"adc.d", "adc", 2, 1, 0x65, false},
	"ror.d":    Opcode{"ror.d", "ror", 2, 1, 0x66, false},
	"adc.dil":  Opcode{"adc.dil", "adc", 2, 1, 0x67, false},
	"pla":      Opcode{"pla", "pla", 1, 0, 0x68, false},
	"adc.#":    Opcode{"adc.#", "adc", 2, 1, 0x69, false},
	"ror.a":    Opcode{"ror.a", "ror", 1, 0, 0x6a, false},
	"rts.l":    Opcode{"rts.l", "rtl", 1, 1, 0x6b, false},
	"jmp.i":    Opcode{"jmp.i", "jmp", 3, 1, 0x6c, false},
	"adc":      Opcode{"adc", "adc", 3, 1, 0x6d, false},
	"ror":      Opcode{"ror", "ror", 3, 1, 0x6e, false},
	"adc.l":    Opcode{"adc.l", "adv", 4, 1, 0x6f, false},
	"bvs":      Opcode{"bvs", "bvs", 2, 1, 0x70, false},
	"adc.diy":  Opcode{"adc.diy", "adc", 2, 1, 0x71, false},
	"adc.di":   Opcode{"adc.di", "adc", 2, 1, 0x72, false},
	"adc.siy":  Opcode{"adc.siy", "adc", 2, 1, 0x73, false},
	"stz.dx":   Opcode{"stz.dx", "stz", 2, 1, 0x74, false},
	"adc.dx":   Opcode{"adc.dx", "adc", 2, 1, 0x75, false},
	"ror.dx":   Opcode{"ror.dx", "ror", 2, 1, 0x76, false},
	"adc.dily": Opcode{"adc.dily", "adc", 2, 1, 0x77, false},
	"sei":      Opcode{"sei", "sei", 1, 0, 0x78, false},
	"adc.y":    Opcode{"adc.y", "adc", 3, 1, 0x79, false},
	"ply":      Opcode{"ply", "ply", 1, 0, 0x7a, false},
	"tdc":      Opcode{"tdc", "tdc", 1, 0, 0x7b, false},
	"jmp.xi":   Opcode{"jmp.xi", "jmp", 3, 1, 0x7c, false},
	"adc.x":    Opcode{"adc.x", "adc", 3, 1, 0x7d, false},
	"ror.x":    Opcode{"ror.x", "ror", 3, 1, 0x7e, false},
	"adc.lx":   Opcode{"adc.lx", "adc", 4, 1, 0x7f, false},
	"bra":      Opcode{"bra", "bra", 2, 1, 0x80, false},
	"sta.dxi":  Opcode{"sta.dxi", "sta", 2, 1, 0x81, false},
	"bra.l":    Opcode{"bra.l", "brl", 3, 1, 0x82, false},
	"sta.s":    Opcode{"sta.s", "sta", 2, 1, 0x83, false},
	"sty.d":    Opcode{"sty.d", "sty", 2, 1, 0x84, false},
	"sta.d":    Opcode{"sta.d", "sta", 2, 1, 0x85, false},
	"stx.d":    Opcode{"stx.d", "stx", 2, 1, 0x86, false},
	"sta.dil":  Opcode{"sta.dil", "sta", 2, 1, 0x87, false},
	"dey":      Opcode{"dey", "dey", 1, 0, 0x88, false},
	"bit.#":    Opcode{"bit.#", "bit", 2, 1, 0x89, false},
	"txa":      Opcode{"txa", "txa", 1, 0, 0x8a, false},
	"phb":      Opcode{"phb", "phb", 1, 0, 0x8b, false},
	"sty":      Opcode{"sty", "sty", 3, 1, 0x8c, false},
	"sta":      Opcode{"sta", "sta", 3, 1, 0x8d, false},
	"stx":      Opcode{"stx", "stx", 3, 1, 0x8e, false},
	"sta.l":    Opcode{"sta.l", "sta", 4, 1, 0x8f, false},
	"bcc":      Opcode{"bcc", "bbc", 2, 1, 0x90, false},
	"sta.diy":  Opcode{"sta.diy", "sta", 2, 1, 0x91, false},
	"sta.di":   Opcode{"sta.di", "sta", 2, 1, 0x92, false},
	"sta.siy":  Opcode{"sta.siy", "sta", 2, 1, 0x93, false},
	"sty.dx":   Opcode{"sty.dx", "sty", 2, 1, 0x94, false},
	"sta.dx":   Opcode{"sta.dx", "sta", 2, 1, 0x95, false},
	"stx.dy":   Opcode{"stx.dy", "stx", 2, 1, 0x96, false},
	"sta.dily": Opcode{"sta.dily", "sta", 2, 1, 0x97, false},
	"tya":      Opcode{"tya", "tya", 1, 0, 0x98, false},
	"sta.y":    Opcode{"sta.y", "sta", 3, 1, 0x99, false},
	"txs":      Opcode{"txs", "txs", 1, 0, 0x9a, false},
	"txy":      Opcode{"txy", "txy", 1, 0, 0x9b, false},
	"stz":      Opcode{"stz", "stz", 3, 1, 0x9c, false},
	"sta.x":    Opcode{"sta.x", "sta", 3, 1, 0x9d, false},
	"stz.x":    Opcode{"stz.x", "stz", 3, 1, 0x9e, false},
	"sta.lx":   Opcode{"sta.lx", "sta", 4, 1, 0x9f, false},
	"ldy.#":    Opcode{"ldy.#", "ldy", 2, 1, 0xa0, false},
	"lda.dxi":  Opcode{"lda.dxi", "lda", 2, 1, 0xa1, false},
	"ldx.#":    Opcode{"ldx.#", "ldx", 2, 1, 0xa2, false},
	"lda.s":    Opcode{"lda.s", "lda", 2, 1, 0xa3, false},
	"ldy.d":    Opcode{"ldy.d", "ldy", 2, 1, 0xa4, false},
	"lda.d":    Opcode{"lda.d", "lda", 2, 1, 0xa5, false},
	"ldx.d":    Opcode{"ldx.d", "ldx", 2, 1, 0xa6, false},
	"lda.dil":  Opcode{"lda.dil", "lda", 2, 1, 0xa7, false},
	"tay":      Opcode{"tay", "tay", 1, 0, 0xa8, false},
	"lda.#":    Opcode{"lda.#", "lda", 2, 1, 0xa9, false},
	"tax":      Opcode{"tax", "tax", 1, 0, 0xaa, false},
	"plb":      Opcode{"plb", "plb", 1, 0, 0xab, false},
	"ldy":      Opcode{"ldy", "ldy", 3, 1, 0xac, false},
	"lda":      Opcode{"lda", "lda", 3, 1, 0xad, false},
	"ldx":      Opcode{"ldx", "ldx", 3, 1, 0xae, false},
	"lda.l":    Opcode{"lda.l", "lda", 4, 1, 0xaf, false},
	"bcs":      Opcode{"bcs", "bcs", 2, 1, 0xb0, false},
	"lda.diy":  Opcode{"lda.diy", "lda", 2, 1, 0xb1, false},
	"lda.di":   Opcode{"lda.di", "lda", 2, 1, 0xb2, false},
	"lda.siy":  Opcode{"lda.siy", "lda", 2, 1, 0xb3, false},
	"ldy.dx":   Opcode{"ldy.dx", "ldy", 2, 1, 0xb4, false},
	"lda.dx":   Opcode{"lda.dx", "lda", 2, 1, 0xb5, false},
	"ldx.dy":   Opcode{"ldx.dy", "ldx", 2, 1, 0xb6, false},
	"lda.dily": Opcode{"lda.dily", "lda", 2, 1, 0xb7, false},
	"clv":      Opcode{"clv", "clv", 1, 0, 0xb8, false},
	"lda.y":    Opcode{"lda.y", "lda", 3, 1, 0xb9, false},
	"tsx":      Opcode{"tsx", "tsx", 1, 0, 0xba, false},
	"tyx":      Opcode{"tyx", "tyx", 1, 0, 0xbb, false},
	"ldy.x":    Opcode{"ldy.x", "ldy", 3, 1, 0xbc, false},
	"lda.x":    Opcode{"lda.x", "lda", 3, 1, 0xbd, false},
	"ldx.y":    Opcode{"ldx.y", "ldx", 3, 1, 0xbe, false},
	"lda.lx":   Opcode{"lda.lx", "lda", 4, 1, 0xbf, false},
	"cpy.#":    Opcode{"cpy.#", "cpy", 2, 1, 0xc0, false},
	"cmp.dxi":  Opcode{"cmp.dxi", "cmp", 2, 1, 0xc1, false},
	"rep":      Opcode{"rep", "rep", 2, 1, 0xc2, false},
	"cmp.s":    Opcode{"cmp.s", "cmp", 2, 1, 0xc3, false},
	"cpy.d":    Opcode{"cpy.d", "cpy", 2, 1, 0xc4, false},
	"cmp.d":    Opcode{"cmp.d", "cmp", 2, 1, 0xc5, false},
	"dec.d":    Opcode{"dec.d", "dec", 2, 1, 0xc6, false},
	"cmp.dil":  Opcode{"cmp.dil", "cmp", 2, 1, 0xc7, false},
	"iny":      Opcode{"iny", "iny", 1, 0, 0xc8, false},
	"cmp.#":    Opcode{"cmp.#", "cmp", 2, 1, 0xc9, false},
	"dex":      Opcode{"dex", "dex", 1, 0, 0xca, false},
	"wai":      Opcode{"wai", "wai", 1, 0, 0xcb, false},
	"cpy":      Opcode{"cpy", "cpy", 3, 0, 0xcc, false},
	"cmp":      Opcode{"cmp", "cmp", 3, 0, 0xcd, false},
	"dec":      Opcode{"dec", "dec", 3, 0, 0xce, false},
	"cmp.l":    Opcode{"cmp.l", "cmp", 4, 0, 0xcf, false},
	"bne":      Opcode{"bne", "bne", 2, 1, 0xd0, false},
	"cmp.diy":  Opcode{"cmp.diy", "cmp", 2, 1, 0xd1, false},
	"cmp.di":   Opcode{"cmp.di", "cmp", 2, 1, 0xd2, false},
	"cmp.siy":  Opcode{"cmp.siy", "cmp", 2, 1, 0xd3, false},
	"phe.d":    Opcode{"phe.d", "pei", 2, 1, 0xd4, false},
	"cmp.dx":   Opcode{"cmp.dx", "cmp", 2, 1, 0xd5, false},
	"dec.dx":   Opcode{"dec.dx", "dec", 2, 1, 0xd6, false},
	"cmp.dily": Opcode{"cmp.dily", "cmp", 2, 1, 0xd7, false},
	"cld":      Opcode{"cld", "cld", 1, 0, 0xd8, false},
	"cmp.y":    Opcode{"cmp.y", "cmp", 3, 1, 0xd9, false},
	"phx":      Opcode{"phx", "phx", 1, 0, 0xda, false},
	"stp":      Opcode{"stp", "stp", 1, 0, 0xdb, false},
	"jmp.il":   Opcode{"jmp.il", "jml", 3, 1, 0xdc, false},
	"cmp.x":    Opcode{"cmp.x", "cmp", 3, 1, 0xdd, false},
	"dec.x":    Opcode{"dec.x", "dec", 3, 1, 0xde, false},
	"cmp.lx":   Opcode{"cmp.lx", "cmp", 4, 1, 0xdf, false},
	"cpx.#":    Opcode{"cpx.#", "cpx", 2, 1, 0xe0, false},
	"sbc.dxi":  Opcode{"sbc.dxi", "sbc", 2, 1, 0xe1, false},
	"sep":      Opcode{"sep", "sep", 2, 1, 0xe2, false},
	"sbc.s":    Opcode{"sbc.s", "sbc", 2, 1, 0xe3, false},
	"cpx.d":    Opcode{"cpx.d", "cpx", 2, 1, 0xe4, false},
	"sbc.d":    Opcode{"sbc.d", "sbc", 2, 1, 0xe5, false},
	"inc.d":    Opcode{"inc.d", "inc", 2, 1, 0xe6, false},
	"sbc.dil":  Opcode{"sbc.dil", "sbc", 2, 1, 0xe7, false},
	"inx":      Opcode{"inx", "inx", 1, 0, 0xe8, false},
	"sbc.#":    Opcode{"sbc.#", "sbc", 2, 1, 0xe9, false},
	"nop":      Opcode{"nop", "nop", 1, 0, 0xea, false},
	"xba":      Opcode{"xba", "xba", 1, 0, 0xeb, false},
	"cpx":      Opcode{"cpx", "cpx", 3, 1, 0xec, false},
	"sbc":      Opcode{"sbc", "sbc", 3, 1, 0xed, false},
	"inc":      Opcode{"inc", "inc", 3, 1, 0xee, false},
	"sbc.l":    Opcode{"sbc.l", "sbc", 4, 1, 0xef, false},
	"beq":      Opcode{"beq", "beq", 2, 1, 0xf0, false},
	"sbc.diy":  Opcode{"sbc.diy", "sbc", 2, 1, 0xf1, false},
	"sbc.di":   Opcode{"sbc.di", "sbc", 2, 1, 0xf2, false},
	"sbc.siy":  Opcode{"sbc.siy", "sbc", 2, 1, 0xf3, false},
	"phe.#":    Opcode{"phe.#", "pea", 3, 1, 0xf4, false},
	"sbc.dx":   Opcode{"sbc.dx", "sbc", 2, 1, 0xf5, false},
	"inc.dx":   Opcode{"inc.dx", "inc", 2, 1, 0xf6, false},
	"sbc.dily": Opcode{"sbc.dily", "sbc", 2, 1, 0xf7, false},
	"sed":      Opcode{"sed", "sed", 1, 0, 0xf8, false},
	"sbc.y":    Opcode{"sbc.y", "sbc", 3, 1, 0xf9, false},
	"plx":      Opcode{"plx", "plx", 1, 0, 0xfa, false},
	"xce":      Opcode{"xce", "xce", 1, 0, 0xfb, false},
	"jsr.xi":   Opcode{"jsr.xi", "jsr", 3, 1, 0xfc, false},
	"sbc.x":    Opcode{"sbc.x", "sbc", 3, 1, 0xfd, false},
	"inc.x":    Opcode{"inc.x", "inc", 3, 1, 0xfe, false},
	"sbc.lx":   Opcode{"sbc.lx", "sbc", 4, 1, 0xff, false},
}
