package x52

import (
	"log"
	"time"

	"github.com/google/gousb"
)

const (
	// Line size on each of the MFDs
	mfdLineSize = 16

	// Number of lines
	mfdLines = 3

	// Number of clocks
	mfdClocks = 3
)

// usbDevice is an interface that implements the public methods of
// gousb.Device that are used by the library. This is used for testing
// the callers of those methods.
type usbDevice interface {
	Close() error
	Control(rType, request uint8, val, idx uint16, data []byte) (int, error)
	Reset() error
}

// Context manages all resources related to device handling
type Context struct {
	usbContext    *gousb.Context
	device        usbDevice
	logger        *log.Logger
	logLevel      int
	featureFlags  uint32
	updateMask    uint32
	ledMask       uint32
	mfdBrightness uint16
	ledBrightness uint16
	mfdLine       [mfdLines][]byte
	time          time.Time
	dateFormat    DateFormat
	timeFormat    [mfdClocks]ClockFormat
	timeZone      [mfdClocks]*time.Location
}

// NewContext returns a new Context against which to run device operations
func NewContext() *Context {
	ctx := new(Context)

	// Create a new usb Context
	ctx.usbContext = gousb.NewContext()

	ctx.initialize()

	return ctx
}

// initialize sets defaults in the Context
func (ctx *Context) initialize() {
	// Setup the logger
	ctx.setupLogger()

	// Reset the feature flags
	ctx.featureFlags = 0

	// Reset the masks
	ctx.updateMask = 0
	ctx.ledMask = 0

	// Reset the brightness values to 0
	ctx.mfdBrightness = 0
	ctx.ledBrightness = 0

	// Set the time to the zero time
	ctx.time = time.Time{}

	// Set the date format to DDMMYY
	ctx.dateFormat = DateFormatDDMMYY

	// Set timezone on all clocks to UTC, and time format to 12 hour
	for i := 0; i < mfdClocks; i++ {
		ctx.timeFormat[i] = ClockFormat12Hr
		ctx.timeZone[i] = time.UTC
	}
}

// Close closes the context, and any devices that may have been opened will also
// be closed
func (ctx *Context) Close() error {
	ctx.devClose()

	ctx.initialize()

	if ctx.usbContext != nil {
		defer func() { ctx.usbContext = nil }()
		return ctx.usbContext.Close()
	}

	return nil
}

func bitSet(mask *uint32, val uint32) {
	*mask |= uint32(1 << val)
}

func bitClear(mask *uint32, val uint32) {
	*mask &= ^uint32(1 << val)
}

func bitTest(mask, val uint32) bool {
	return (mask & uint32(1<<val)) != 0
}
