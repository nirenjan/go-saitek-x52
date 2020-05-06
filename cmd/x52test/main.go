// x52test is a program that will verify the functionality of an attached
// Saitek X52/X52Pro joystick.
package main // import "nirenjan.org/saitek-x52/cmd/x52test"

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"nirenjan.org/saitek-x52/x52"
)

var selectTests bool

type TC struct {
	name    string
	handler func(*x52.Context) error
}

func (tc *TC) IsBoolFlag() bool {
	return true
}

func (tc *TC) String() string {
	return "true"
}

func (tc *TC) Set(v string) error {
	bv, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}

	if bv {
		selectTests = true
	}

	return nil
}

var mockTests bool

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "\nAll tests are run if no flags are specified")
	}

	flag.BoolVar(&mockTests, "mock", false, "Don't actually run the tests, just simulate the output")

	flag.Parse()

	// gousb has an annoying tendency to log interrupts to stderr, and these
	// mess up the progressbar display. To work around this, we force the log
	// package to send its output to ioutil.Discard
	log.SetOutput(ioutil.Discard)

	// Create an X52 context, and connect to the device
	ctx := x52.NewContext()
	defer ctx.Close()

	// Make sure that the device is opened
	if !mockTests && !ctx.Connect() {
		return
	}
	defer ctx.Reset()

	var abortTests bool
	runTests := func(f *flag.Flag) {
		if abortTests {
			return
		}
		value, ok := f.Value.(*TC)
		if !ok {
			return
		}
		fmt.Printf("Running %s tests\n", value.name)
		if err := value.handler(ctx); err != nil {
			fmt.Printf("%s tests failed: %v\n", value.name, err)
			fmt.Println("Aborting any remaining tests")
			abortTests = true
		}
	}

	if selectTests {
		flag.Visit(runTests)
	} else {
		flag.VisitAll(runTests)
	}
}
