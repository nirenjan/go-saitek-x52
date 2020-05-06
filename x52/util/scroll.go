package util

import (
	"bytes"
	"errors"
)

// Scroller allows the application to save a long string, and scroll it
// automatically for use in the X52 MFD text API
type Scroller struct {
	text    []byte
	prefix  []byte
	suffix  []byte
	buffer  []byte
	flags   ScrollFlags
	textPos int
	start   int
	end     int
}

// ScrollFlags control the behavior of the scroller
type ScrollFlags uint

const (
	// ScrollFromOffscreen scrolls the text from off the side of the screen,
	// instead of starting with the maximum visible length
	ScrollFromOffscreen ScrollFlags = 1 << iota

	// ScrollTextOffscreen will scroll the entire text offscreen before
	// starting a new scroll cycle
	ScrollTextOffscreen

	// ScrollLeftToRight will scroll the text left-to-right instead of the
	// default right-to-left
	ScrollLeftToRight
)

// NewScroller returns a Scroller with the given parameters setup. It may
// return an error if the prefix and suffix combined will not leave sufficient
// space for the text to be displayed. The text, prefix and suffix must be in
// the code page of the MFD display.
func NewScroller(text, prefix, suffix []byte, flags ScrollFlags) (*Scroller, error) {
	scroller := &Scroller{
		text:   text,
		prefix: prefix,
		suffix: suffix,
		flags:  flags,
	}

	// Maximum length of prefix and suffix combined is 8, otherwise it won't
	// leave enough space to allow for a reasonable scroll
	if len(prefix)+len(suffix) > 8 {
		return nil, errors.New("prefix and suffix combined length too long")
	}

	scroller.SetFlags(flags)

	return scroller, nil
}

// Bytes returns a slice of bytes for passing to the MFD text API
func (sc *Scroller) Bytes() []byte {
	var data = bytes.Repeat([]byte{0x20}, 16)
	var pos int
	var tpos int

	for pos = 0; pos < len(sc.prefix); pos++ {
		data[pos] = sc.prefix[pos]
	}

	for pos = 16 - len(sc.suffix); pos < 16; pos++ {
		data[pos] = sc.suffix[pos-16+len(sc.suffix)]
	}

	for tpos, pos = 0, len(sc.prefix); pos < 16-len(sc.suffix); tpos, pos = tpos+1, pos+1 {
		if sc.textPos+tpos >= len(sc.buffer) {
			break
		}

		data[pos] = sc.buffer[sc.textPos+tpos]
	}

	return data[:]
}

// Scroll scrolls the text by a single character and returns a slice of bytes
// for passing to the MFD text API
func (sc *Scroller) Scroll() []byte {
	if sc.start != sc.end {
		if sc.start < sc.end {
			sc.textPos++
		} else {
			sc.textPos--
		}
	}

	if sc.textPos == sc.end {
		sc.textPos = sc.start
	}

	return sc.Bytes()
}

// SetFlags updates the scroller flags
func (sc *Scroller) SetFlags(flags ScrollFlags) {
	defer sc.Reset()
	// Add a buffer of 16-len(prefix)-len(suffix) spaces on both sides
	// to allow for scrolling
	buflen := 16 - len(sc.prefix) - len(sc.suffix)
	buffer := bytes.Repeat([]byte{0x20}, buflen)

	sc.buffer = sc.text
	if len(sc.text) <= buflen {
		sc.start = 0
		sc.end = 0
		return
	}

	if flags&ScrollLeftToRight == 0 {
		// Scrolling right-to-left, default
		if flags&ScrollFromOffscreen != 0 {
			sc.buffer = append(buffer, sc.buffer...)
		}
		if flags&ScrollTextOffscreen != 0 {
			sc.buffer = append(sc.buffer, buffer...)
		}

		sc.start = 0
		sc.end = len(sc.buffer) - buflen + 1
	} else {
		// Scrolling left to right
		if flags&ScrollFromOffscreen != 0 {
			sc.buffer = append(sc.buffer, buffer...)
		}
		if flags&ScrollTextOffscreen != 0 {
			sc.buffer = append(buffer, sc.buffer...)
		}
		sc.start = len(sc.buffer) - buflen
		sc.end = -1
	}
}

// Reset resets the scroller to the default position
func (sc *Scroller) Reset() {
	sc.textPos = sc.start
}
