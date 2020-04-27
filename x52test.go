package main

import (
	"time"

	"nirenjan.org/saitek-x52/x52"
)

func main() {
	ctx := x52.NewContext()
	defer ctx.Close()

	ctx.Debug(4)

	ok := ctx.Connect()
	if !ok {
		return
	}

	for i := uint16(0); i <= 0x80; i++ {
		ctx.SetMFDBrightness(i)
		ctx.SetLEDBrightness(i)
		ctx.Update()
		time.Sleep(50 * time.Millisecond)
	}
}
