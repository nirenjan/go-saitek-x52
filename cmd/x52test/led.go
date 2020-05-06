package main

// LED function test

import (
	"flag"
	"fmt"
	"strings"

	"nirenjan.org/saitek-x52/x52"
)

func init() {
	flag.Var(&TC{"LED state", testLED}, "led", "Test LED states")
}

func testLED(ctx *x52.Context) error {
	// Run the blink function first
	if err := testBlink(ctx); err != nil {
		return err
	}

	// No point testing the LEDs if the device doesn't support it
	if !ctx.HasFeature(x52.FeatureLED) && !mockTests {
		return nil
	}

	testLED := func(leds []x52.LED, states []x52.LedState) error {
		const repeatCount = 2
		for _, led := range leds {
			// Build the progressbar text
			text := fmt.Sprintf("%v ", led)
			for _, state := range states {
				text += fmt.Sprintf("%v/", state)
			}
			text = strings.TrimRight(text, "/")
			bar := progressBar(text, repeatCount*len(states))

			// Wait for 1 second before actually starting the test
			delayMs(1000)

			for i := 0; i < repeatCount; i++ {
				for _, state := range states {
					ctx.SetLed(led, state)
					if err := updateDev(ctx, bar); err != nil {
						return err
					}
					delayMs(250)
					bar.Add(1)
				}
			}
			delayMs(500)
		}

		return nil
	}

	onOffLEDs := []x52.LED{x52.LedFire, x52.LedThrottle}
	onOffStates := []x52.LedState{x52.LedOff, x52.LedOn}
	if err := testLED(onOffLEDs, onOffStates); err != nil {
		return err
	}

	colorLEDs := []x52.LED{x52.LedA, x52.LedB, x52.LedD, x52.LedE, x52.LedT1, x52.LedT2, x52.LedT3, x52.LedPOV, x52.LedClutch}
	colorStates := []x52.LedState{x52.LedOff, x52.LedRed, x52.LedAmber, x52.LedGreen}
	if err := testLED(colorLEDs, colorStates); err != nil {
		return err
	}

	return nil
}

func testBlink(ctx *x52.Context) error {
	bar := progressBar("LED Blink On/Off", 2)

	for i := 0; i < 2; i++ {
		ctx.SetBlink(i == 0)
		if err := updateDev(ctx, bar); err != nil {
			return err
		}
		delayMs(2000)
		bar.Add(1)
	}

	return nil
}
