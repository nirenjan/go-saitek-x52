package main

// LED function test

import (
	"flag"
	"fmt"

	"nirenjan.org/saitek-x52/x52"
)

func init() {
	flag.Var(&TC{"LED state", testLED}, "led", "Test LED states")
}

func testLED(ctx *x52.Context) error {
	// No point testing the LEDs if the device doesn't support it
	if !ctx.HasFeature(x52.FeatureLED) && !mockTests {
		return nil
	}

	bar := progressBar("LED State", 81)

	onOffLEDs := []x52.LED{x52.LedFire, x52.LedThrottle}
	colorLEDs := []x52.LED{x52.LedA, x52.LedB, x52.LedD, x52.LedE, x52.LedT1, x52.LedT2, x52.LedT3, x52.LedPOV, x52.LedClutch}

	onOffStates := []x52.LedState{x52.LedOff, x52.LedOn}
	colorStates := []x52.LedState{x52.LedOff, x52.LedRed, x52.LedAmber, x52.LedGreen}

	for _, led := range onOffLEDs {
		for i := 0; i < 2; i++ {
			for _, state := range onOffStates {
				bar.Describe(fmt.Sprintf("%8v %-5v", led, state))
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

	for _, led := range colorLEDs {
		for i := 0; i < 2; i++ {
			for _, state := range colorStates {
				bar.Describe(fmt.Sprintf("%8v %-5v", led, state))
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

	bar.Describe("LED State")
	bar.Finish()

	return nil
}
