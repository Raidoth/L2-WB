package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {

	data := map[string]string{
		"a4bc2d5e":  "aaaabccddddde",
		"abcd":      "abcd",
		"45":        "",
		"":          "",
		"qwe\\4\\5": "qwe45",
		"qwe\\45":   "qwe44444",
		"qwe\\\\5":  "qwe\\\\\\\\\\",
	}
	for i, v := range data {
		if s, err := UnpackString(i); s != v {
			t.Error("Func no correct")
		} else {
			t.Log("Test ok value:", s, "err", err)
		}
	}

}
