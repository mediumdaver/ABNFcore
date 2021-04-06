/* Implements ABNF RFC 5234 core elements.
 * License: MIT, as distributed with this source code.
 */

package abnfcore

func RuneIsALPHA(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func RuneIsBIT(c rune) bool {
	return c == '0' || c == '1'
}

func RuneIsCHAR(c rune) bool {
	return c >= 1 && c <= '\x7f'
}

func RuneIsCR(c rune) bool {
	return c == '\r'
}

func StringIsCRLF(s string) bool {
	return s == "\r\n"
}

func RuneIsCTL(c rune) bool {
	return c == '\x7f' || (c >= '\x00' && c <= '\x1f')
}

func RuneIsDIGIT(c rune) bool {
	return c >= '0' && c <= '9'
}

func RuneIsDQUOTE(c rune) bool {
	return c == '"'
}

func RuneIsHEXDIG(c rune) bool {
	return RuneIsDIGIT(c) || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f')
}

func RuneIsHTAB(c rune) bool {
	return c == '\t'
}

func RuneIsLF(c rune) bool {
	return c == '\n'
}

func StringIsLWSP(s string) bool {
	inCRLF := false
	for i := 0; i < len(s); i++ {
		if !inCRLF && RuneIsCR(rune(s[i])) {
			inCRLF = true
		} else if inCRLF && RuneIsLF(rune(s[i])) {
			inCRLF = false
		} else if !RuneIsWSP(rune(s[i])) {
			return false
		}
	}
	if inCRLF {
		return false
	}

	return true
}

func RuneIsOCTET(c rune) bool {
	return c >= '\x00' && c <= '\xff'
}

func RuneIsSP(c rune) bool {
	return c == ' '
}

func RuneIsVCHAR(c rune) bool {
	return c >= '!' && c <= '~'
}

func RuneIsWSP(c rune) bool {
	return RuneIsSP(c) || RuneIsHTAB(c)
}
