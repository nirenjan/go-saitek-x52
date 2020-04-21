package x52

// SetLed sets the state of the given LED. Not all LEDs support all states,
// LedFire and LedThrottle only support LedOn and LedOff states,
// the remaining LEDs support every state except LedOn.
// **Limitation**: This function will not work on a non-pro X52 at this time.
func (ctx *Context) SetLed(led LED, state LedState) error {
	// Make sure that this is a supported device
	// The non-pro X52 doesn't support setting LED states
	if !bitTest(ctx.featureFlags, featureLed) {
		return ErrNotSupported("setting LED state")
	}

	switch led {
	case LedFire, LedThrottle:
		if state == LedOff {
			bitClear(&ctx.ledMask, uint32(led))
			bitSet(&ctx.updateMask, uint32(led))
		} else if state == LedOn {
			bitSet(&ctx.ledMask, uint32(led))
			bitSet(&ctx.updateMask, uint32(led))
		} else {
			return ErrNotSupported("invalid state for on/off LED")
		}

	case LedA, LedB, LedD, LedE, LedT1, LedT2, LedT3, LedPOV, LedClutch:
		led_id := uint32(led)
		switch state {
		case LedOff:
			bitClear(&ctx.ledMask, led_id+0)
			bitClear(&ctx.ledMask, led_id+1)

		case LedRed:
			bitSet(&ctx.ledMask, led_id+0)
			bitClear(&ctx.ledMask, led_id+1)

		case LedAmber:
			bitSet(&ctx.ledMask, led_id+0)
			bitSet(&ctx.ledMask, led_id+1)

		case LedGreen:
			bitClear(&ctx.ledMask, led_id+0)
			bitSet(&ctx.ledMask, led_id+1)

		default:
			return ErrNotSupported("invalid state for color LED")
		}

		bitSet(&ctx.updateMask, led_id+0)
		bitSet(&ctx.updateMask, led_id+1)

	default:
		return ErrNotSupported("invalid LED identifier")
	}

	return nil
}
