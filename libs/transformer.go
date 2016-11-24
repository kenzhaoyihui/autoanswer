package libs

import k "github.com/micmonay/keybd_event"

// SpecialChars is
const SpecialChars = "~!@#$%^&*()_+:\"<>?{}"

// Mapper translate single char to key number
var Mapper = map[rune]int{
	'a':  k.VK_A,
	'b':  k.VK_B,
	'c':  k.VK_C,
	'd':  k.VK_D,
	'e':  k.VK_E,
	'f':  k.VK_F,
	'g':  k.VK_G,
	'h':  k.VK_H,
	'i':  k.VK_I,
	'j':  k.VK_J,
	'k':  k.VK_K,
	'l':  k.VK_L,
	'm':  k.VK_M,
	'n':  k.VK_N,
	'o':  k.VK_O,
	'p':  k.VK_P,
	'q':  k.VK_Q,
	'r':  k.VK_R,
	's':  k.VK_S,
	't':  k.VK_T,
	'u':  k.VK_U,
	'v':  k.VK_V,
	'w':  k.VK_W,
	'x':  k.VK_X,
	'y':  k.VK_Y,
	'z':  k.VK_Z,
	'-':  k.VK_MINUS,
	'.':  k.VK_DOT,
	' ':  k.VK_SPACE,
	'/':  k.VK_SLASH,
	'\n': k.VK_ENTER,
	'=': k.VK_

	// special chars below
	'|': k.VK_BACKSLASH * -100,
	'*': k.VK_8 * -100,
	'>': k.VK_DOT * -100,
	'<': k.VK_COMMA * -100,
}

// StringsToByteAry translate string to byte arrary
func StringsToByteAry(s string) []int {
	keySeq := make([]int, 0, len(s))

	for _, c := range s {
		keySeq = append(keySeq, Mapper[c])
	}
	return keySeq
}
