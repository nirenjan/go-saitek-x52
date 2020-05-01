package util

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestTranslation(t *testing.T) {
	// Walk through the entire map, creating strings with single character
	// that exists in the map, and verify that the convert maps to the correct
	// character
	for k, v := range charmap {
		s := string(k)
		b := ConvertStringToX52Charmap(s, false, 0)
		if len(b) != 1 || b[0] != v {
			t.Errorf("Mismatch in conversion, expected [%x], got %#x", v, b)
		}
	}
}

func TestLongString(t *testing.T) {
	// Create a number of random strings, with rune values in the range
	// 0x0000-0xFFFF, and verify that the byte output is as expected
	for i := int32(0); i < rand.Int31()&int32(0xFFFF); i++ {
		// Create a random string upto 255 bytes long
		r := make([]rune, 0, rand.Int31()&int32(0xFF))
		withRepl := make([]byte, 0, cap(r))
		withoutRepl := make([]byte, 0, cap(r))
		for j := 0; j < cap(r); j++ {
			// Get a random rune less than 0x10000
			c := rand.Int31() & 0xFFFF
			r = append(r, c)
			b, ok := charmap[c]
			if ok {
				withRepl = append(withRepl, b)
				withoutRepl = append(withoutRepl, b)
			} else {
				withRepl = append(withRepl, ReplaceMissing)
			}
		}

		s := string(r)
		t.Logf("TestLongString %04x - checking %v, expecting %#x, %#x\n", i, s, withRepl, withoutRepl)
		gotWithRepl := ConvertStringToX52Charmap(s, true, ReplaceMissing)
		gotWithoutRepl := ConvertStringToX52Charmap(s, false, 0)

		if !bytes.Equal(gotWithRepl, withRepl) {
			t.Errorf("String %v, expected %#x, got %#x", s, withRepl, gotWithRepl)
		}

		if !bytes.Equal(gotWithoutRepl, withoutRepl) {
			t.Errorf("String %v, expected %#x, got %#x", s, withoutRepl, gotWithoutRepl)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertStringToX52Charmap("\uFF71\uFF72\uFF73\uFF74\uFF75", false, 0)
	}
}
