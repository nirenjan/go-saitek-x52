package main

// MFD tests

import (
	"bytes"
	"flag"
	"fmt"

	"nirenjan.org/saitek-x52/x52"
)

func init() {
	flag.Var(&TC{"MFD", testMFD}, "mfd", "Test multifunction display")
}

func testMFD(ctx *x52.Context) error {
	mfdFuncs := []func(*x52.Context) error{
		testMFDText,
		testMFDChars,
		testShift,
	}

	for _, f := range mfdFuncs {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func testMFDText(ctx *x52.Context) error {
	bar := progressBar("MFD Text", 256)

	for i := 0; i < 256; i += 16 {
		line1 := []byte(fmt.Sprintf("0x%02x - 0x%02x", i, i+15))
		ctx.SetMFDText(0, line1)

		line23 := bytes.Repeat([]byte{0x20}, 32)
		for j := 0; j < 16; j++ {
			line23[j<<1] = byte(i + j)
		}
		ctx.SetMFDText(1, line23[:16])
		ctx.SetMFDText(2, line23[16:])

		if err := updateDev(ctx, bar); err != nil {
			return err
		}
		delayMs(2000)
		bar.Add(16)
	}

	return nil
}

func testMFDChars(ctx *x52.Context) error {
	bar := progressBar("MFD Characters", 256)

	for i := 0; i < 256; i++ {
		line := bytes.Repeat([]byte{byte(i)}, 16)

		ctx.SetMFDText(0, line)
		ctx.SetMFDText(1, line)
		ctx.SetMFDText(2, line)

		if err := updateDev(ctx, bar); err != nil {
			return err
		}
		delayMs(500)
		bar.Add(1)
	}

	return nil
}

func testShift(ctx *x52.Context) error {
	bar := progressBar("MFD Shift On/Off", 2)

	for i := 0; i < 2; i++ {
		ctx.SetShift(i == 0)
		if err := updateDev(ctx, bar); err != nil {
			return err
		}
		delayMs(2000)
		bar.Add(1)
	}

	return nil
}
