package x52

import (
	"github.com/google/gousb"
)

// Raw sends a raw vendor control packet to the device
func (ctx *Context) Raw(index, value uint16) error {
	if ctx.device == nil {
		ctx.log(logWarning, "not connected")
		return errNotConnected(nil)
	}

	// gousb takes care of retries internally, so we don't have to
	// do it ourselves
	ctx.logf(logDebug, "sending raw %04x %04x", index, value)
	_, err := ctx.device.Control(
		gousb.ControlVendor|gousb.ControlDevice|gousb.ControlOut,
		0x91, value, index, nil)
	if err != nil {
		ctx.logf(logError, "error updating device: %v", err)

		if err == gousb.ErrorNoDevice {
			// Device has been unplugged, close it
			ctx.devClose()
			ctx.log(logWarning, "device has been disconnected")

			return errNotConnected(err)
		}

		return err
	}

	return nil
}

// Update updates the X52 with the saved data
func (ctx *Context) Update() error {
	updated := ctx.updateMask
	written := uint32(0)
	var value uint16
	var index uint16
	var err error

	ctx.logf(logDebug, "updated bitmask %08x", updated)
	for i := updateShift; i < updateMax; i++ {
		if !bitTest(updated, i) {
			// Bit is not set
			continue
		}

		ctx.log(logDebug, "checking bit", i)

		switch i {
		case updateShift:
			value = 0x50 // Shift OFF
			if bitTest(ctx.ledMask, updateShift) {
				// Shift ON
				value |= 1
			}
			// Shift indicator
			index = 0xfd

			err = ctx.Raw(index, value)

		case updateLedFire,
			updateLedARed, updateLedAGreen,
			updateLedBRed, updateLedBGreen,
			updateLedDRed, updateLedDGreen,
			updateLedERed, updateLedEGreen,
			updateLedT1Red, updateLedT1Green,
			updateLedT2Red, updateLedT2Green,
			updateLedT3Red, updateLedT3Green,
			updateLedPOVRed, updateLedPOVGreen,
			updateLedClutchRed, updateLedClutchGreen,
			updateLedThrottle:

			value = uint16(i << 8)
			if bitTest(ctx.ledMask, i) {
				value |= 1
			}
			err = ctx.Raw(0xb8, value)

		case updateMfdLine1, updateMfdLine2, updateMfdLine3:
			err = ctx.writeLine(uint8(i - updateMfdLine1))

		case updatePOVBlink:
			value = 0x50 // Blink OFF
			if bitTest(ctx.ledMask, updateShift) {
				// Blink ON
				value |= 1
			}
			// Blink indicator
			index = 0xb4

			err = ctx.Raw(index, value)

		case updateBrightnessMFD:
			index = 0xb1
			value = ctx.mfdBrightness
			err = ctx.Raw(index, value)

		case updateBrightnessLED:
			index = 0xb2
			value = ctx.ledBrightness
			err = ctx.Raw(index, value)

		default:
			err = nil
		}

		if err == nil {
			bitSet(&written, i)
			bitClear(&updated, i)
		} else {
			break
		}
	}

	return err
}

func (ctx *Context) writeLine(line uint8) error {
	data := ctx.mfdLine[line]
	if len(data)%2 != 0 {
		data = append(data, byte(0))
	}

	var index uint16
	var value uint16

	// Clear the line first
	index = 0xd8 | uint16(1<<line)
	value = 0
	err := ctx.Raw(index, value)
	if err != nil {
		return err
	}

	// Write the line
	index = 0xd0 | uint16(1<<line)
	for i := 0; i < len(data); i += 2 {
		value = uint16(data[i+1])<<8 | uint16(data[i])
		err = ctx.Raw(index, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Context) writeDate() error {
	t, _ := convertTime(ctx.time)

	var dd, mm, yy uint16
	switch ctx.dateFormat {
	case DateFormatDDMMYY:
		dd = uint16(t.day)
		mm = uint16(t.month)
		yy = uint16(t.year)

	case DateFormatMMDDYY:
		mm = uint16(t.day)
		dd = uint16(t.month)
		yy = uint16(t.year)

	case DateFormatYYMMDD:
		yy = uint16(t.day)
		mm = uint16(t.month)
		dd = uint16(t.year)

	default:
		return errStructCorrupted("")
	}

	ddmm := mm<<8 | dd
	err := ctx.Raw(0xc4, ddmm)
	if err != nil {
		return err
	}

	return ctx.Raw(0xc8, yy)
}

func (ctx *Context) writeTime() error {
	_, t := convertTime(ctx.time)

	cf := uint16(ctx.timeFormat[Clock1])
	hh := uint16(t.hour)
	mm := uint16(t.minute)

	value := (cf << 15) | (hh << 8) | mm
	return ctx.Raw(0xc0, value)
}

func (ctx *Context) writeOffset(clock ClockID) error {
	offs := ctx.computeOffset(clock)
	cf := uint16(ctx.timeFormat[clock])
	value := (cf << 15) | offs
	index := uint16(clock) | 0xc0
	return ctx.Raw(index, value)
}
