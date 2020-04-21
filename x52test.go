package main

import (
	"time"

	"github.com/google/gousb"
)

func vendor(dev *gousb.Device, val, idx uint16) error {
	_, err := dev.Control(gousb.ControlVendor|gousb.ControlDevice|gousb.ControlOut,
		0x91, val, idx, nil)
	return err
}

func main() {
	ctx := gousb.NewContext()
	defer ctx.Close()

	ctx.Debug(4)

	const vid = gousb.ID(0x06a3)
	const pid = gousb.ID(0x0762)
	dev, err := ctx.OpenDeviceWithVIDPID(vid, pid)

	if err != nil {
		panic(err)
	}
	defer dev.Close()

	for i := uint16(0); i <= 0x80; i++ {
		vendor(dev, i, 0xb1)
		vendor(dev, i, 0xb2)

		time.Sleep(50 * time.Millisecond)
	}
}
