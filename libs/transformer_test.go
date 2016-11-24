package libs

import "testing"

func TestStringsToByteAry(t *testing.T) {
	ret := StringsToByteAry("yes ")

	if len(ret) != 4 {
		t.Fail()
	}
}
