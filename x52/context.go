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

// Context manages all resources related to device handling
type Context struct {
	usbContext    *gousb.Context
	device        *gousb.Device
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

	// Set timezone on all clocks to UTC
	for i := 0; i < mfdClocks; i++ {
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
