package util

import (
	"bytes"
	"testing"
)

// This file tests the scroll functionality

func testScroll(sc *Scroller, expected [][]byte, t *testing.T) {
	// Make sure that the scroll cycle matches the expected cycle for
	// at least 3 loops
	for cycle := 0; cycle < 3; cycle++ {
		for i, exp := range expected {
			got := sc.Bytes()
			t.Logf("%v Cycle %d Step %2d:\n\tgot: %q\n\texp: %q\n", t.Name(), cycle, i, string(got), string(exp))
			if !bytes.Equal(got, exp) {
				t.Errorf("%#v\n%v %q\n\n", sc, sc.textPos, string(sc.buffer[sc.textPos:]))
			}
			sc.Scroll()
		}
	}
}

// Scroll from right-to-left
// |>>>          <<<|
// |>>>        A <<<|
// |>>>       AB <<<|
// |>>>      ABC <<<|
// |>>>     ABCD <<<|
// |>>>    ABCDE <<<|
// |>>>   ABCDEF <<<|
// |>>>  ABCDEFG <<<|
// |>>> ABCDEFGH <<<| # Starts from here if scroll from offscreen is not set
// |>>> BCDEFGHI <<<|
// |>>> CDEFGHIJ <<<| # Ends here if scroll to offscreen is not set
// |>>> DEFGHIJ  <<<|
// |>>> EFGHIJ   <<<|
// |>>> FGHIJ    <<<|
// |>>> GHIJ     <<<|
// |>>> HIJ      <<<|
// |>>> IJ       <<<|
// |>>> J        <<<|
// |>>>          <<<| # Back to step 1
func TestScrollRTLfYtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJ"), []byte(">>> "), []byte(" <<<"),
		ScrollFromOffscreen|ScrollTextOffscreen)

	expected := [][]byte{
		[]byte(">>>          <<<"),
		[]byte(">>>        A <<<"),
		[]byte(">>>       AB <<<"),
		[]byte(">>>      ABC <<<"),
		[]byte(">>>     ABCD <<<"),
		[]byte(">>>    ABCDE <<<"),
		[]byte(">>>   ABCDEF <<<"),
		[]byte(">>>  ABCDEFG <<<"),
		[]byte(">>> ABCDEFGH <<<"),
		[]byte(">>> BCDEFGHI <<<"),
		[]byte(">>> CDEFGHIJ <<<"),
		[]byte(">>> DEFGHIJ  <<<"),
		[]byte(">>> EFGHIJ   <<<"),
		[]byte(">>> FGHIJ    <<<"),
		[]byte(">>> GHIJ     <<<"),
		[]byte(">>> HIJ      <<<"),
		[]byte(">>> IJ       <<<"),
		[]byte(">>> J        <<<"),
		[]byte(">>>          <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollRTLfYtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJ"), []byte(">>> "), []byte(" <<<"),
		ScrollFromOffscreen)

	expected := [][]byte{
		[]byte(">>>          <<<"),
		[]byte(">>>        A <<<"),
		[]byte(">>>       AB <<<"),
		[]byte(">>>      ABC <<<"),
		[]byte(">>>     ABCD <<<"),
		[]byte(">>>    ABCDE <<<"),
		[]byte(">>>   ABCDEF <<<"),
		[]byte(">>>  ABCDEFG <<<"),
		[]byte(">>> ABCDEFGH <<<"),
		[]byte(">>> BCDEFGHI <<<"),
		[]byte(">>> CDEFGHIJ <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollRTLfNtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJ"), []byte(">>> "), []byte(" <<<"),
		ScrollTextOffscreen)

	expected := [][]byte{
		[]byte(">>> ABCDEFGH <<<"),
		[]byte(">>> BCDEFGHI <<<"),
		[]byte(">>> CDEFGHIJ <<<"),
		[]byte(">>> DEFGHIJ  <<<"),
		[]byte(">>> EFGHIJ   <<<"),
		[]byte(">>> FGHIJ    <<<"),
		[]byte(">>> GHIJ     <<<"),
		[]byte(">>> HIJ      <<<"),
		[]byte(">>> IJ       <<<"),
		[]byte(">>> J        <<<"),
		[]byte(">>>          <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollRTLfNtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJ"), []byte(">>> "), []byte(" <<<"), 0)

	expected := [][]byte{
		[]byte(">>> ABCDEFGH <<<"),
		[]byte(">>> BCDEFGHI <<<"),
		[]byte(">>> CDEFGHIJ <<<"),
	}

	testScroll(scroller, expected, t)
}

// Scroll from left-to-right
// |>>>          <<<|
// |>>> J        <<<|
// |>>> IJ       <<<|
// |>>> HIJ      <<<|
// |>>> GHIJ     <<<|
// |>>> FGHIJ    <<<|
// |>>> EFGHIJ   <<<|
// |>>> DEFGHIJ  <<<|
// |>>> CDEFGHIJ <<<| # Starts from here if scroll from offscreen is not set
// |>>> BCDEFGHI <<<|
// |>>> ABCDEFGH <<<| # Ends here if scroll to offscreen is not set
// |>>>  ABCDEFG <<<|
// |>>>   ABCDEF <<<|
// |>>>    ABCDE <<<|
// |>>>     ABCD <<<|
// |>>>      ABC <<<|
// |>>>       AB <<<|
// |>>>        A <<<|
// |>>>          <<< # Back to step 1
func TestScrollLTRfYtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJ"), []byte(">>> "), []byte(" <<<"),
		ScrollFromOffscreen|ScrollTextOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte(">>>          <<<"),
		[]byte(">>> J        <<<"),
		[]byte(">>> IJ       <<<"),
		[]byte(">>> HIJ      <<<"),
		[]byte(">>> GHIJ     <<<"),
		[]byte(">>> FGHIJ    <<<"),
		[]byte(">>> EFGHIJ   <<<"),
		[]byte(">>> DEFGHIJ  <<<"),
		[]byte(">>> CDEFGHIJ <<<"),
		[]byte(">>> BCDEFGHI <<<"),
		[]byte(">>> ABCDEFGH <<<"),
		[]byte(">>>  ABCDEFG <<<"),
		[]byte(">>>   ABCDEF <<<"),
		[]byte(">>>    ABCDE <<<"),
		[]byte(">>>     ABCD <<<"),
		[]byte(">>>      ABC <<<"),
		[]byte(">>>       AB <<<"),
		[]byte(">>>        A <<<"),
		[]byte(">>>          <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollLTRfYtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJ"), []byte(">>> "), []byte(" <<<"),
		ScrollFromOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte(">>>          <<<"),
		[]byte(">>> J        <<<"),
		[]byte(">>> IJ       <<<"),
		[]byte(">>> HIJ      <<<"),
		[]byte(">>> GHIJ     <<<"),
		[]byte(">>> FGHIJ    <<<"),
		[]byte(">>> EFGHIJ   <<<"),
		[]byte(">>> DEFGHIJ  <<<"),
		[]byte(">>> CDEFGHIJ <<<"),
		[]byte(">>> BCDEFGHI <<<"),
		[]byte(">>> ABCDEFGH <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollLTRfNtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJ"), []byte(">>> "), []byte(" <<<"),
		ScrollTextOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte(">>> CDEFGHIJ <<<"),
		[]byte(">>> BCDEFGHI <<<"),
		[]byte(">>> ABCDEFGH <<<"),
		[]byte(">>>  ABCDEFG <<<"),
		[]byte(">>>   ABCDEF <<<"),
		[]byte(">>>    ABCDE <<<"),
		[]byte(">>>     ABCD <<<"),
		[]byte(">>>      ABC <<<"),
		[]byte(">>>       AB <<<"),
		[]byte(">>>        A <<<"),
		[]byte(">>>          <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollLTRfNtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJ"), []byte(">>> "), []byte(" <<<"),
		ScrollLeftToRight)

	expected := [][]byte{
		[]byte(">>> CDEFGHIJ <<<"),
		[]byte(">>> BCDEFGHI <<<"),
		[]byte(">>> ABCDEFGH <<<"),
	}

	testScroll(scroller, expected, t)
}

// Scroll from right-to-left, no suffix
// |>>>             |
// |>>>            A|
// |>>>           AB|
// |>>>          ABC|
// |>>>         ABCD|
// |>>>        ABCDE|
// |>>>       ABCDEF|
// |>>>      ABCDEFG|
// |>>>     ABCDEFGH|
// |>>>    ABCDEFGHI|
// |>>>   ABCDEFGHIJ|
// |>>>  ABCDEFGHIJK|
// |>>> ABCDEFGHIJKL| # Starts from here if scroll from offscreen is not set
// |>>> BCDEFGHIJKLM|
// |>>> CDEFGHIJKLMN|
// |>>> DEFGHIJKLMNO|
// |>>> EFGHIJKLMNOP| # Ends here if scroll to offscreen is not set
// |>>> FGHIJKLMNOP |
// |>>> GHIJKLMNOP  |
// |>>> HIJKLMNOP   |
// |>>> IJKLMNOP    |
// |>>> JKLMNOP     |
// |>>> KLMNOP      |
// |>>> LMNOP       |
// |>>> MNOP        |
// |>>> NOP         |
// |>>> OP          |
// |>>> P           |
// |>>>             | #     Back to step 1

func TestScrollNoSuffixRTLfYtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), []byte(">>> "), nil,
		ScrollFromOffscreen|ScrollTextOffscreen)

	expected := [][]byte{
		[]byte(">>>             "),
		[]byte(">>>            A"),
		[]byte(">>>           AB"),
		[]byte(">>>          ABC"),
		[]byte(">>>         ABCD"),
		[]byte(">>>        ABCDE"),
		[]byte(">>>       ABCDEF"),
		[]byte(">>>      ABCDEFG"),
		[]byte(">>>     ABCDEFGH"),
		[]byte(">>>    ABCDEFGHI"),
		[]byte(">>>   ABCDEFGHIJ"),
		[]byte(">>>  ABCDEFGHIJK"),
		[]byte(">>> ABCDEFGHIJKL"),
		[]byte(">>> BCDEFGHIJKLM"),
		[]byte(">>> CDEFGHIJKLMN"),
		[]byte(">>> DEFGHIJKLMNO"),
		[]byte(">>> EFGHIJKLMNOP"),
		[]byte(">>> FGHIJKLMNOP "),
		[]byte(">>> GHIJKLMNOP  "),
		[]byte(">>> HIJKLMNOP   "),
		[]byte(">>> IJKLMNOP    "),
		[]byte(">>> JKLMNOP     "),
		[]byte(">>> KLMNOP      "),
		[]byte(">>> LMNOP       "),
		[]byte(">>> MNOP        "),
		[]byte(">>> NOP         "),
		[]byte(">>> OP          "),
		[]byte(">>> P           "),
		[]byte(">>>             "),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoSuffixRTLfYtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), []byte(">>> "), nil,
		ScrollFromOffscreen)

	expected := [][]byte{
		[]byte(">>>             "),
		[]byte(">>>            A"),
		[]byte(">>>           AB"),
		[]byte(">>>          ABC"),
		[]byte(">>>         ABCD"),
		[]byte(">>>        ABCDE"),
		[]byte(">>>       ABCDEF"),
		[]byte(">>>      ABCDEFG"),
		[]byte(">>>     ABCDEFGH"),
		[]byte(">>>    ABCDEFGHI"),
		[]byte(">>>   ABCDEFGHIJ"),
		[]byte(">>>  ABCDEFGHIJK"),
		[]byte(">>> ABCDEFGHIJKL"),
		[]byte(">>> BCDEFGHIJKLM"),
		[]byte(">>> CDEFGHIJKLMN"),
		[]byte(">>> DEFGHIJKLMNO"),
		[]byte(">>> EFGHIJKLMNOP"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoSuffixRTLfNtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), []byte(">>> "), nil,
		ScrollTextOffscreen)

	expected := [][]byte{
		[]byte(">>> ABCDEFGHIJKL"),
		[]byte(">>> BCDEFGHIJKLM"),
		[]byte(">>> CDEFGHIJKLMN"),
		[]byte(">>> DEFGHIJKLMNO"),
		[]byte(">>> EFGHIJKLMNOP"),
		[]byte(">>> FGHIJKLMNOP "),
		[]byte(">>> GHIJKLMNOP  "),
		[]byte(">>> HIJKLMNOP   "),
		[]byte(">>> IJKLMNOP    "),
		[]byte(">>> JKLMNOP     "),
		[]byte(">>> KLMNOP      "),
		[]byte(">>> LMNOP       "),
		[]byte(">>> MNOP        "),
		[]byte(">>> NOP         "),
		[]byte(">>> OP          "),
		[]byte(">>> P           "),
		[]byte(">>>             "),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoSuffixRTLfNtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), []byte(">>> "), nil,
		0)

	expected := [][]byte{
		[]byte(">>> ABCDEFGHIJKL"),
		[]byte(">>> BCDEFGHIJKLM"),
		[]byte(">>> CDEFGHIJKLMN"),
		[]byte(">>> DEFGHIJKLMNO"),
		[]byte(">>> EFGHIJKLMNOP"),
	}

	testScroll(scroller, expected, t)
}

// Scroll from left-to-right, no suffix
// |>>>             |
// |>>> P           |
// |>>> OP          |
// |>>> NOP         |
// |>>> MNOP        |
// |>>> LMNOP       |
// |>>> KLMNOP      |
// |>>> JKLMNOP     |
// |>>> IJKLMNOP    |
// |>>> HIJKLMNOP   |
// |>>> GHIJKLMNOP  |
// |>>> FGHIJKLMNOP |
// |>>> EFGHIJKLMNOP| # Starts from here if scroll from offscreen is not set
// |>>> DEFGHIJKLMNO|
// |>>> CDEFGHIJKLMN|
// |>>> BCDEFGHIJKLM|
// |>>> ABCDEFGHIJKL| # Ends here if scroll to offscreen is not set
// |>>>  ABCDEFGHIJK|
// |>>>   ABCDEFGHIJ|
// |>>>    ABCDEFGHI|
// |>>>     ABCDEFGH|
// |>>>      ABCDEFG|
// |>>>       ABCDEF|
// |>>>        ABCDE|
// |>>>         ABCD|
// |>>>          ABC|
// |>>>           AB|
// |>>>            A|
// |>>>             | #     Back to step 1

func TestScrollNoSuffixLTRfYtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), []byte(">>> "), nil,
		ScrollFromOffscreen|ScrollTextOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte(">>>             "),
		[]byte(">>> P           "),
		[]byte(">>> OP          "),
		[]byte(">>> NOP         "),
		[]byte(">>> MNOP        "),
		[]byte(">>> LMNOP       "),
		[]byte(">>> KLMNOP      "),
		[]byte(">>> JKLMNOP     "),
		[]byte(">>> IJKLMNOP    "),
		[]byte(">>> HIJKLMNOP   "),
		[]byte(">>> GHIJKLMNOP  "),
		[]byte(">>> FGHIJKLMNOP "),
		[]byte(">>> EFGHIJKLMNOP"),
		[]byte(">>> DEFGHIJKLMNO"),
		[]byte(">>> CDEFGHIJKLMN"),
		[]byte(">>> BCDEFGHIJKLM"),
		[]byte(">>> ABCDEFGHIJKL"),
		[]byte(">>>  ABCDEFGHIJK"),
		[]byte(">>>   ABCDEFGHIJ"),
		[]byte(">>>    ABCDEFGHI"),
		[]byte(">>>     ABCDEFGH"),
		[]byte(">>>      ABCDEFG"),
		[]byte(">>>       ABCDEF"),
		[]byte(">>>        ABCDE"),
		[]byte(">>>         ABCD"),
		[]byte(">>>          ABC"),
		[]byte(">>>           AB"),
		[]byte(">>>            A"),
		[]byte(">>>             "),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoSuffixLTRfYtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), []byte(">>> "), nil,
		ScrollFromOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte(">>>             "),
		[]byte(">>> P           "),
		[]byte(">>> OP          "),
		[]byte(">>> NOP         "),
		[]byte(">>> MNOP        "),
		[]byte(">>> LMNOP       "),
		[]byte(">>> KLMNOP      "),
		[]byte(">>> JKLMNOP     "),
		[]byte(">>> IJKLMNOP    "),
		[]byte(">>> HIJKLMNOP   "),
		[]byte(">>> GHIJKLMNOP  "),
		[]byte(">>> FGHIJKLMNOP "),
		[]byte(">>> EFGHIJKLMNOP"),
		[]byte(">>> DEFGHIJKLMNO"),
		[]byte(">>> CDEFGHIJKLMN"),
		[]byte(">>> BCDEFGHIJKLM"),
		[]byte(">>> ABCDEFGHIJKL"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoSuffixLTRfNtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), []byte(">>> "), nil,
		ScrollTextOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte(">>> EFGHIJKLMNOP"),
		[]byte(">>> DEFGHIJKLMNO"),
		[]byte(">>> CDEFGHIJKLMN"),
		[]byte(">>> BCDEFGHIJKLM"),
		[]byte(">>> ABCDEFGHIJKL"),
		[]byte(">>>  ABCDEFGHIJK"),
		[]byte(">>>   ABCDEFGHIJ"),
		[]byte(">>>    ABCDEFGHI"),
		[]byte(">>>     ABCDEFGH"),
		[]byte(">>>      ABCDEFG"),
		[]byte(">>>       ABCDEF"),
		[]byte(">>>        ABCDE"),
		[]byte(">>>         ABCD"),
		[]byte(">>>          ABC"),
		[]byte(">>>           AB"),
		[]byte(">>>            A"),
		[]byte(">>>             "),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoSuffixLTRfNtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), []byte(">>> "), nil,
		ScrollLeftToRight)

	expected := [][]byte{
		[]byte(">>> EFGHIJKLMNOP"),
		[]byte(">>> DEFGHIJKLMNO"),
		[]byte(">>> CDEFGHIJKLMN"),
		[]byte(">>> BCDEFGHIJKLM"),
		[]byte(">>> ABCDEFGHIJKL"),
	}

	testScroll(scroller, expected, t)
}

// Scroll from right-to-left, no prefix
// |             <<<|
// |           A <<<|
// |          AB <<<|
// |         ABC <<<|
// |        ABCD <<<|
// |       ABCDE <<<|
// |      ABCDEF <<<|
// |     ABCDEFG <<<|
// |    ABCDEFGH <<<|
// |   ABCDEFGHI <<<|
// |  ABCDEFGHIJ <<<|
// | ABCDEFGHIJK <<<|
// |ABCDEFGHIJKL <<<| # Starts from here if scroll from offscreen is not set
// |BCDEFGHIJKLM <<<|
// |CDEFGHIJKLMN <<<|
// |DEFGHIJKLMNO <<<|
// |EFGHIJKLMNOP <<<| # Ends here if scroll to offscreen is not set
// |FGHIJKLMNOP  <<<|
// |GHIJKLMNOP   <<<|
// |HIJKLMNOP    <<<|
// |IJKLMNOP     <<<|
// |JKLMNOP      <<<|
// |KLMNOP       <<<|
// |LMNOP        <<<|
// |MNOP         <<<|
// |NOP          <<<|
// |OP           <<<|
// |P            <<<|
// |             <<<| #     Back to step 1

func TestScrollNoPrefixRTLfYtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), nil, []byte(" <<<"),
		ScrollFromOffscreen|ScrollTextOffscreen)

	expected := [][]byte{
		[]byte("             <<<"),
		[]byte("           A <<<"),
		[]byte("          AB <<<"),
		[]byte("         ABC <<<"),
		[]byte("        ABCD <<<"),
		[]byte("       ABCDE <<<"),
		[]byte("      ABCDEF <<<"),
		[]byte("     ABCDEFG <<<"),
		[]byte("    ABCDEFGH <<<"),
		[]byte("   ABCDEFGHI <<<"),
		[]byte("  ABCDEFGHIJ <<<"),
		[]byte(" ABCDEFGHIJK <<<"),
		[]byte("ABCDEFGHIJKL <<<"),
		[]byte("BCDEFGHIJKLM <<<"),
		[]byte("CDEFGHIJKLMN <<<"),
		[]byte("DEFGHIJKLMNO <<<"),
		[]byte("EFGHIJKLMNOP <<<"),
		[]byte("FGHIJKLMNOP  <<<"),
		[]byte("GHIJKLMNOP   <<<"),
		[]byte("HIJKLMNOP    <<<"),
		[]byte("IJKLMNOP     <<<"),
		[]byte("JKLMNOP      <<<"),
		[]byte("KLMNOP       <<<"),
		[]byte("LMNOP        <<<"),
		[]byte("MNOP         <<<"),
		[]byte("NOP          <<<"),
		[]byte("OP           <<<"),
		[]byte("P            <<<"),
		[]byte("             <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixRTLfYtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), nil, []byte(" <<<"),
		ScrollFromOffscreen)

	expected := [][]byte{
		[]byte("             <<<"),
		[]byte("           A <<<"),
		[]byte("          AB <<<"),
		[]byte("         ABC <<<"),
		[]byte("        ABCD <<<"),
		[]byte("       ABCDE <<<"),
		[]byte("      ABCDEF <<<"),
		[]byte("     ABCDEFG <<<"),
		[]byte("    ABCDEFGH <<<"),
		[]byte("   ABCDEFGHI <<<"),
		[]byte("  ABCDEFGHIJ <<<"),
		[]byte(" ABCDEFGHIJK <<<"),
		[]byte("ABCDEFGHIJKL <<<"),
		[]byte("BCDEFGHIJKLM <<<"),
		[]byte("CDEFGHIJKLMN <<<"),
		[]byte("DEFGHIJKLMNO <<<"),
		[]byte("EFGHIJKLMNOP <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixRTLfNtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), nil, []byte(" <<<"),
		ScrollTextOffscreen)

	expected := [][]byte{
		[]byte("ABCDEFGHIJKL <<<"),
		[]byte("BCDEFGHIJKLM <<<"),
		[]byte("CDEFGHIJKLMN <<<"),
		[]byte("DEFGHIJKLMNO <<<"),
		[]byte("EFGHIJKLMNOP <<<"),
		[]byte("FGHIJKLMNOP  <<<"),
		[]byte("GHIJKLMNOP   <<<"),
		[]byte("HIJKLMNOP    <<<"),
		[]byte("IJKLMNOP     <<<"),
		[]byte("JKLMNOP      <<<"),
		[]byte("KLMNOP       <<<"),
		[]byte("LMNOP        <<<"),
		[]byte("MNOP         <<<"),
		[]byte("NOP          <<<"),
		[]byte("OP           <<<"),
		[]byte("P            <<<"),
		[]byte("             <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixRTLfNtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), nil, []byte(" <<<"), 0)

	expected := [][]byte{
		[]byte("ABCDEFGHIJKL <<<"),
		[]byte("BCDEFGHIJKLM <<<"),
		[]byte("CDEFGHIJKLMN <<<"),
		[]byte("DEFGHIJKLMNO <<<"),
		[]byte("EFGHIJKLMNOP <<<"),
	}

	testScroll(scroller, expected, t)
}

// Scroll from left-to-right, no prefix
// |             <<<|
// |P            <<<|
// |OP           <<<|
// |NOP          <<<|
// |MNOP         <<<|
// |LMNOP        <<<|
// |KLMNOP       <<<|
// |JKLMNOP      <<<|
// |IJKLMNOP     <<<|
// |HIJKLMNOP    <<<|
// |GHIJKLMNOP   <<<|
// |FGHIJKLMNOP  <<<|
// |EFGHIJKLMNOP <<<| # Starts from here if scroll from offscreen is not set
// |DEFGHIJKLMNO <<<|
// |CDEFGHIJKLMN <<<|
// |BCDEFGHIJKLM <<<|
// |ABCDEFGHIJKL <<<| # Ends here if scroll to offscreen is not set
// | ABCDEFGHIJK <<<|
// |  ABCDEFGHIJ <<<|
// |   ABCDEFGHI <<<|
// |    ABCDEFGH <<<|
// |     ABCDEFG <<<|
// |      ABCDEF <<<|
// |       ABCDE <<<|
// |        ABCD <<<|
// |         ABC <<<|
// |          AB <<<|
// |           A <<<|
// |             <<<| #     Back to step 1

func TestScrollNoPrefixLTRfYtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), nil, []byte(" <<<"),
		ScrollFromOffscreen|ScrollTextOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte("             <<<"),
		[]byte("P            <<<"),
		[]byte("OP           <<<"),
		[]byte("NOP          <<<"),
		[]byte("MNOP         <<<"),
		[]byte("LMNOP        <<<"),
		[]byte("KLMNOP       <<<"),
		[]byte("JKLMNOP      <<<"),
		[]byte("IJKLMNOP     <<<"),
		[]byte("HIJKLMNOP    <<<"),
		[]byte("GHIJKLMNOP   <<<"),
		[]byte("FGHIJKLMNOP  <<<"),
		[]byte("EFGHIJKLMNOP <<<"),
		[]byte("DEFGHIJKLMNO <<<"),
		[]byte("CDEFGHIJKLMN <<<"),
		[]byte("BCDEFGHIJKLM <<<"),
		[]byte("ABCDEFGHIJKL <<<"),
		[]byte(" ABCDEFGHIJK <<<"),
		[]byte("  ABCDEFGHIJ <<<"),
		[]byte("   ABCDEFGHI <<<"),
		[]byte("    ABCDEFGH <<<"),
		[]byte("     ABCDEFG <<<"),
		[]byte("      ABCDEF <<<"),
		[]byte("       ABCDE <<<"),
		[]byte("        ABCD <<<"),
		[]byte("         ABC <<<"),
		[]byte("          AB <<<"),
		[]byte("           A <<<"),
		[]byte("             <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixLTRfYtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), nil, []byte(" <<<"),
		ScrollFromOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte("             <<<"),
		[]byte("P            <<<"),
		[]byte("OP           <<<"),
		[]byte("NOP          <<<"),
		[]byte("MNOP         <<<"),
		[]byte("LMNOP        <<<"),
		[]byte("KLMNOP       <<<"),
		[]byte("JKLMNOP      <<<"),
		[]byte("IJKLMNOP     <<<"),
		[]byte("HIJKLMNOP    <<<"),
		[]byte("GHIJKLMNOP   <<<"),
		[]byte("FGHIJKLMNOP  <<<"),
		[]byte("EFGHIJKLMNOP <<<"),
		[]byte("DEFGHIJKLMNO <<<"),
		[]byte("CDEFGHIJKLMN <<<"),
		[]byte("BCDEFGHIJKLM <<<"),
		[]byte("ABCDEFGHIJKL <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixLTRfNtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), nil, []byte(" <<<"),
		ScrollTextOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte("EFGHIJKLMNOP <<<"),
		[]byte("DEFGHIJKLMNO <<<"),
		[]byte("CDEFGHIJKLMN <<<"),
		[]byte("BCDEFGHIJKLM <<<"),
		[]byte("ABCDEFGHIJKL <<<"),
		[]byte(" ABCDEFGHIJK <<<"),
		[]byte("  ABCDEFGHIJ <<<"),
		[]byte("   ABCDEFGHI <<<"),
		[]byte("    ABCDEFGH <<<"),
		[]byte("     ABCDEFG <<<"),
		[]byte("      ABCDEF <<<"),
		[]byte("       ABCDE <<<"),
		[]byte("        ABCD <<<"),
		[]byte("         ABC <<<"),
		[]byte("          AB <<<"),
		[]byte("           A <<<"),
		[]byte("             <<<"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixLTRfNtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOP"), nil, []byte(" <<<"),
		ScrollLeftToRight)

	expected := [][]byte{
		[]byte("EFGHIJKLMNOP <<<"),
		[]byte("DEFGHIJKLMNO <<<"),
		[]byte("CDEFGHIJKLMN <<<"),
		[]byte("BCDEFGHIJKLM <<<"),
		[]byte("ABCDEFGHIJKL <<<"),
	}

	testScroll(scroller, expected, t)
}

// Scroll from right-to-left, no prefix or suffix
// |                |
// |               A|
// |              AB|
// |             ABC|
// |            ABCD|
// |           ABCDE|
// |          ABCDEF|
// |         ABCDEFG|
// |        ABCDEFGH|
// |       ABCDEFGHI|
// |      ABCDEFGHIJ|
// |     ABCDEFGHIJK|
// |    ABCDEFGHIJKL|
// |   ABCDEFGHIJKLM|
// |  ABCDEFGHIJKLMN|
// | ABCDEFGHIJKLMNO|
// |ABCDEFGHIJKLMNOP| # Starts from here if scroll from offscreen is not set
// |BCDEFGHIJKLMNOPQ| # Ends here if scroll to offscreen is not set
// |CDEFGHIJKLMNOPQ |
// |DEFGHIJKLMNOPQ  |
// |EFGHIJKLMNOPQ   |
// |FGHIJKLMNOPQ    |
// |GHIJKLMNOPQ     |
// |HIJKLMNOPQ      |
// |IJKLMNOPQ       |
// |JKLMNOPQ        |
// |KLMNOPQ         |
// |LMNOPQ          |
// |MNOPQ           |
// |NOPQ            |
// |OPQ             |
// |PQ              |
// |Q               |
// |                | #     Back to step 1

func TestScrollNoPrefixOrSuffixRTLfYtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOPQ"), nil, nil,
		ScrollFromOffscreen|ScrollTextOffscreen)

	expected := [][]byte{
		[]byte("                "),
		[]byte("               A"),
		[]byte("              AB"),
		[]byte("             ABC"),
		[]byte("            ABCD"),
		[]byte("           ABCDE"),
		[]byte("          ABCDEF"),
		[]byte("         ABCDEFG"),
		[]byte("        ABCDEFGH"),
		[]byte("       ABCDEFGHI"),
		[]byte("      ABCDEFGHIJ"),
		[]byte("     ABCDEFGHIJK"),
		[]byte("    ABCDEFGHIJKL"),
		[]byte("   ABCDEFGHIJKLM"),
		[]byte("  ABCDEFGHIJKLMN"),
		[]byte(" ABCDEFGHIJKLMNO"),
		[]byte("ABCDEFGHIJKLMNOP"),
		[]byte("BCDEFGHIJKLMNOPQ"),
		[]byte("CDEFGHIJKLMNOPQ "),
		[]byte("DEFGHIJKLMNOPQ  "),
		[]byte("EFGHIJKLMNOPQ   "),
		[]byte("FGHIJKLMNOPQ    "),
		[]byte("GHIJKLMNOPQ     "),
		[]byte("HIJKLMNOPQ      "),
		[]byte("IJKLMNOPQ       "),
		[]byte("JKLMNOPQ        "),
		[]byte("KLMNOPQ         "),
		[]byte("LMNOPQ          "),
		[]byte("MNOPQ           "),
		[]byte("NOPQ            "),
		[]byte("OPQ             "),
		[]byte("PQ              "),
		[]byte("Q               "),
		[]byte("                "),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixOrSuffixRTLfYtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOPQ"), nil, nil,
		ScrollFromOffscreen)

	expected := [][]byte{
		[]byte("                "),
		[]byte("               A"),
		[]byte("              AB"),
		[]byte("             ABC"),
		[]byte("            ABCD"),
		[]byte("           ABCDE"),
		[]byte("          ABCDEF"),
		[]byte("         ABCDEFG"),
		[]byte("        ABCDEFGH"),
		[]byte("       ABCDEFGHI"),
		[]byte("      ABCDEFGHIJ"),
		[]byte("     ABCDEFGHIJK"),
		[]byte("    ABCDEFGHIJKL"),
		[]byte("   ABCDEFGHIJKLM"),
		[]byte("  ABCDEFGHIJKLMN"),
		[]byte(" ABCDEFGHIJKLMNO"),
		[]byte("ABCDEFGHIJKLMNOP"),
		[]byte("BCDEFGHIJKLMNOPQ"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixOrSuffixRTLfNtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOPQ"), nil, nil,
		ScrollTextOffscreen)

	expected := [][]byte{
		[]byte("ABCDEFGHIJKLMNOP"),
		[]byte("BCDEFGHIJKLMNOPQ"),
		[]byte("CDEFGHIJKLMNOPQ "),
		[]byte("DEFGHIJKLMNOPQ  "),
		[]byte("EFGHIJKLMNOPQ   "),
		[]byte("FGHIJKLMNOPQ    "),
		[]byte("GHIJKLMNOPQ     "),
		[]byte("HIJKLMNOPQ      "),
		[]byte("IJKLMNOPQ       "),
		[]byte("JKLMNOPQ        "),
		[]byte("KLMNOPQ         "),
		[]byte("LMNOPQ          "),
		[]byte("MNOPQ           "),
		[]byte("NOPQ            "),
		[]byte("OPQ             "),
		[]byte("PQ              "),
		[]byte("Q               "),
		[]byte("                "),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixOrSuffixRTLfNtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOPQ"), nil, nil, 0)

	expected := [][]byte{
		[]byte("ABCDEFGHIJKLMNOP"),
		[]byte("BCDEFGHIJKLMNOPQ"),
	}

	testScroll(scroller, expected, t)
}

// Scroll from left-to-right, no prefix or suffix
// |                |
// |Q               |
// |PQ              |
// |OPQ             |
// |NOPQ            |
// |MNOPQ           |
// |LMNOPQ          |
// |KLMNOPQ         |
// |JKLMNOPQ        |
// |IJKLMNOPQ       |
// |HIJKLMNOPQ      |
// |GHIJKLMNOPQ     |
// |FGHIJKLMNOPQ    |
// |EFGHIJKLMNOPQ   |
// |DEFGHIJKLMNOPQ  |
// |CDEFGHIJKLMNOPQ |
// |BCDEFGHIJKLMNOPQ| # Starts from here if scroll from offscreen is not set
// |ABCDEFGHIJKLMNOP| # Ends here if scroll to offscreen is not set
// | ABCDEFGHIJKLMNO|
// |  ABCDEFGHIJKLMN|
// |   ABCDEFGHIJKLM|
// |    ABCDEFGHIJKL|
// |     ABCDEFGHIJK|
// |      ABCDEFGHIJ|
// |       ABCDEFGHI|
// |        ABCDEFGH|
// |         ABCDEFG|
// |          ABCDEF|
// |           ABCDE|
// |            ABCD|
// |             ABC|
// |              AB|
// |               A|
// |                | #     Back to step 1

func TestScrollNoPrefixOrSuffixLTRfYtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOPQ"), nil, nil,
		ScrollFromOffscreen|ScrollTextOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte("                "),
		[]byte("Q               "),
		[]byte("PQ              "),
		[]byte("OPQ             "),
		[]byte("NOPQ            "),
		[]byte("MNOPQ           "),
		[]byte("LMNOPQ          "),
		[]byte("KLMNOPQ         "),
		[]byte("JKLMNOPQ        "),
		[]byte("IJKLMNOPQ       "),
		[]byte("HIJKLMNOPQ      "),
		[]byte("GHIJKLMNOPQ     "),
		[]byte("FGHIJKLMNOPQ    "),
		[]byte("EFGHIJKLMNOPQ   "),
		[]byte("DEFGHIJKLMNOPQ  "),
		[]byte("CDEFGHIJKLMNOPQ "),
		[]byte("BCDEFGHIJKLMNOPQ"),
		[]byte("ABCDEFGHIJKLMNOP"),
		[]byte(" ABCDEFGHIJKLMNO"),
		[]byte("  ABCDEFGHIJKLMN"),
		[]byte("   ABCDEFGHIJKLM"),
		[]byte("    ABCDEFGHIJKL"),
		[]byte("     ABCDEFGHIJK"),
		[]byte("      ABCDEFGHIJ"),
		[]byte("       ABCDEFGHI"),
		[]byte("        ABCDEFGH"),
		[]byte("         ABCDEFG"),
		[]byte("          ABCDEF"),
		[]byte("           ABCDE"),
		[]byte("            ABCD"),
		[]byte("             ABC"),
		[]byte("              AB"),
		[]byte("               A"),
		[]byte("                "),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixOrSuffixLTRfYtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOPQ"), nil, nil,
		ScrollFromOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte("                "),
		[]byte("Q               "),
		[]byte("PQ              "),
		[]byte("OPQ             "),
		[]byte("NOPQ            "),
		[]byte("MNOPQ           "),
		[]byte("LMNOPQ          "),
		[]byte("KLMNOPQ         "),
		[]byte("JKLMNOPQ        "),
		[]byte("IJKLMNOPQ       "),
		[]byte("HIJKLMNOPQ      "),
		[]byte("GHIJKLMNOPQ     "),
		[]byte("FGHIJKLMNOPQ    "),
		[]byte("EFGHIJKLMNOPQ   "),
		[]byte("DEFGHIJKLMNOPQ  "),
		[]byte("CDEFGHIJKLMNOPQ "),
		[]byte("BCDEFGHIJKLMNOPQ"),
		[]byte("ABCDEFGHIJKLMNOP"),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixOrSuffixLTRfNtY(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOPQ"), nil, nil,
		ScrollTextOffscreen|ScrollLeftToRight)

	expected := [][]byte{
		[]byte("BCDEFGHIJKLMNOPQ"),
		[]byte("ABCDEFGHIJKLMNOP"),
		[]byte(" ABCDEFGHIJKLMNO"),
		[]byte("  ABCDEFGHIJKLMN"),
		[]byte("   ABCDEFGHIJKLM"),
		[]byte("    ABCDEFGHIJKL"),
		[]byte("     ABCDEFGHIJK"),
		[]byte("      ABCDEFGHIJ"),
		[]byte("       ABCDEFGHI"),
		[]byte("        ABCDEFGH"),
		[]byte("         ABCDEFG"),
		[]byte("          ABCDEF"),
		[]byte("           ABCDE"),
		[]byte("            ABCD"),
		[]byte("             ABC"),
		[]byte("              AB"),
		[]byte("               A"),
		[]byte("                "),
	}

	testScroll(scroller, expected, t)
}

func TestScrollNoPrefixOrSuffixLTRfNtN(t *testing.T) {
	scroller, _ := NewScroller([]byte("ABCDEFGHIJKLMNOPQ"), nil, nil,
		ScrollLeftToRight)

	expected := [][]byte{
		[]byte("BCDEFGHIJKLMNOPQ"),
		[]byte("ABCDEFGHIJKLMNOP"),
	}

	testScroll(scroller, expected, t)
}

func testShort(prefix, suffix, exp []byte, t *testing.T) {
	scroller, _ := NewScroller([]byte("foobar"), prefix, suffix, 0)

	expected := [][]byte{exp}

	testScroll(scroller, expected, t)
}

func TestScrollerShortText(t *testing.T) {
	pre := []byte(">>> ")
	suf := []byte(" <<<")
	testShort(pre, suf, []byte(">>> foobar   <<<"), t)
	testShort(pre, nil, []byte(">>> foobar      "), t)
	testShort(nil, suf, []byte("foobar       <<<"), t)
	testShort(nil, nil, []byte("foobar          "), t)
}
