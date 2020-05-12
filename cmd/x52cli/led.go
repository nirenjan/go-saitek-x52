package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"nirenjan.org/saitek-x52/x52"
)

var ledCommand *cobra.Command

func init() {
	ledCommand = &cobra.Command{
		Use:   "led",
		Short: "Control the LED state",
		Long: `Set the LED state of the individual LEDs on the connected X52Pro
joystick. This command is not supported on the non-pro X52 due to
hardware limitations.
`,
	}

	register := func(leds []x52.LED, states []x52.LedState) {
		for _, led := range leds {
			ledStr := strings.ToLower(led.String())
			ledCmd := &cobra.Command{
				Use:   ledStr,
				Short: fmt.Sprintf("Set the state of the %v LED on the X52Pro", led),
			}

			for _, state := range states {
				stateStr := strings.ToLower(state.String())
				stateCmd := &cobra.Command{
					Use:   stateStr,
					Short: fmt.Sprintf("Set the %v LED to %v", led, state),
					Run:   ledHandler(led, state),
				}

				ledCmd.AddCommand(stateCmd)
			}

			ledCommand.AddCommand(ledCmd)
		}
	}

	onOffLEDs := []x52.LED{x52.LedFire, x52.LedThrottle}
	onOffStates := []x52.LedState{x52.LedOff, x52.LedOn}
	colorLEDs := []x52.LED{x52.LedA, x52.LedB, x52.LedD, x52.LedE,
		x52.LedT1, x52.LedT2, x52.LedT3, x52.LedPOV, x52.LedClutch}
	colorStates := []x52.LedState{x52.LedOff, x52.LedRed, x52.LedAmber, x52.LedGreen}

	register(onOffLEDs, onOffStates)
	register(colorLEDs, colorStates)
}

func ledHandler(led x52.LED, state x52.LedState) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		ctx := connectToX52()

		if cliVerbose {
			fmt.Println(led, state)
		}

		ctx.SetLed(led, state)
		ctx.Update()
	}
}
