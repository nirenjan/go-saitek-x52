Saitek X52/X52Pro joystick driver for Unix/Linux
================================================

This project adds a new driver for the Saitek/MadCatz X52/X52Pro flight
control system. The X52/X52Pro is a HOTAS (hand on throttle and stick)
with 7 axes, 39 buttons, 1 hat and 1 thumbstick and a multi-function
display which is programmable.

Currently, only Windows drivers are available from Saitek, which led me
to develop a new Linux driver which can program the MFD and the individual
LEDs on the joystick. The standard usbhid driver is capable of reading
the joystick, but it cannot control the MFD or LEDs.

Most of the extra functionality can be handled from userspace. See the
individual folders for README information.

This project is a rewrite of https://github.com/nirenjan/x52pro-linux in Go.
