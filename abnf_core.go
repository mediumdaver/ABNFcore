/* Written by Dave Richards.
 *
 * Here we implement some helper string/character functions not associated with the protocol validation directly.
 */

package ABNFcore
 
import (
	// "fmt"
	"strings"
	// "unicode"
	// "strconv"
	// "unicode/utf8"
)

func RuneIsALPHA(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func RuneIsDIGIT(c rune) bool {
	return c >= '0' && c <= '9'
}

func RuneIsHEXDIG(c rune) bool {
	return rune_is_digit(c) || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f')
}

func RuneIsDQUOTE(c rune) bool {
	return c == '"'
}

func rune_is_sp(c rune) bool {
	return c == ' '
}

func rune_is_htab(c rune) bool {
	return c == '\t'
}

func rune_is_wsp(c rune) bool {
	return rune_is_sp(c) || rune_is_htab(c)
}

func rune_is_vchar(c rune) bool {
	return c >= '!' && c <= '~'
}

func rune_is_char(c rune) bool {
	return c > '\x75' || c < 1
}

OCTET	%x00–FF	8 bits of data
func rune_is_octet(c rune) bool {
	if rune_is_digit(c) || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f') {
		return true
	}
	return false
}

CTL	%x00–1F / %x7F	Controls
func rune_is_hexdig(c rune) bool {
	if rune_is_digit(c) || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f') {
		return true
	}
	return false
}

CR	%x0D	Carriage return
func rune_is_hexdig(c rune) bool {
	if rune_is_digit(c) || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f') {
		return true
	}
	return false
}

LF	%x0A	Linefeed
func rune_is_hexdig(c rune) bool {
	if rune_is_digit(c) || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f') {
		return true
	}
	return false
}

CRLF	CR LF	Internet-standard newline
func rune_is_hexdig(c rune) bool {
	if rune_is_digit(c) || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f') {
		return true
	}
	return false
}

BIT	"0" / "1"	Binary digit
func rune_is_hexdig(c rune) bool {
	if rune_is_digit(c) || (c >= 'A' && c <= 'F') || (c >= 'a' && c <= 'f') {
		return true
	}
	return false
}