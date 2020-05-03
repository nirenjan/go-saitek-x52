package x52

// This file controls the date and time display on the X52 MFD

import (
	"time"
)

type tmDate struct {
	year  int
	month time.Month
	day   int
}

type tmTime struct {
	hour   int
	minute int
}

// convertTime converts a time.Time to a struct tm. This ignores the seconds
// because they are not displayed in the X52 MFD
func convertTime(t time.Time) (dt tmDate, tm tmTime) {
	dt.year, dt.month, dt.day = t.Date()
	tm.hour, tm.minute, _ = t.Clock()

	return
}

// SetTime sets the time of the primary clock. The secondary and tertiary clocks
// are derived by setting a programmable offset from the primary clock.
func (ctx *Context) SetTime(t time.Time) error {
	if t.Location() != ctx.timeZone[Clock1] {
		// Location has changed, we need to update all the clocks and date
		ctx.timeZone[Clock1] = t.Location()

		for i := updateDate; i <= updateOffs2; i++ {
			bitSet(&ctx.updateMask, i)
		}
	}

	savedDt, savedTm := convertTime(ctx.time)
	inputDt, inputTm := convertTime(t)

	if savedDt == inputDt && savedTm == inputTm {
		// No change to display time
		return nil
	}

	ctx.time = t

	if savedTm != inputTm {
		bitSet(&ctx.updateMask, updateTime)
	}
	if savedDt != inputDt {
		bitSet(&ctx.updateMask, updateDate)
	}
	return nil
}

// SetLocation updates the location of the given clock. You may not update the
// location of the primary clock, as it is computed when you call SetTime
func (ctx *Context) SetLocation(clock ClockID, loc *time.Location) error {
	switch clock {
	case Clock1:
		return errInvalidParam("cannot set location of primary clock")

	case Clock2, Clock3:
		if loc == nil {
			return errInvalidParam("must specify a valid location")
		}

		ctx.timeZone[clock] = loc
		bitSet(&ctx.updateMask, updateTime+uint32(clock))

	default:
		return errInvalidParam("invalid clock ID")
	}

	return nil
}

// SetClockFormat sets the clock format of the given clock
func (ctx *Context) SetClockFormat(clock ClockID, format ClockFormat) error {
	// Validate parameters
	switch format {
	case ClockFormat12Hr, ClockFormat24Hr:
		// acceptable values, proceed
		break

	default:
		return errInvalidParam("invalid clock format")
	}

	switch clock {
	case Clock1, Clock2, Clock3:
		ctx.timeFormat[clock] = format
		bitSet(&ctx.updateMask, updateTime+uint32(clock))

	default:
		return errInvalidParam("invalid clock ID")
	}

	return nil
}

// SetDateFormat sets the date format
func (ctx *Context) SetDateFormat(format DateFormat) error {
	// Validate parameters
	switch format {
	case DateFormatDDMMYY, DateFormatMMDDYY, DateFormatYYMMDD:
		// acceptable values, proceed
		break

	default:
		return errInvalidParam("invalid date format")
	}

	ctx.dateFormat = format
	bitSet(&ctx.updateMask, updateDate)
	return nil
}

// computeOffset computes the offset of the secondary and tertiary clocks from
// the base clock. The X52 only supports offsets in the range [-1023, +1023],
// but the wide range of timezones means that the actual offsets can be outside
// this range. This function handles converting a large offset into one that
// can be fit within the range, using the logic that t+16h === t-8h
func (ctx *Context) computeOffset(clock ClockID) uint16 {
	t1 := ctx.time

	// Make sure that we have a valid timezone, otherwise, treat it as the
	// same timezone as clock1 (offset 0)
	tz := ctx.timeZone[clock]
	if tz == nil {
		return 0
	}

	t2 := t1.In(tz)
	z1, o1 := t1.Zone()
	z2, o2 := t2.Zone()

	ctx.log(logDebug, "Primary clock timezone", z1, o1)
	ctx.log(logDebug, "Clock", clock+1, " timezone", z2, o2)

	// The returned offsets are in seconds east of GMT. We need to compute the
	// offset between the two clocks in minutes
	offset := (o2 - o1) / 60
	negative := (offset < 0)
	if negative {
		offset = -offset
	}
	ctx.log(logDebug, "Raw offset", offset, "negative", negative)

	// The X52 packet formats takes in a sign bit, followed by the offset value
	// which is a 10 bit value. This can have a maximum value of 1023, but the
	// raw offset from the base clock can easily exceed this. In order to handle
	// this, we repeatedly subtract a 24 hour offset, in order to convert the
	// offsets of e.g. -16h into +8h. The clock selector doesn't change the date
	// display on the MFD, so we only need to care about the time value.
	for offset > 1023 {
		offset -= 1440 // Subtract 24 hours
	}

	// It is possible that we've subtracted too much and the offset has gone
	// negative again. In this case, we need to flip the negative flag, and
	// make the offset a positive value
	if offset < 0 {
		offset = -offset
		negative = !negative
	}

	var neg uint16
	if negative {
		neg = 1 << 10
	}
	return (neg | uint16(offset))
}
