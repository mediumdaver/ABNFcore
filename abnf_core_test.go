package ABNFcore

import (
	"testing"
)

func Test_rune_is_alpha(t *testing.T) {
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
		if rune_is_alpha(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_rune_is_digit(t *testing.T) {
	testData := map[rune]bool{
		'a':	false,
		'\x8a':	false,
		'%':	false,
		'*':	false,
		'G':	false,
		'7':	true,
		'\x00':	false,
		-1:		false,
	}

	for d, result := range testData {
		if rune_is_digit(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}

func Test_rune_is_char(t *testing.T) {
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
		if rune_is_char(d) != result {
			t.Errorf("failed test of '%c' (%x)", d, d)
		}
	}
}