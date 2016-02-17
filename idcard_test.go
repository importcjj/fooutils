package fooutils

import (
	"testing"
)

const(
	id1 = "350122199301067888"
	id2 = "53012219850804624X"
	id3 = "370405199411014860"
	id4 = "52262619780228998X"
	id5 = "411721197701290358"
	id6 = "230714199203187491"
)

func TestVerifyId(t *testing.T) {
	trueIds := []string{id1, id2, id3}
	falseIds := []string{id4, id5, id6}
	for _, v := range trueIds {
		if !IsValid(v) {
			t.Errorf("%s is true!", v)
		}
	}
	for _, v := range falseIds {
		if IsValid(v) {
			t.Errorf("%s is fake!", v)
		}
	}
}
