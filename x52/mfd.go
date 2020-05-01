package x52

// SetMFDText sets the display on the given MFD line.
// The data must be in the code page recognized by the MFD.
// This function only accepts line lengths of up to 16 bytes, with any
// additional data being silently discarded.
func (ctx *Context) SetMFDText(line uint8, data []byte) error {
	if line >= mfdLines {
		return errInvalidParam("line number out of range")
	}

	// Restrict the data to the line size of the MFD
	if len(data) > mfdLineSize {
		data = data[:mfdLineSize]
	}

	ctx.mfdLine[line] = data
	bitSet(&ctx.updateMask, updateMfdLine1+uint32(line))

	return nil
}

// setBrightness sets the brightness of either the MFD or the LEDs
func (ctx *Context) setBrightness(led bool, brightness uint16) error {
	if led {
		ctx.ledBrightness = brightness
		bitSet(&ctx.updateMask, updateBrightnessLED)
	} else {
		ctx.mfdBrightness = brightness
		bitSet(&ctx.updateMask, updateBrightnessMFD)
	}

	return nil
}

// SetMFDBrightness will set the brightness of the MFD.
func (ctx *Context) SetMFDBrightness(brightness uint16) error {
	return ctx.setBrightness(false, brightness)
}

// SetLEDBrightness will set the brightness of the LED.
func (ctx *Context) SetLEDBrightness(brightness uint16) error {
	return ctx.setBrightness(true, brightness)
}

// setBlinkShift will enable or disable the blink/shift functionality
func (ctx *Context) setBlinkShift(enable bool, bit uint32) error {
	if enable {
		bitSet(&ctx.ledMask, bit)
	} else {
		bitClear(&ctx.ledMask, bit)
	}

	bitSet(&ctx.updateMask, bit)

	return nil
}

// SetBlink will enable or disable the blink functionality
func (ctx *Context) SetBlink(enable bool) error {
	return ctx.setBlinkShift(enable, updatePOVBlink)
}

// SetShift will enable or disable the shift functionality
func (ctx *Context) SetShift(enable bool) error {
	return ctx.setBlinkShift(enable, updateShift)
}
