package fooutils

import (
	"testing"
)

var(
	id1 = "350122199301067888"
	id2 = "53012219850804624X"
	id3 = "370405199411014860"
	id4 = "52262619780228998X"
	id5 = "411721197701290358"
	id6 = "230714199203187491"
	id7 = "123"
	id8 = "身份证"
	id9 = "idcode-abcdefghijklmnopq"
)

func TestIsValidIDCode(t *testing.T) {
	validCodes := []string{id1, id2, id3}
	for _, code := range validCodes {
		if ret := IsValidIDCode(code); ret != true {
			t.Errorf("Falid IsValidIDCode(%s). Got %v, expected true", code, ret)
		}
	}
}

func TestIsValidIDCode2(t *testing.T) {
	inValidCodes := []string{id4, id5, id6, id7, id8, id9}
	for _, code := range inValidCodes {
		if ret := IsValidIDCode(code); ret != false {
			t.Errorf("Falid IsValidIDCode(%s). Got %v, expected false", code, ret)
		}
	}
}
