package x52

import (
	"bytes"
	"testing"
)

func TestBlinkOn(t *testing.T) {
	ctx := NewContext()
	defer ctx.Close()

	ctx.SetBlink(true)

	if ctx.ledMask != (1 << 24) {
		t.Errorf("Unexpected update mask value, expected %v, got %v", 1<<24, ctx.ledMask)
	}
	if ctx.updateMask != (1 << 24) {
		t.Errorf("Unexpected update mask value, expected %v, got %v", 1<<24, ctx.updateMask)
	}
}

func TestBlinkOff(t *testing.T) {
	ctx := NewContext()
	defer ctx.Close()

	ctx.SetBlink(false)

	if ctx.ledMask != 0 {
		t.Errorf("Unexpected update mask value, expected %v, got %v", 0, ctx.ledMask)
	}
	if ctx.updateMask != (1 << 24) {
		t.Errorf("Unexpected update mask value, expected %v, got %v", 1<<24, ctx.updateMask)
	}
}

func TestShiftOn(t *testing.T) {
	ctx := NewContext()
	defer ctx.Close()

	ctx.SetShift(true)

	if ctx.ledMask != 1 {
		t.Errorf("Unexpected update mask value, expected %v, got %v", 1, ctx.ledMask)
	}
	if ctx.updateMask != 1 {
		t.Errorf("Unexpected update mask value, expected %v, got %v", 1, ctx.updateMask)
	}
}

func TestShiftOff(t *testing.T) {
	ctx := NewContext()
	defer ctx.Close()

	ctx.SetShift(false)

	if ctx.ledMask != 0 {
		t.Errorf("Unexpected update mask value, expected %v, got %v", 0, ctx.ledMask)
	}
	if ctx.updateMask != 1 {
		t.Errorf("Unexpected update mask value, expected %v, got %v", 1, ctx.updateMask)
	}
}

func TestMFDBrightness(t *testing.T) {
	updateMask := uint32(1 << updateBrightnessMFD)
	ctx := NewContext()
	defer ctx.Close()

	for i := uint16(0); i <= 0x80; i++ {
		ctx.updateMask = 0
		err := ctx.SetMFDBrightness(i)
		if err != nil {
			t.Errorf("Unexpected error setting MFD brightness to %v: %v", i, err)
			continue
		}

		if ctx.updateMask != updateMask {
			t.Errorf("Update mask wrong value - exp %v, got %v", updateMask, ctx.updateMask)
		}

		if ctx.mfdBrightness != i {
			t.Errorf("MFD Brightness wrong value - exp %v, got %v", i, ctx.mfdBrightness)
		}
	}
}

func TestMFDText(t *testing.T) {
	ctx := NewContext()
	defer ctx.Close()

	text := []byte("abcdefghijklmnopqr")

	for line := uint8(0); line < mfdLines; line++ {
		updateMask := uint32(1 << (updateMfdLine1 + uint32(line)))
		for i := 0; i <= len(text); i++ {
			ctx.updateMask = 0
			data := text[:i]
			err := ctx.SetMFDText(line, data)
			if err != nil {
				t.Error("Unexpected error", err)
				continue
			}

			lim := i
			if lim > mfdLineSize {
				lim = mfdLineSize
			}
			data = data[:lim]

			if ctx.updateMask != updateMask {
				t.Errorf("Update mask wrong value - exp %v, got %v", updateMask, ctx.updateMask)
			}

			if !bytes.Equal(ctx.mfdLine[line], data) {
				t.Errorf("MFD text wrong value - exp %v, got %v", data, ctx.mfdLine[line])
			}
		}
	}

	// Test setting an invalid line
	errInvalid := ErrInvalidParam("line number out of range")
	err := ctx.SetMFDText(3, []byte("foobar"))
	if err.Error() != errInvalid.Error() {
		t.Errorf("Mismatch in error strings, exp %v, got %v", errInvalid, err)
	}
}
