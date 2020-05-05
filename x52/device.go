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

	// Reset any flags that may have been set
	ctx.featureFlags = 0
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

	// We have at least 1 device, use the first in the list, and close
	// all the others
	for i, dev := range devlist {
		if i == 0 {
			// Pick the first device
			ctx.logf(logInfo, "Picking device on bus %v, address %v, port %v",
				dev.Desc.Bus, dev.Desc.Address, dev.Desc.Port)

			ctx.device = dev

			// Set flags based on the device
			if isPro := (dev.Desc.Product == productX52Pro); isPro {
				bitSet(&ctx.featureFlags, featureLed)
			}
		} else {
			// Close the remaining devices
			ctx.logf(logInfo, "Closing device on bus %v, address %v, port %v",
				dev.Desc.Bus, dev.Desc.Address, dev.Desc.Port)
			defer dev.Close()
		}
	}
	return true
}

func (ctx *Context) checkDisconnect(action string, err error) error {
	if err != nil {
		ctx.logf(logError, "error %v device: %v", action, err)

		if err == gousb.ErrorNoDevice {
			// Device has been unplugged, close it
			ctx.devClose()
			ctx.log(logWarning, "device has been disconnected")

			return errNotConnected(err)
		}

		return err
	}

	return nil
}

// Reset resets the connected device. If there is no device connected, it
// returns errNotConnected, otherwise it will return a corresponding USB
// error
func (ctx *Context) Reset() error {
	if ctx.device == nil {
		ctx.log(logWarning, "not connected")
		return errNotConnected(nil)
	}

	ctx.log(logDebug, "resetting device")
	return ctx.checkDisconnect("resetting", ctx.device.Reset())
}

// Raw sends a raw vendor control packet to the device
func (ctx *Context) Raw(index, value uint16) error {
	if ctx.device == nil {
		ctx.log(logWarning, "not connected")
		return errNotConnected(nil)
	}

	// gousb takes care of retries internally, so we don't have to
	// do it ourselves
	ctx.logf(logDebug, "sending raw %04x %04x", index, value)
	_, err := ctx.device.Control(
		gousb.ControlVendor|gousb.ControlDevice|gousb.ControlOut,
		0x91, value, index, nil)

	return ctx.checkDisconnect("updating", err)
}
