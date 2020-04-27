package x52

import (
	"github.com/google/gousb"
)

// devClose closes the device
func (ctx *Context) devClose() {
	if ctx.device != nil {
		ctx.device.Close()
		ctx.device = nil
	}
}

const (
	vendorSaitek = gousb.ID(0x06a3)

	productX52_1  = gousb.ID(0x0255)
	productX52_2  = gousb.ID(0x075c)
	productX52Pro = gousb.ID(0x0762)
)

// devSupported returns a boolean if the device is a supported one
func devSupported(desc *gousb.DeviceDesc) bool {
	if desc.Vendor != vendorSaitek {
		return false
	}

	switch desc.Product {
	case productX52_1, productX52_2, productX52Pro:
		return true
	}

	return false
}

// devSetFlags sets the feature flags based on the product ID
func (ctx *Context) devSetFlags() {
	if ctx.device == nil {
		return
	}

	if isPro := (ctx.device.Desc.Product == productX52Pro); isPro {
		bitSet(&ctx.featureFlags, featureLed)
	}
}

// Connect will try to connect to a supported X52/X52Pro joystick. If the
// joystick is plugged in and the function succeeds, it returns true, otherwise
// it returns false. If multiple supported devices are plugged in, then it will
// pick one of the supported devices in an unspecified manner.
func (ctx *Context) Connect() bool {
	devlist, err := ctx.usbContext.OpenDevices(devSupported)

	if err != nil {
		ctx.logf(logError, "error opening devices: %v", err)
		// Close any opened devices
		for _, dev := range devlist {
			defer dev.Close()
		}

		return false
	}

	// No matching device
	if len(devlist) == 0 {
		ctx.log(logInfo, "no matching devices found")
		return false
	}

	// We have at least 1 device, use the first in the list
	ctx.device = devlist[0]

	// More than 1 matching device
	if len(devlist) > 1 {
		ctx.log(logInfo, "found multiple matching devices")
		// Close all but the first
		for i, dev := range devlist {
			if i == 0 {
				// Pick the first device
				ctx.logf(logInfo, "Picking device on bus %v, address %v, port %v",
					dev.Desc.Bus, dev.Desc.Address, dev.Desc.Port)
			} else {
				// Close the remaining devices
				ctx.logf(logInfo, "Closing device on bus %v, address %v, port %v",
					dev.Desc.Bus, dev.Desc.Address, dev.Desc.Port)
				defer dev.Close()
			}
		}
	}
	return true
}
