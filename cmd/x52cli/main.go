package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"nirenjan.org/saitek-x52/x52"
)

var cliVerbose bool
var rootCmd = &cobra.Command{
	Use:   "x52cli",
	Short: "x52cli is a utility program to control the X52/X52Pro LEDs and MFD",
}

func main() {
	// Add flags to the root command
	rootCmd.PersistentFlags().BoolVarP(&cliVerbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(ledCommand)
	rootCmd.AddCommand(mfdCommand)
	rootCmd.Execute()
}

// Common code to connect to the X52 joystick
func connectToX52() *x52.Context {
	ctx := x52.NewContext()

	if !ctx.Connect() {
		ctx.Close()
		fmt.Println("Unable to connect to X52 joystick. Is it plugged in?")
		os.Exit(1)
	}
	return ctx
}
