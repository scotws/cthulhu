// Fragements moved from other parts of the code. This is temporary until we can
// actually write this program

package converter

func procWDCMne(rs []rune, mpu string) (int, int, bool) {
	var o int

	f := false

	// We allow uppercase letters for mnenomics because we are nice that
	// way, but internally all is lower case
	s0 := strings.ToLower(string(rs))

	// Any WDC opcode must be three characters long exactly, which means
	// that either the word is exactly three characters long or the fourth
	// letter is a blank
	if (len(s0) > 3 && s0[3] == ' ') || len(s0) == 3 {
		s1 := s0[0:3]
		_, ok := data.OpcodesWDC[mpu][s1]

		if ok {

			// To make life easier for the parser, we check to see
			// if this is an instruction that definitely doesn't
			// take any operands
			_, ok := data.MneWDC65816NoPara[s1]

			if ok {
				o = token.WDC_NOPARA
			} else {
				o = token.WDC
			}

			f = true
		}
	}
	return o, 3, f
}
