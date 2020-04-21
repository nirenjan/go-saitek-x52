// Package x52 provides handling for the X52/X52Pro devices
package x52 // import "nirenjan.org/saitek-x52/x52"

// ClockID identifies the clock on the X52 multifunction display
type ClockID uint

const (
	Clock1 ClockID = iota
	Clock2
	Clock3
)

// ClockFormat describes the time format on the MFD
type ClockFormat uint

const (
	ClockFormat12Hr ClockFormat = iota
	ClockFormat24Hr
)

// DateFormat describes the date format on the MFD
type DateFormat uint

const (
	DateFormatDDMMYY DateFormat = iota
	DateFormatMMDDYY
	DateFormatYYMMDD
)

// LED contains the supported LEDs
type LED uint

const (
	LedFire     = LED(0x01)
	LedA        = LED(0x02)
	LedB        = LED(0x04)
	LedD        = LED(0x06)
	LedE        = LED(0x08)
	LedT1       = LED(0x0a)
	LedT2       = LED(0x0c)
	LedT3       = LED(0x0e)
	LedPOV      = LED(0x10)
	LedClutch   = LED(0x12)
	LedThrottle = LED(0x14)
)

// LedState contains the possible LED states. Each LED only supports a subset
// of the states
type LedState uint

const (
	LedOff LedState = iota
	LedOn
	LedRed
	LedAmber
	LedGreen
)

// Feature flags
const (
	featureLed uint32 = 1 << iota
)

// Update bits
const (
	updateShift uint32 = iota
	updateLedFire
	updateLedARed
	updateLedAGreen
	updateLedBRed
	updateLedBGreen
	updateLedDRed
	updateLedDGreen
	updateLedERed
	updateLedEGreen
	updateLedT1Red
	updateLedT1Green
	updateLedT2Red
	updateLedT2Green
	updateLedT3Red
	updateLedT3Green
	updateLedPOVRed
	updateLedPOVGreen
	updateLedClutchRed
	updateLedClutchGreen
	updateLedThrottle
	updateMfdLine1
	updateMfdLine2
	updateMfdLine3
	updatePOVBlink
	updateBrightnessMFD
	updateBrightnessLED
	updateDate
	updateTime
	updateOffs1
	updateOffs2
	updateMax
)
