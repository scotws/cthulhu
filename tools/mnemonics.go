"brk": Opcode{"brk", 2, X, 0x00, false},
"ora.dxi": Opcode{"ora.dxi", 2, X, 0x01, false},
"cop": Opcode{"cop", 2, X, 0x02, false},
"ora.s": Opcode{"ora.s", 2, X, 0x03, false},
"tsb.d": Opcode{"tsb.d", 2, X, 0x04, false},
"ora.d": Opcode{"ora.d", 2, X, 0x05, false},
"asl.d": Opcode{"asl.d", 2, X, 0x06, false},
"ora.dil": Opcode{"ora.dil", 2, X, 0x07, false},
"php": Opcode{"php", 1, X, 0x08, false},
"ora.#": Opcode{"ora.#", 2, X, 0x09, false},
"asl.a": Opcode{"asl.a", 1, X, 0x0a, false},
"phd": Opcode{"phd", 1, X, 0x0b, false},
"tsb": Opcode{"tsb", 3, X, 0x0c, false},
"ora": Opcode{"ora", 3, X, 0x0d, false},
"asl": Opcode{"asl", 3, X, 0x0e, false},
"ora.l": Opcode{"ora.l", 4, X, 0x0f, false},
"bpl": Opcode{"bpl", 2, X, 0x10, false},
"ora.diy": Opcode{"ora.diy", 2, X, 0x11, false},
"ora.di": Opcode{"ora.di", 2, X, 0x12, false},
"ora.siy": Opcode{"ora.siy", 2, X, 0x13, false},
"trb.d": Opcode{"trb.d", 2, X, 0x14, false},
"ora.dx": Opcode{"ora.dx", 2, X, 0x15, false},
"asl.dx": Opcode{"asl.dx", 2, X, 0x16, false},
"ora.dily": Opcode{"ora.dily", 2, X, 0x17, false},
"clc": Opcode{"clc", 1, X, 0x18, false},
"ora.y": Opcode{"ora.y", 3, X, 0x19, false},
"inc.a": Opcode{"inc.a", 1, X, 0x1a, false},
"tcs": Opcode{"tcs", 1, X, 0x1b, false},
"trb": Opcode{"trb", 3, X, 0x1c, false},
"ora.x": Opcode{"ora.x", 3, X, 0x1d, false},
"asl.x": Opcode{"asl.x", 3, X, 0x1e, false},
"ora.lx": Opcode{"ora.lx", 4, X, 0x1f, false},
"jsr": Opcode{"jsr", 3, X, 0x20, false},
"and.dxi": Opcode{"and.dxi", 2, X, 0x21, false},
"jsr.l": Opcode{"jsr.l", 4, X, 0x22, false},
"and.s": Opcode{"and.s", 2, X, 0x23, false},
"bit.d": Opcode{"bit.d", 2, X, 0x24, false},
"and.d": Opcode{"and.d", 2, X, 0x25, false},
"rol.d": Opcode{"rol.d", 2, X, 0x26, false},
"and.dil": Opcode{"and.dil", 2, X, 0x27, false},
"plp": Opcode{"plp", 1, X, 0x28, false},
"and.#": Opcode{"and.#", 2, X, 0x29, false},
"rol.a": Opcode{"rol.a", 1, X, 0x2a, false},
"pld": Opcode{"pld", 1, X, 0x2b, false},
"bit": Opcode{"bit", 3, X, 0x2c, false},
"and": Opcode{"and", 3, X, 0x2d, false},
"rol": Opcode{"rol", 3, X, 0x2e, false},
"and.l": Opcode{"and.l", 4, X, 0x2f, false},
"bmi": Opcode{"bmi", 2, X, 0x30, false},
"and.diy": Opcode{"and.diy", 2, X, 0x31, false},
"and.di": Opcode{"and.di", 2, X, 0x32, false},
"and.siy": Opcode{"and.siy", 2, X, 0x33, false},
"bit.dx": Opcode{"bit.dx", 2, X, 0x34, false},
"and.dx": Opcode{"and.dx", 2, X, 0x35, false},
"rol.dx": Opcode{"rol.dx", 2, X, 0x36, false},
"and.dily": Opcode{"and.dily", 2, X, 0x37, false},
"sec": Opcode{"sec", 1, X, 0x38, false},
"and.y": Opcode{"and.y", 3, X, 0x39, false},
"dec.a": Opcode{"dec.a", 1, X, 0x3a, false},
"tsc": Opcode{"tsc", 1, X, 0x3b, false},
"bit.x": Opcode{"bit.x", 3, X, 0x3c, false},
"and.x": Opcode{"and.x", 3, X, 0x3d, false},
"rol.x": Opcode{"rol.x", 3, X, 0x3e, false},
"and.lx": Opcode{"and.lx", 4, X, 0x3f, false},
"rti": Opcode{"rti", 1, X, 0x40, false},
"eor.dxi": Opcode{"eor.dxi", 2, X, 0x41, false},
"wdm": Opcode{"wdm", 2, X, 0x42, false},
"eor.s": Opcode{"eor.s", 2, X, 0x43, false},
"mvp": Opcode{"mvp", 3, X, 0x44, false},
"eor.d": Opcode{"eor.d", 2, X, 0x45, false},
"lsr.d": Opcode{"lsr.d", 2, X, 0x46, false},
"eor.dil": Opcode{"eor.dil", 2, X, 0x47, false},
"pha": Opcode{"pha", 1, X, 0x48, false},
"eor.#": Opcode{"eor.#", 2, X, 0x49, false},
"lsr.a": Opcode{"lsr.a", 1, X, 0x4a, false},
"phk": Opcode{"phk", 1, X, 0x4b, false},
"jmp": Opcode{"jmp", 3, X, 0x4c, false},
"eor": Opcode{"eor", 3, X, 0x4d, false},
"lsr": Opcode{"lsr", 3, X, 0x4e, false},
"eor.l": Opcode{"eor.l", 4, X, 0x4f, false},
"bvc": Opcode{"bvc", 2, X, 0x50, false},
"eor.diy": Opcode{"eor.diy", 2, X, 0x51, false},
"eor.di": Opcode{"eor.di", 2, X, 0x52, false},
"eor.siy": Opcode{"eor.siy", 2, X, 0x53, false},
"mvn": Opcode{"mvn", 3, X, 0x54, false},
"eor.dx": Opcode{"eor.dx", 2, X, 0x55, false},
"lsr.dx": Opcode{"lsr.dx", 2, X, 0x56, false},
"eor.dily": Opcode{"eor.dily", 2, X, 0x57, false},
"cli": Opcode{"cli", 1, X, 0x58, false},
"eor.y": Opcode{"eor.y", 3, X, 0x59, false},
"phy": Opcode{"phy", 1, X, 0x5a, false},
"tcd": Opcode{"tcd", 1, X, 0x5b, false},
"jmp.l": Opcode{"jmp.l", 4, X, 0x5c, false},
"eor.x": Opcode{"eor.x", 3, X, 0x5d, false},
"lsr.x": Opcode{"lsr.x", 3, X, 0x5e, false},
"eor.lx": Opcode{"eor.lx", 4, X, 0x5f, false},
"rts": Opcode{"rts", 1, X, 0x60, false},
"adc.dxi": Opcode{"adc.dxi", 2, X, 0x61, false},
"phe.r": Opcode{"phe.r", 3, X, 0x62, false},
"adc.s": Opcode{"adc.s", 2, X, 0x63, false},
"stz.d": Opcode{"stz.d", 2, X, 0x64, false},
"adc.d": Opcode{"adc.d", 2, X, 0x65, false},
"ror.d": Opcode{"ror.d", 2, X, 0x66, false},
"adc.dil": Opcode{"adc.dil", 2, X, 0x67, false},
"pla": Opcode{"pla", 1, X, 0x68, false},
"adc.#": Opcode{"adc.#", 2, X, 0x69, false},
"ror.a": Opcode{"ror.a", 1, X, 0x6a, false},
"rts.l": Opcode{"rts.l", 1, X, 0x6b, false},
"jmp.i": Opcode{"jmp.i", 3, X, 0x6c, false},
"adc": Opcode{"adc", 3, X, 0x6d, false},
"ror": Opcode{"ror", 3, X, 0x6e, false},
"adc.l": Opcode{"adc.l", 4, X, 0x6f, false},
"bvs": Opcode{"bvs", 2, X, 0x70, false},
"adc.diy": Opcode{"adc.diy", 2, X, 0x71, false},
"adc.di": Opcode{"adc.di", 2, X, 0x72, false},
"adc.siy": Opcode{"adc.siy", 2, X, 0x73, false},
"stz.dx": Opcode{"stz.dx", 2, X, 0x74, false},
"adc.dx": Opcode{"adc.dx", 2, X, 0x75, false},
"ror.dx": Opcode{"ror.dx", 2, X, 0x76, false},
"adc.dily": Opcode{"adc.dily", 2, X, 0x77, false},
"sei": Opcode{"sei", 1, X, 0x78, false},
"adc.y": Opcode{"adc.y", 3, X, 0x79, false},
"ply": Opcode{"ply", 1, X, 0x7a, false},
"tdc": Opcode{"tdc", 1, X, 0x7b, false},
"jmp.xi": Opcode{"jmp.xi", 3, X, 0x7c, false},
"adc.x": Opcode{"adc.x", 3, X, 0x7d, false},
"ror.x": Opcode{"ror.x", 3, X, 0x7e, false},
"adc.lx": Opcode{"adc.lx", 4, X, 0x7f, false},
"bra": Opcode{"bra", 2, X, 0x80, false},
"sta.dxi": Opcode{"sta.dxi", 2, X, 0x81, false},
"bra.l": Opcode{"bra.l", 3, X, 0x82, false},
"sta.s": Opcode{"sta.s", 2, X, 0x83, false},
"sty.d": Opcode{"sty.d", 2, X, 0x84, false},
"sta.d": Opcode{"sta.d", 2, X, 0x85, false},
"stx.d": Opcode{"stx.d", 2, X, 0x86, false},
"sta.dil": Opcode{"sta.dil", 2, X, 0x87, false},
"dey": Opcode{"dey", 1, X, 0x88, false},
"bit.#": Opcode{"bit.#", 2, X, 0x89, false},
"txa": Opcode{"txa", 1, X, 0x8a, false},
"phb": Opcode{"phb", 1, X, 0x8b, false},
"sty": Opcode{"sty", 3, X, 0x8c, false},
"sta": Opcode{"sta", 3, X, 0x8d, false},
"stx": Opcode{"stx", 3, X, 0x8e, false},
"sta.l": Opcode{"sta.l", 4, X, 0x8f, false},
"bcc": Opcode{"bcc", 2, X, 0x90, false},
"sta.diy": Opcode{"sta.diy", 2, X, 0x91, false},
"sta.di": Opcode{"sta.di", 2, X, 0x92, false},
"sta.siy": Opcode{"sta.siy", 2, X, 0x93, false},
"sty.dx": Opcode{"sty.dx", 2, X, 0x94, false},
"sta.dx": Opcode{"sta.dx", 2, X, 0x95, false},
"stx.dy": Opcode{"stx.dy", 2, X, 0x96, false},
"sta.dily": Opcode{"sta.dily", 2, X, 0x97, false},
"tya": Opcode{"tya", 1, X, 0x98, false},
"sta.y": Opcode{"sta.y", 3, X, 0x99, false},
"txs": Opcode{"txs", 1, X, 0x9a, false},
"txy": Opcode{"txy", 1, X, 0x9b, false},
"stz": Opcode{"stz", 3, X, 0x9c, false},
"sta.x": Opcode{"sta.x", 3, X, 0x9d, false},
"stz.x": Opcode{"stz.x", 3, X, 0x9e, false},
"sta.lx": Opcode{"sta.lx", 4, X, 0x9f, false},
"ldy.#": Opcode{"ldy.#", 2, X, 0xa0, false},
"lda.dxi": Opcode{"lda.dxi", 2, X, 0xa1, false},
"ldx.#": Opcode{"ldx.#", 2, X, 0xa2, false},
"lda.s": Opcode{"lda.s", 2, X, 0xa3, false},
"ldy.d": Opcode{"ldy.d", 2, X, 0xa4, false},
"lda.d": Opcode{"lda.d", 2, X, 0xa5, false},
"ldx.d": Opcode{"ldx.d", 2, X, 0xa6, false},
"lda.dil": Opcode{"lda.dil", 2, X, 0xa7, false},
"tay": Opcode{"tay", 1, X, 0xa8, false},
"lda.#": Opcode{"lda.#", 2, X, 0xa9, false},
"tax": Opcode{"tax", 1, X, 0xaa, false},
"plb": Opcode{"plb", 1, X, 0xab, false},
"ldy": Opcode{"ldy", 3, X, 0xac, false},
"lda": Opcode{"lda", 3, X, 0xad, false},
"ldx": Opcode{"ldx", 3, X, 0xae, false},
"lda.l": Opcode{"lda.l", 4, X, 0xaf, false},
"bcs": Opcode{"bcs", 2, X, 0xb0, false},
"lda.diy": Opcode{"lda.diy", 2, X, 0xb1, false},
"lda.di": Opcode{"lda.di", 2, X, 0xb2, false},
"lda.siy": Opcode{"lda.siy", 2, X, 0xb3, false},
"ldy.dx": Opcode{"ldy.dx", 2, X, 0xb4, false},
"lda.dx": Opcode{"lda.dx", 2, X, 0xb5, false},
"ldx.dy": Opcode{"ldx.dy", 2, X, 0xb6, false},
"lda.dily": Opcode{"lda.dily", 2, X, 0xb7, false},
"clv": Opcode{"clv", 1, X, 0xb8, false},
"lda.y": Opcode{"lda.y", 3, X, 0xb9, false},
"tsx": Opcode{"tsx", 1, X, 0xba, false},
"tyx": Opcode{"tyx", 1, X, 0xbb, false},
"ldy.x": Opcode{"ldy.x", 3, X, 0xbc, false},
"lda.x": Opcode{"lda.x", 3, X, 0xbd, false},
"ldx.y": Opcode{"ldx.y", 3, X, 0xbe, false},
"lda.lx": Opcode{"lda.lx", 4, X, 0xbf, false},
"cpy.#": Opcode{"cpy.#", 2, X, 0xc0, false},
"cmp.dxi": Opcode{"cmp.dxi", 2, X, 0xc1, false},
"rep": Opcode{"rep", 2, X, 0xc2, false},
"cmp.s": Opcode{"cmp.s", 2, X, 0xc3, false},
"cpy.d": Opcode{"cpy.d", 2, X, 0xc4, false},
"cmp.d": Opcode{"cmp.d", 2, X, 0xc5, false},
"dec.d": Opcode{"dec.d", 2, X, 0xc6, false},
"cmp.dil": Opcode{"cmp.dil", 2, X, 0xc7, false},
"iny": Opcode{"iny", 1, X, 0xc8, false},
"cmp.#": Opcode{"cmp.#", 2, X, 0xc9, false},
"dex": Opcode{"dex", 1, X, 0xca, false},
"wai": Opcode{"wai", 1, X, 0xcb, false},
"cpy": Opcode{"cpy", 3, X, 0xcc, false},
"cmp": Opcode{"cmp", 3, X, 0xcd, false},
"dec": Opcode{"dec", 3, X, 0xce, false},
"cmp.l": Opcode{"cmp.l", 4, X, 0xcf, false},
"bne": Opcode{"bne", 2, X, 0xd0, false},
"cmp.diy": Opcode{"cmp.diy", 2, X, 0xd1, false},
"cmp.di": Opcode{"cmp.di", 2, X, 0xd2, false},
"cmp.siy": Opcode{"cmp.siy", 2, X, 0xd3, false},
"phe.d": Opcode{"phe.d", 2, X, 0xd4, false},
"cmp.dx": Opcode{"cmp.dx", 2, X, 0xd5, false},
"dec.dx": Opcode{"dec.dx", 2, X, 0xd6, false},
"cmp.dily": Opcode{"cmp.dily", 2, X, 0xd7, false},
"cld": Opcode{"cld", 1, X, 0xd8, false},
"cmp.y": Opcode{"cmp.y", 3, X, 0xd9, false},
"phx": Opcode{"phx", 1, X, 0xda, false},
"stp": Opcode{"stp", 1, X, 0xdb, false},
"jmp.il": Opcode{"jmp.il", 3, X, 0xdc, false},
"cmp.x": Opcode{"cmp.x", 3, X, 0xdd, false},
"dec.x": Opcode{"dec.x", 3, X, 0xde, false},
"cmp.lx": Opcode{"cmp.lx", 4, X, 0xdf, false},
"cpx.#": Opcode{"cpx.#", 2, X, 0xe0, false},
"sbc.dxi": Opcode{"sbc.dxi", 2, X, 0xe1, false},
"sep": Opcode{"sep", 2, X, 0xe2, false},
"sbc.s": Opcode{"sbc.s", 2, X, 0xe3, false},
"cpx.d": Opcode{"cpx.d", 2, X, 0xe4, false},
"sbc.d": Opcode{"sbc.d", 2, X, 0xe5, false},
"inc.d": Opcode{"inc.d", 2, X, 0xe6, false},
"sbc.dil": Opcode{"sbc.dil", 2, X, 0xe7, false},
"inx": Opcode{"inx", 1, X, 0xe8, false},
"sbc.#": Opcode{"sbc.#", 2, X, 0xe9, false},
"nop": Opcode{"nop", 1, X, 0xea, false},
"xba": Opcode{"xba", 1, X, 0xeb, false},
"cpx": Opcode{"cpx", 3, X, 0xec, false},
"sbc": Opcode{"sbc", 3, X, 0xed, false},
"inc": Opcode{"inc", 3, X, 0xee, false},
"sbc.l": Opcode{"sbc.l", 4, X, 0xef, false},
"beq": Opcode{"beq", 2, X, 0xf0, false},
"sbc.diy": Opcode{"sbc.diy", 2, X, 0xf1, false},
"sbc.di": Opcode{"sbc.di", 2, X, 0xf2, false},
"sbc.siy": Opcode{"sbc.siy", 2, X, 0xf3, false},
"phe.#": Opcode{"phe.#", 3, X, 0xf4, false},
"sbc.dx": Opcode{"sbc.dx", 2, X, 0xf5, false},
"inc.dx": Opcode{"inc.dx", 2, X, 0xf6, false},
"sbc.dily": Opcode{"sbc.dily", 2, X, 0xf7, false},
"sed": Opcode{"sed", 1, X, 0xf8, false},
"sbc.y": Opcode{"sbc.y", 3, X, 0xf9, false},
"plx": Opcode{"plx", 1, X, 0xfa, false},
"xce": Opcode{"xce", 1, X, 0xfb, false},
"jsr.xi": Opcode{"jsr.xi", 3, X, 0xfc, false},
"sbc.x": Opcode{"sbc.x", 3, X, 0xfd, false},
"inc.x": Opcode{"inc.x", 3, X, 0xfe, false},
"sbc.lx": Opcode{"sbc.lx", 4, X, 0xff, false},
