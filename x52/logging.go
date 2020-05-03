package x52

import (
	"fmt"
	"log"
	"os"
)

// The logging levels correspond to the same levels in libusb
const (
	logNone = iota
	logError
	logWarning
	logInfo
	logDebug
)

// Debug changes the debug level. Level 0 means no debug, higher levels will
// print out more debugging information.
func (ctx *Context) Debug(level int) {
	if level < logNone {
		level = logNone
	} else if level > logDebug {
		level = logDebug
	}

	ctx.logLevel = level

	// If the log level is set to debug, then add the file info to the flags
	var flags = log.LstdFlags | log.Lmsgprefix
	if level == logDebug {
		flags |= log.Lshortfile
	}
	ctx.logger.SetFlags(flags)
}

// DebugUSB changes the debug level of the USB subsystem. Level 0 means no
// debug, higher levels will print out more debugging information.
func (ctx *Context) DebugUSB(level int) {
	ctx.usbContext.Debug(level)
}

func (ctx *Context) setupLogger() {
	ctx.logger = log.New(os.Stderr, "x52: ", log.LstdFlags)
	ctx.logLevel = logNone
}

func (ctx *Context) log(level int, v ...interface{}) {
	if level <= ctx.logLevel {
		ctx.logger.Output(2, fmt.Sprint(v...))
	}
}

func (ctx *Context) logf(level int, format string, v ...interface{}) {
	if level <= ctx.logLevel {
		ctx.logger.Output(2, fmt.Sprintf(format, v...))
	}
}
