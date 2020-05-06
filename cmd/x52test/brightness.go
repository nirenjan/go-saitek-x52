package main

// Brightness tests

import (
	"flag"
	"time"

	"nirenjan.org/saitek-x52/x52"
)

func init() {
	flag.Var(&TC{"brightness scale", testBrightness}, "brightness", "Test brightness scale")
}

func testBrightness(ctx *x52.Context) error {
	bar := progressBar("MFD Brightness", 0x81)

	for i := 0; i < 0x81; i++ {
		ctx.SetMFDBrightness(uint16(i))
		err := ctx.Update()
		if err != nil {
			bar.Clear()
			return err
		}
		time.Sleep(100 * time.Millisecond)
		bar.Add(1)
	}

	bar = progressBar("LED Brightness", 0x81)
	for i := 0; i < 0x81; i++ {
		ctx.SetLEDBrightness(uint16(i))
		err := ctx.Update()
		if err != nil {
			bar.Clear()
			return err
		}
		time.Sleep(100 * time.Millisecond)
		bar.Add(1)
	}

	return nil
}
