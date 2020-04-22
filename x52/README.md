Saitek X52/X52Pro Library
=========================

The Saitek X52/X52Pro library is used to interface with a supported Saitek
X52/X52Pro joystick. The library uses the [gousb] library to communicate with
the joystick.

# API Overview

In order to use the library, you will need to create a context. All joystick LED
and MFD data is managed through the context. You should close the context when
you no longer need it.

```go
ctx := x52.NewContext()
defer ctx.Close()
```

Unlike the corresponding [C library], context creation does not require you to
have the joystick plugged in until you wish to actually update it. Therefore,
you will need to call `OpenDevice` to actually connect to the device. The
`OpenDevice` will return a `nil` error if it was successful in connecting to the
joystick, and a corresponding `error` value if it failed for any reason.

```go
err := ctx.OpenDevice()
if err != nil {
    // ...
}
```

The library also provides a corresponding `CloseDevice` method, which can be
used when the application no longer needs access to the joystick. However, the
library will automatically call `CloseDevice` when the context is closed.

All the `Set*` methods set internal data structures within the context, and
don't actually update the joystick. In order to do so, the application must call
the `Update` method of the context.

```go
err = ctx.SetLed(x52.LedFire, x52.LedOff)
if err != nil {
    // ...
}
// Write the updated state to the joystick
err = ctx.Update()

```

The library will also check and close the device automatically if updating it
fails because the joystick was unplugged.

# LED and MFD control

Currently, the library supports setting the state of all LEDs, the brightness of
the LEDs and MFD, and the text on the MFD. The update API is still not working,
so the sets won't take effect anyway. The API for these can be considered
frozen. The library does NOT yet support setting the date and time display on
the MFD, and the API for this is still a work in progress.

# Limitations

The library can maintain a connection to only 1 supported device at a time. This
means that if you, for some reason, have multiple X52 joysticks connected, only
1 of them will be controlled by the library. No guarantees are made about which
of these will be selected, as this is a function of the order in which the
devices are enumerated by the USB subsystem.



[gousb]: https://github.com/google/gousb
[C library]: https://github.com/nirenjan/x52pro-linux
