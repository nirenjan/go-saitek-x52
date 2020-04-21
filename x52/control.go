package x52

import (
// "github.com/google/gousb"
)

// Raw sends a raw vendor control packet to the device
func (ctx *Context) Raw(index, value uint16) error {
	// TODO: Implement
	return nil
}

// Update updates the X52 with the saved data
func (ctx *Context) Update() error {
	updated := ctx.updateMask
	written := uint32(0)
	var value uint16
	var index uint16
	var err error

	for i := updateShift; i < updateMax; i++ {
		if bitTest(updated, i) {
			// Bit is not set
			continue
		}

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
		}
	}

	return nil
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
