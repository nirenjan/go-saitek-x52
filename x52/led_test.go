package x52

import (
	"testing"
)

func TestLeds(t *testing.T) {
	tests := []struct {
		led        LED
		state      LedState
		ledMask    uint32
		updateMask uint32
		errexp     bool
		errmsg     string
	}{
		{LedFire, LedOff, 0x00000000, 0x00000002, false, ""},
		{LedFire, LedOn, 0x00000002, 0x00000002, false, ""},
		{LedFire, LedRed, 0x00000002, 0x00000002, true, "invalid state for on/off LED"},
		{LedFire, LedAmber, 0x00000002, 0x00000002, true, "invalid state for on/off LED"},
		{LedFire, LedGreen, 0x00000002, 0x00000002, true, "invalid state for on/off LED"},

		{LedA, LedOff, 0x00000000, 0x0000000c, false, ""},
		{LedA, LedOn, 0x00000000, 0x00000000, true, "invalid state for color LED"},
		{LedA, LedRed, 0x00000004, 0x0000000c, false, ""},
		{LedA, LedAmber, 0x0000000c, 0x0000000c, false, ""},
		{LedA, LedGreen, 0x00000008, 0x0000000c, false, ""},

		{LedB, LedOff, 0x00000000, 0x00000030, false, ""},
		{LedB, LedOn, 0x00000000, 0x00000000, true, "invalid state for color LED"},
		{LedB, LedRed, 0x00000010, 0x00000030, false, ""},
		{LedB, LedAmber, 0x00000030, 0x00000030, false, ""},
		{LedB, LedGreen, 0x00000020, 0x00000030, false, ""},

		{LedD, LedOff, 0x00000000, 0x000000c0, false, ""},
		{LedD, LedOn, 0x00000000, 0x00000000, true, "invalid state for color LED"},
		{LedD, LedRed, 0x00000040, 0x000000c0, false, ""},
		{LedD, LedAmber, 0x000000c0, 0x000000c0, false, ""},
		{LedD, LedGreen, 0x00000080, 0x000000c0, false, ""},

		{LedE, LedOff, 0x00000000, 0x00000300, false, ""},
		{LedE, LedOn, 0x00000000, 0x00000000, true, "invalid state for color LED"},
		{LedE, LedRed, 0x00000100, 0x00000300, false, ""},
		{LedE, LedAmber, 0x00000300, 0x00000300, false, ""},
		{LedE, LedGreen, 0x00000200, 0x00000300, false, ""},

		{LedT1, LedOff, 0x00000000, 0x00000c00, false, ""},
		{LedT1, LedOn, 0x00000000, 0x00000000, true, "invalid state for color LED"},
		{LedT1, LedRed, 0x00000400, 0x00000c00, false, ""},
		{LedT1, LedAmber, 0x00000c00, 0x00000c00, false, ""},
		{LedT1, LedGreen, 0x00000800, 0x00000c00, false, ""},

		{LedT2, LedOff, 0x00000000, 0x00003000, false, ""},
		{LedT2, LedOn, 0x00000000, 0x00000000, true, "invalid state for color LED"},
		{LedT2, LedRed, 0x00001000, 0x00003000, false, ""},
		{LedT2, LedAmber, 0x00003000, 0x00003000, false, ""},
		{LedT2, LedGreen, 0x00002000, 0x00003000, false, ""},

		{LedT3, LedOff, 0x00000000, 0x0000c000, false, ""},
		{LedT3, LedOn, 0x00000000, 0x00000000, true, "invalid state for color LED"},
		{LedT3, LedRed, 0x00004000, 0x0000c000, false, ""},
		{LedT3, LedAmber, 0x0000c000, 0x0000c000, false, ""},
		{LedT3, LedGreen, 0x00008000, 0x0000c000, false, ""},

		{LedPOV, LedOff, 0x00000000, 0x00030000, false, ""},
		{LedPOV, LedOn, 0x00000000, 0x00000000, true, "invalid state for color LED"},
		{LedPOV, LedRed, 0x00010000, 0x00030000, false, ""},
		{LedPOV, LedAmber, 0x00030000, 0x00030000, false, ""},
		{LedPOV, LedGreen, 0x00020000, 0x00030000, false, ""},

		{LedClutch, LedOff, 0x00000000, 0x000c0000, false, ""},
		{LedClutch, LedOn, 0x00000000, 0x00000000, true, "invalid state for color LED"},
		{LedClutch, LedRed, 0x00040000, 0x000c0000, false, ""},
		{LedClutch, LedAmber, 0x000c0000, 0x000c0000, false, ""},
		{LedClutch, LedGreen, 0x00080000, 0x000c0000, false, ""},

		{LedThrottle, LedOff, 0x00000000, 0x00100000, false, ""},
		{LedThrottle, LedOn, 0x00100000, 0x00100000, false, ""},
		{LedThrottle, LedRed, 0x00000000, 0x00000000, true, "invalid state for on/off LED"},
		{LedThrottle, LedAmber, 0x00000000, 0x00000000, true, "invalid state for on/off LED"},
		{LedThrottle, LedGreen, 0x00000000, 0x00000000, true, "invalid state for on/off LED"},

		{LED(21), LedOff, 0x00000000, 0x00000000, true, "invalid LED identifier"},
	}

	ctx := NewContext()
	defer ctx.Close()

	// Check that if the LED feature is not enabled,
	// then SetLed will return an error
	if err := ctx.SetLed(LedFire, LedOff); err == nil {
		t.Error("Setting LED when featureLed is not set does not return an error")
	}

	// Enable setting LEDs
	bitSet(&ctx.featureFlags, featureLed)

	for i, tc := range tests {
		ctx.ledMask = tc.updateMask
		ctx.updateMask = 0

		err := ctx.SetLed(tc.led, tc.state)
		if !tc.errexp {
			if err != nil {
				t.Logf("id %v %v\n", i, tc)
				t.Logf("%#v\n", ctx)
				t.Error("Unexpected error:", err)
				continue
			}

			if ctx.ledMask != tc.ledMask || ctx.updateMask != tc.updateMask {
				t.Logf("id %v %v\n", i, tc)
				t.Logf("%#v\n", ctx)
				t.Errorf("Unexpected mask values:\n\texp: %08x %08x\n\tgot: %08x %08x\n",
					tc.ledMask, tc.updateMask, ctx.ledMask, ctx.updateMask)
			}
		} else {
			experr := ErrNotSupported(tc.errmsg)
			if experr.Error() != err.Error() {
				t.Logf("id %v %v\n", i, tc)
				t.Logf("%#v\n", ctx)
				t.Errorf("Unexpected error message:\n\texp: %v\n\tgot: %v\n", experr, err)
			}
		}
	}
}
