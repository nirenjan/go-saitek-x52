package util

import (
	"errors"
)

// Scroller allows the application to save a long string, and scroll it
// automatically for use in the X52 MFD text API
type Scroller struct {
	text    []byte
	prefix  []byte
	suffix  []byte
	flags   ScrollFlags
	textPos uint
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

	return scroller, nil
}

// Bytes returns a slice of bytes for passing to the MFD text API

// Scroll scrolls the text by a single character and returns a slice of bytes
// for passing to the MFD text API

// SetFlags updates the scroller flags

// Reset resets the scroller to the default position

// prefix + text[pos:pos+16-lp-ls] + suffix
