package abnfcore

import (
	"testing"
)

func Test_RuneIsALPHA(t *testing.T) {
	testData := map[rune]bool{
		'a':	true,
		'\x8a':	false,
		'%':	false,
		'*':	false,
		'G':	true,
		'7':	false,
		'\x00':	false,
		-1:		false,
	}

	for d, result := range testData {
		if RuneIsALPHA(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsBIT(t *testing.T) {
	testData := map[rune]bool{
		'a':	false,
		'\x8a':	false,
		'%':	false,
		'*':	false,
		'0':	true,
		'1':	true,
		'\x00':	false,
		-1:		false,
	}

	for d, result := range testData {
		if RuneIsBIT(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsCHAR(t *testing.T) {
	testData := map[rune]bool{
		'a':    true,
		'⌘':    false,
		']':    true,
		'(':    true,
		'*':    true,
		'\\':   true,
		0:      false,
		'\r':   true,
		'\n':   true,
		'9':    true,
		-1:     false,
		'\x8a': false,
	}

	for d, result := range testData {
		if RuneIsCHAR(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsCR(t *testing.T) {
	testData := map[rune]bool{
		'a':    false,
		'⌘':    false,
		']':    false,
		'(':    false,
		'*':    false,
		'\\':   false,
		0:      false,
		'\r':   true,
		'\n':   false,
		'9':    false,
		-1:     false,
		'\x8a': false,
	}

	for d, result := range testData {
		if RuneIsCR(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_StringIsCRLF(t *testing.T) {
	testData := map[string]bool{
		`a`:    false,
		`⌘`:    false,
		`]`:    false,
		`(`:    false,
		`*dfa`:    false,
		"\\":   false,
		`\\`:	false,
		"\r\n":   true,
		"\n\r":   false,
		"\r":   false,
		"\n":   false,
		"\r\x00":    false,
		"\n\x00":    false,
		"":     false,
		`\x8a`: false,
	}

	for d, result := range testData {
		if StringIsCRLF(d) != result {
			t.Errorf("failed test of '%s' (%x)", d, d)
		}
	}
}

func Test_RuneIsCTL(t *testing.T) {
	testData := map[rune]bool{
		'a':    false,
		'⌘':    false,
		']':    false,
		'(':    false,
		'*':    false,
		'\\':   false,
		0:      true,
		'\x7f':   true,
		'\n':   true,
		'\x1a':	true,
		'9':    false,
		-1:     false,
		'\x8a': false,
	}

	for d, result := range testData {
		if RuneIsCTL(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsDIGIT(t *testing.T) {
	testData := map[rune]bool{
		'a':    false,
		'⌘':    false,
		']':    false,
		'(':    false,
		'*':    false,
		'\\':   false,
		0:      false,
		'5':   true,
		'0':   true,
		'9':    true,
		-1:     false,
		'\x8a': false,
	}

	for d, result := range testData {
		if RuneIsDIGIT(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsDQUOTE(t *testing.T) {
	testData := map[rune]bool{
		'a':    false,
		'⌘':    false,
		']':    false,
		'(':    false,
		'*':    false,
		'\\':   false,
		0:      false,
		'"':   true,
		'\'':   false,
		-1:     false,
		'\x8a': false,
	}

	for d, result := range testData {
		if RuneIsDQUOTE(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsHEXDIG(t *testing.T) {
	testData := map[rune]bool{
		'g':    false,
		'⌘':    false,
		']':    false,
		' ':    false,
		'*':    false,
		'\\':   false,
		0:      false,
		'0':	true,
		'a':	true,
		'7':	true,
		'F':	true,
		-1:		false,
		'\x8a':	false,
	}

	for d, result := range testData {
		if RuneIsHEXDIG(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsHTAB(t *testing.T) {
	testData := map[rune]bool{
		'a':    false,
		'⌘':    false,
		']':    false,
		'(':    false,
		'*':    false,
		'\\':   false,
		0:      false,
		'\t':	true,
		'7':	false,
		'\r':	false,
		-1:		false,
		'\x8a':	false,
	}

	for d, result := range testData {
		if RuneIsHTAB(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsLF(t *testing.T) {
	testData := map[rune]bool{
		'a':    false,
		'⌘':    false,
		']':    false,
		'(':    false,
		'*':    false,
		'\\':   false,
		0:      false,
		'\n':	true,
		'\t':	false,
		'7':	false,
		'\r':	false,
		-1:		false,
		'\x8a':	false,
	}

	for d, result := range testData {
		if RuneIsLF(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_StringIsLWSP(t *testing.T) {
	testData := [...]map[string]bool{
		{`a`:    false},
		{`⌘`:    false},
		{`]`:    false},
		{`(`:    false},
		{`*dfa`:    false},
		{"\\":   false},
		{`\\`:	false},
		{"0":      false},
		{"\r\n":   true},
		{"\n\r":   false},
		{"\r":   false},
		{"\n":   false},
		{"\r\x00":    false},
		{"\n\x00":    false},
		{"":     true},
		{`\x8a`: false},
		{"\r\n   ": 	true},
		{"\t  "	: 	true},
		{"\t\r\n \r\n   \t\r\n":	true},
		{"\t\r\n \r\n   \t\r\n ":	true},
		{" \t\r\n \r\n   \t":	true},
	}

	for i, d := range testData {
		for k, result := range d {
			if StringIsLWSP(k) != result {
				t.Errorf("failed test %d of '%s' (%x)", i, k, k)
			}
		}
	}
}

func Test_RuneIsOCTET(t *testing.T) {
	testData := map[rune]bool{
		'a':    true,
		'⌘':    false,
		']':    true,
		'(':    true,
		'*':    true,
		'\\':   true,
		0:      true,
		'\n':	true,
		'\t':	true,
		'7':	true,
		'\r':	true,
		-1:		false,
		'\x8a':	true,
		'\xff':	true,
	}

	for d, result := range testData {
		if RuneIsOCTET(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsSP(t *testing.T) {
	testData := map[rune]bool{
		'a':    false,
		'⌘':    false,
		']':    false,
		'(':    false,
		'*':    false,
		'\\':   false,
		0:      false,
		' ':	true,
		'\t':	false,
		'7':	false,
		'\r':	false,
		-1:		false,
		'\x8a':	false,
	}

	for d, result := range testData {
		if RuneIsSP(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsVCHAR(t *testing.T) {
	testData := map[rune]bool{
		'a':    true,
		'⌘':    false,
		']':    true,
		'(':    true,
		'*':    true,
		'\\':   true,
		0:      false,
		' ':	false,
		'\t':	false,
		'7':	true,
		'\r':	false,
		-1:		false,
		'\x8a':	false,
	}

	for d, result := range testData {
		if RuneIsVCHAR(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_RuneIsWSP(t *testing.T) {
	testData := map[rune]bool{
		'a':    false,
		'⌘':    false,
		']':    false,
		'(':    false,
		'*':    false,
		'\\':   false,
		0:      false,
		' ':	true,
		'\t':	true,
		'7':	false,
		'\r':	false,
		'\n':	false,
		-1:		false,
		'\x8a':	false,
	}

	for d, result := range testData {
		if RuneIsWSP(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}
