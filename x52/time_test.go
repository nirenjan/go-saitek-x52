package x52

import (
	"fmt"
	"testing"
	"time"
)

// TestSetLocation verifies that setting the location on clock 2/3 updates
// the updateMask correctly, and sets the right members in the context
func TestSetLocation(t *testing.T) {
	tests := []struct {
		clock      ClockID
		loc        *time.Location
		updateMask uint32
		err        error
	}{
		{Clock1, nil, 0, errInvalidParam("cannot set location of primary clock")},
		{Clock2, nil, 0, errInvalidParam("must specify a valid location")},
		{Clock3, nil, 0, errInvalidParam("must specify a valid location")},
		{ClockID(3), nil, 0, errInvalidParam("invalid clock ID")},
		{Clock2, time.Local, 0x20000000, nil},
		{Clock3, time.Local, 0x40000000, nil},
		{Clock2, time.UTC, 0x20000000, nil},
		{Clock3, time.UTC, 0x40000000, nil},
	}

	ctx := NewContext()
	defer ctx.Close()
	defer ctx.Close()
	for i, tc := range tests {
		ctx.initialize()

		tcID := fmt.Sprintf("%s%d", t.Name(), i+1)
		err := ctx.SetLocation(tc.clock, tc.loc)
		if err != nil {
			if tc.err == nil {
				t.Errorf("%v: unexpected error %v", tcID, err)
			} else if err.Error() != tc.err.Error() {
				t.Errorf("%v: mismatched error messages\n\tgot: %v\n\texp: %v\n",
					tcID, err, tc.err)
			}
		} else {
			if tc.err != nil {
				t.Errorf("%v: expected error %q, but got none", tcID, tc.err)
			} else if ctx.updateMask != tc.updateMask {
				t.Errorf("%v: mismatched update masks\n\tgot: %08x\n\texp: %08x\n",
					tcID, ctx.updateMask, tc.updateMask)
			}
		}
	}
}

// TestSetClockFormat verifies that setting the clock format works correctly
// and sets the right updatemask
func TestSetClockFormat(t *testing.T) {
	tests := []struct {
		clock      ClockID
		format     ClockFormat
		updateMask uint32
		err        error
	}{
		{Clock1, ClockFormat(2), 0, errInvalidParam("invalid clock format")},
		{Clock2, ClockFormat(2), 0, errInvalidParam("invalid clock format")},
		{Clock3, ClockFormat(2), 0, errInvalidParam("invalid clock format")},
		{ClockID(3), ClockFormat(2), 0, errInvalidParam("invalid clock format")},
		{ClockID(3), ClockFormat12Hr, 0, errInvalidParam("invalid clock ID")},

		{Clock1, ClockFormat12Hr, 0x10000000, nil},
		{Clock1, ClockFormat24Hr, 0x10000000, nil},
		{Clock2, ClockFormat12Hr, 0x20000000, nil},
		{Clock2, ClockFormat24Hr, 0x20000000, nil},
		{Clock3, ClockFormat12Hr, 0x40000000, nil},
		{Clock3, ClockFormat24Hr, 0x40000000, nil},
	}

	ctx := NewContext()
	defer ctx.Close()
	for i, tc := range tests {
		ctx.initialize()

		tcID := fmt.Sprintf("%s%d", t.Name(), i+1)
		err := ctx.SetClockFormat(tc.clock, tc.format)
		if err != nil {
			if tc.err == nil {
				t.Errorf("%v: unexpected error %v", tcID, err)
			} else if err.Error() != tc.err.Error() {
				t.Errorf("%v: mismatched error messages\n\tgot: %v\n\texp: %v\n",
					tcID, err, tc.err)
			}
		} else {
			if tc.err != nil {
				t.Errorf("%v: expected error %q, but got none", tcID, tc.err)
			} else if ctx.updateMask != tc.updateMask {
				t.Errorf("%v: mismatched update masks\n\tgot: %08x\n\texp: %08x\n",
					tcID, ctx.updateMask, tc.updateMask)
			} else if ctx.timeFormat[tc.clock] != tc.format {
				t.Errorf("%v: wrong value saved\n\tgot: %v\n\texp: %v\n",
					tcID, ctx.timeFormat[tc.clock], tc.format)
			}
		}
	}
}

// TestSetDateFormat verifies that setting the date format works correctly
// and sets the right updatemask
func TestSetDateFormat(t *testing.T) {
	tests := []struct {
		format     DateFormat
		updateMask uint32
		err        error
	}{
		{DateFormat(3), 0, errInvalidParam("invalid date format")},

		{DateFormatDDMMYY, 0x08000000, nil},
		{DateFormatMMDDYY, 0x08000000, nil},
		{DateFormatYYMMDD, 0x08000000, nil},
	}

	ctx := NewContext()
	defer ctx.Close()
	for i, tc := range tests {
		ctx.initialize()

		tcID := fmt.Sprintf("%s%d", t.Name(), i+1)
		err := ctx.SetDateFormat(tc.format)
		if err != nil {
			if tc.err == nil {
				t.Errorf("%v: unexpected error %v", tcID, err)
			} else if err.Error() != tc.err.Error() {
				t.Errorf("%v: mismatched error messages\n\tgot: %v\n\texp: %v\n",
					tcID, err, tc.err)
			}
		} else {
			if tc.err != nil {
				t.Errorf("%v: expected error %q, but got none", tcID, tc.err)
			} else if ctx.updateMask != tc.updateMask {
				t.Errorf("%v: mismatched update masks\n\tgot: %08x\n\texp: %08x\n",
					tcID, ctx.updateMask, tc.updateMask)
			} else if ctx.dateFormat != tc.format {
				t.Errorf("%v: wrong value saved\n\tgot: %v\n\texp: %v\n",
					tcID, ctx.dateFormat, tc.format)
			}
		}
	}
}

// TestSetTime verifies that the SetTime function works as expected
// Unlike previous test cases, this doesn't reinitialize the entire context
// but only the updateMask, in order to verify that the updateMask is set
// correctly in all cases
func TestSetTime(t *testing.T) {
	ctx := NewContext()
	defer ctx.Close()

	zone := time.FixedZone("UTC-8", -8*60*60)
	now := time.Date(2006, 1, 2, 15, 4, 5, 0, zone)

	tests := []struct {
		time       time.Time
		updateMask uint32
	}{
		// First case, use the zero time, and don't change the location
		// in order to simulate the no-change condition
		{time.Time{}, 0x00000000},

		// Set the clock to 2006-01-02T15:04:05-0800, expect all clock fields
		// to get updated
		{now, 0x78000000},

		// Set the clock to 2006-01-02T15:04:05-0800 again, but expect no
		// updates since the clock is the same
		{now, 0x00000000},

		// Set the clock to 2006-01-02T15:04:59-0800, but expect no updates
		// since the hour and minute display hasn't changed
		{now.Add(54 * time.Second), 0x00000000},

		// Set the clock to 2006-01-02T15:05:05-0800, and expect only the time
		// field to be updated
		{now.Add(60 * time.Second), 0x10000000},

		// Set the clock to 2006-01-03T15:05:05-0800, and expect only the date
		// field to be updated
		{now.Add(86460 * time.Second), 0x08000000},

		// Set the clock back to 2006-01-02T15:04:05-0800, expect both date
		// and time fields to be updated
		{now, 0x18000000},
	}

	for i, tc := range tests {
		ctx.updateMask = 0

		tcID := fmt.Sprintf("%s%d", t.Name(), i+1)
		t.Log(tcID)
		t.Log(ctx.time.Format(time.RFC3339), ctx.timeZone[Clock1])
		t.Log(tc.time.Format(time.RFC3339), tc.time.Location(), tc.updateMask)
		t.Log("\n")
		err := ctx.SetTime(tc.time)
		if err != nil {
			t.Errorf("%v: unexpected error %v", tcID, err)
		} else {
			if ctx.updateMask != tc.updateMask {
				t.Errorf("%v: mismatched update masks\n\tgot: %08x\n\texp: %08x\n",
					tcID, ctx.updateMask, tc.updateMask)
			} else if ctx.updateMask != 0 && !ctx.time.Equal(tc.time) {
				t.Errorf("%v: wrong value saved\n\tgot: %v\n\texp: %v\n",
					tcID, ctx.time, tc.time)
			}
		}
	}
}

// TestConvertOffset verifies that the offset calculation computes the right
// offset for the secondary and tertiary clocks, even when the timezone
// difference is very large.
func TestConvertOffset(t *testing.T) {
	tests := []struct {
		offset1 int
		offset2 int
	}{
		// PST, GMT
		{-8 * 60 * 60, 0},
		// GMT, PST
		{0, -8 * 60 * 60},
		// LINT, PST
		{14 * 60 * 60, -8 * 60 * 60},
		// LINT, IDLW
		{14 * 60 * 60, -12 * 60 * 60},
		// LINT, PST
		{-8 * 60 * 60, 14 * 60 * 60},
		// LINT, IDLW
		{-12 * 60 * 60, 14 * 60 * 60},
	}

	ctx := NewContext()
	defer ctx.Close()

	for i, tc := range tests {
		tcID := fmt.Sprintf("%s%d", t.Name(), i+1)

		// Set a default time so that we can accurately compute the offsets
		ctx.initialize()
		ctx.SetTime(time.Date(2006, 1, 2, 15, 4, 5, 0, time.FixedZone("PRI", tc.offset1)))
		ctx.SetLocation(Clock2, time.FixedZone("SEC", tc.offset2))

		// Compute what the offset should be
		offset := (tc.offset2 - tc.offset1) / 60

		// If the offset exceeds the range -12h to +12h, add/subtract 24h until
		// it is in the range
		t.Logf("%v: adjusting offset", tcID)
		t.Log("\twas ", offset, offset/60, offset%60)
		for offset < -12*60 || offset > 12*60 {
			if offset < -12*60 {
				offset += 24 * 60
			} else if offset > 12*60 {
				offset -= 24 * 60
			}
			t.Log("\tnow ", offset, offset/60, offset%60)
		}

		var expOffs uint16
		if offset < 0 {
			expOffs = 1 << 10
			offset = -offset
		}
		expOffs |= uint16(offset)

		gotOffs := ctx.computeOffset(Clock2)
		t.Logf("%v:\n\tgot: %04x\n\texp: %04x\n", tcID, gotOffs, expOffs)
		if gotOffs != expOffs {
			t.Error("mismatched offset values")
		}
	}
}
