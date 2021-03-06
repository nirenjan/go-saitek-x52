package util

// ConvertStringToX52Charmap converts a string to a byte slice that is accepted
// by the X52 MFD display. If replace is true, then runes that are not
// recognized by this function are replaced with the replacement byte, otherwise
// they are dropped from the output
func ConvertStringToX52Charmap(s string, replace bool, c byte) []byte {
	out := make([]byte, 0, len(s))
	for _, r := range s {
		ch, ok := charmap[r]
		if !ok && replace {
			ok = replace
			ch = c
		}
		if ok {
			out = append(out, ch)
		}
	}

	return out
}

// ReplaceMissing is the default replacement character for unsupported Unicode
// code points
const ReplaceMissing byte = 0xDB

// Conversion Map for X52 Pro MFD character map

// The X52 Pro MFD uses a single byte character set and encodes multiple
// character ranges in that set. This file defines the mapping from Unicode
// code points to the vendor character set. This is transformed at compilation
// time to a lookup table using UTF-8. All characters must be explicitly
// specified to be added to the lookup table.

// Lines must be formatted as follows
// <Unicode Code point in hex>: <MFD Charmap value in hex>,
// or
// <Unicode Code point in hex> <MFD Charmap value as single character>
// Comment lines begin with the text //
// Comments may begin after the character map value, they are ignored as long
// as the rest of the line is in a valid format.

// Code points which are not found in the list below will translate to
// 0xDB which is the entry in the character map for a box (similar to U+25A1)

// Note that the library will not attempt to perform any additional matching
// steps like iconv does to find a close match in the glyph, so if you need
// to add any such "close matches", you will need to explicitly list them
// in the list below.

// The following chart indicates which glyphs in the character map have been
// listed below, it is intended to give at a quick glance what additional
// characters need to be mapped. Rows are the lower nibble, columns are the
// higher nibble.
// NOTE TO MAINTAINERS: Update the chart below when you add new codepoints to
// the translation list.

// LEGEND
// ======
// ? - Needs to be mapped
// * - Already mapped
// x - There is no mapping for this (0x00-0x07)

// +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
// | \ | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | A | B | C | D | E | F |
// +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
// | 0 | x | * | * | * | * | * | * | * | * | * | * | * | * | * | * | ? |
// | 1 | x | * | * | * | * | * | * | * | * | * | * | * | * | * | * | ? |
// | 2 | x | * | * | * | * | * | * | * | * | * | * | * | * | * | * | ? |
// | 3 | x | * | * | * | * | * | * | * | * | * | * | * | * | * | * | ? |
// | 4 | x | * | * | * | * | * | * | * | * | * | * | * | * | * | * | ? |
// | 5 | x | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// | 6 | x | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// | 7 | x | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// | 8 | ? | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// | 9 | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// | A | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// | B | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// | C | * | * | * | * | * | ? | * | * | * | * | * | * | * | * | * | * |
// | D | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// | E | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// | F | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * | * |
// +---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+

// Map starts below
// ----------------------------------------------------------------------------
var charmap = map[rune]byte{
	// Printable ASCII Range - These map directly onto their respective code
	// points except where specified.
	0x0020: 0x20,
	0x0021: 0x21,
	0x0022: 0x22,
	0x0023: 0x23,
	0x0024: 0x24,
	0x0025: 0x25,
	0x0026: 0x26,
	0x0027: 0x27,
	0x0028: 0x28,
	0x0029: 0x29,
	0x002a: 0x2a,
	0x002b: 0x2b,
	0x002c: 0x2c,
	0x002d: 0x2d,
	0x002e: 0x2e,
	0x002f: 0x2f,
	0x0030: 0x30,
	0x0031: 0x31,
	0x0032: 0x32,
	0x0033: 0x33,
	0x0034: 0x34,
	0x0035: 0x35,
	0x0036: 0x36,
	0x0037: 0x37,
	0x0038: 0x38,
	0x0039: 0x39,
	0x003a: 0x3a,
	0x003b: 0x3b,
	0x003c: 0x3c,
	0x003d: 0x3d,
	0x003e: 0x3e,
	0x003f: 0x3f,
	0x0040: 0x40,
	0x0041: 0x41,
	0x0042: 0x42,
	0x0043: 0x43,
	0x0044: 0x44,
	0x0045: 0x45,
	0x0046: 0x46,
	0x0047: 0x47,
	0x0048: 0x48,
	0x0049: 0x49,
	0x004a: 0x4a,
	0x004b: 0x4b,
	0x004c: 0x4c,
	0x004d: 0x4d,
	0x004e: 0x4e,
	0x004f: 0x4f,
	0x0050: 0x50,
	0x0051: 0x51,
	0x0052: 0x52,
	0x0053: 0x53,
	0x0054: 0x54,
	0x0055: 0x55,
	0x0056: 0x56,
	0x0057: 0x57,
	0x0058: 0x58,
	0x0059: 0x59,
	0x005a: 0x5a,
	0x005b: 0x5b,
	// Backslash (\) does not appear in the character set
	0x005d: 0x5d,
	0x005e: 0x5e,
	0x005f: 0x5f,
	0x0060: 0x60,
	0x0061: 0x61,
	0x0062: 0x62,
	0x0063: 0x63,
	0x0064: 0x64,
	0x0065: 0x65,
	0x0066: 0x66,
	0x0067: 0x67,
	0x0068: 0x68,
	0x0069: 0x69,
	0x006a: 0x6a,
	0x006b: 0x6b,
	0x006c: 0x6c,
	0x006d: 0x6d,
	0x006e: 0x6e,
	0x006f: 0x6f,
	0x0070: 0x70,
	0x0071: 0x71,
	0x0072: 0x72,
	0x0073: 0x73,
	0x0074: 0x74,
	0x0075: 0x75,
	0x0076: 0x76,
	0x0077: 0x77,
	0x0078: 0x78,
	0x0079: 0x79,
	0x007a: 0x7a,
	0x007b: 0x7b,
	0x007c: 0x7c,
	0x007d: 0x7d,
	// Tilde (~) does not appear in the character set

	// Miscellaneous Symbols
	0x00B7: 0x0D, // MIDDLE DOT
	0x00AE: 0x0E, // REGISTERED SIGN
	0x00A9: 0x0F, // COPYRIGHT SIGN
	0x2122: 0x10, // TRADE MARK SIGN
	0x2020: 0x11, // DAGGER
	0x00A7: 0x12, // SECTION SIGN
	0x00B6: 0x13, // PILCROW SIGN
	0x2192: 0x7E, // RIGHTWARDS ARROW
	0x2190: 0x7F, // LEFTWARDS ARROW (also available at 0x08)
	0x00A0: 0xA0, // NO-BREAK SPACE
	0x203E: 0xFF, // OVERLINE

	// Mathematical Symbols
	0x00BD: 0xF5, // VULGAR FRACTION ONE HALF
	0x00BC: 0xF6, // VULGAR FRACTION ONE QUARTER
	0x00D7: 0xF7, // MULTIPLICATION SIGN
	0x00F7: 0xF8, // DIVISION SIGN
	0x2264: 0xF9, // LESS-THAN OR EQUAL TO
	0x2265: 0xFA, // GREATER-THAN OR EQUAL TO
	0x226A: 0xFB, // MUCH LESS-THAN
	0x226B: 0xFC, // MUCH GREATER-THAN
	0x2260: 0xFD, // NOT EQUAL TO
	0x221A: 0xFE, // SQUARE ROOT

	// Accented Latin characters
	0x00C7: 0x80, // LATIN CAPITAL LETTER C WITH CEDILLA
	0x00FC: 0x81, // LATIN SMALL LETTER U WITH DIAERESIS
	0x00E9: 0x82, // LATIN SMALL LETTER E WITH ACUTE
	0x00E2: 0x83, // LATIN SMALL LETTER A WITH CIRCUMFLEX
	0x00E4: 0x84, // LATIN SMALL LETTER A WITH DIAERESIS
	0x00E0: 0x85, // LATIN SMALL LETTER A WITH GRAVE
	0x0227: 0x86, // LATIN SMALL LETTER A WITH DOT ABOVE
	0x00E7: 0x87, // LATIN SMALL LETTER C WITH CEDILLA
	0x00EA: 0x88, // LATIN SMALL LETTER E WITH CIRCUMFLEX
	0x00EB: 0x89, // LATIN SMALL LETTER E WITH DIAERESIS
	0x00E8: 0x8A, // LATIN SMALL LETTER E WITH GRAVE
	0x00EF: 0x8B, // LATIN SMALL LETTER I WITH DIAERESIS
	0x00EE: 0x8C, // LATIN SMALL LETTER I WITH CIRCUMFLEX
	0x00EC: 0x8D, // LATIN SMALL LETTER I WITH GRAVE
	0x00C4: 0x8E, // LATIN CAPITAL LETTER A WITH DIAERESIS
	0x00C2: 0x8F, // LATIN CAPITAL LETTER A WITH CIRCUMFLEX

	0x00C9: 0x90, // LATIN CAPITAL LETTER E WITH ACUTE
	0x00E6: 0x91, // LATIN SMALL LETTER AE
	0x00C6: 0x92, // LATIN CAPITAL LETTER AE
	0x00F4: 0x93, // LATIN SMALL LETTER O WITH CIRCUMFLEX
	0x00F6: 0x94, // LATIN SMALL LETTER O WITH DIAERESIS
	0x00F2: 0x95, // LATIN SMALL LETTER O WITH GRAVE
	0x00FB: 0x96, // LATIN SMALL LETTER U WITH CIRCUMFLEX
	0x00F9: 0x97, // LATIN SMALL LETTER U WITH GRAVE
	0x00FF: 0x98, // LATIN SMALL LETTER Y WITH DIAERESIS
	0x00D6: 0x99, // LATIN CAPITAL LETTER O WITH DIAERESIS
	0x00DC: 0x9A, // LATIN CAPITAL LETTER U WITH DIAERESIS
	0x00F1: 0x9B, // LATIN SMALL LETTER N WITH TILDE
	0x00D1: 0x9C, // LATIN CAPITAL LETTER N WITH TILDE
	0x00AA: 0x9D, // FEMININE ORDINAL INDICATOR
	0x00BA: 0x9E, // MASCULINE ORDINAL INDICATOR
	0x00BF: 0x9F, // INVERTED QUESTION MARK

	0x00E1: 0xE0, // LATIN SMALL LETTER A WITH ACUTE
	0x00ED: 0xE1, // LATIN SMALL LETTER I WITH ACUTE
	0x00F3: 0xE2, // LATIN SMALL LETTER O WITH ACUTE
	0x00FA: 0xE3, // LATIN SMALL LETTER U WITH ACUTE
	0x00A2: 0xE4, // CENT SIGN
	0x00A3: 0xE5, // POUND SIGN
	0x00A5: 0xE6, // YEN SIGN
	0x20A7: 0xE7, // PESETA SIGN
	0x0192: 0xE8, // LATIN SMALL LETTER F WITH HOOK
	0x00A1: 0xE9, // INVERTED EXCLAMATION MARK
	0x00C3: 0xEA, // LATIN CAPITAL LETTER A WITH TILDE
	0x00E3: 0xEB, // LATIN SMALL LETTER A WITH TILDE
	0x00D5: 0xEC, // LATIN CAPITAL LETTER O WITH TILDE
	0x00F5: 0xED, // LATIN SMALL LETTER O WITH TILDE
	0x00D8: 0xEE, // LATIN CAPITAL LETTER O WITH STROKE
	0x00F8: 0xEF, // LATIN SMALL LETTER O WITH STROKE

	// Greek
	0x0393: 0x14, // GREEK CAPITAL LETTER GAMMA
	0x0394: 0x15, // GREEK CAPITAL LETTER DELTA
	0x0398: 0x16, // GREEK CAPITAL LETTER THETA
	0x039B: 0x17, // GREEK CAPITAL LETTER LAMDA
	0x039E: 0x18, // GREEK CAPITAL LETTER XI
	0x03A0: 0x19, // GREEK CAPITAL LETTER PI
	0x03A3: 0x1A, // GREEK CAPITAL LETTER SIGMA
	0x03D2: 0x1B, // GREEK UPSILON WITH HOOK SYMBOL
	0x03A6: 0x1C, // GREEK CAPITAL LETTER PHI
	0x03A8: 0x1D, // GREEK CAPITAL LETTER PSI
	0x03A9: 0x1E, // GREEK CAPITAL LETTER OMEGA
	0x03B1: 0x1F, // GREEK SMALL LETTER ALPHA

	// Box Drawing
	0x250C: 0x09, // BOX DRAWINGS LIGHT DOWN AND RIGHT
	0x2510: 0x0A, // BOX DRAWINGS LIGHT DOWN AND LEFT
	0x2514: 0x0B, // BOX DRAWINGS LIGHT UP AND RIGHT
	0x2518: 0x0C, // BOX DRAWINGS LIGHT UP AND LEFT

	// Halfwidth CJK punctuation
	0xFF61: 0xA1, // HALFWIDTH IDEOGRAPHIC FULL STOP
	0xFF62: 0xA2, // HALFWIDTH LEFT CORNER BRACKET
	0xFF63: 0xA3, // HALFWIDTH RIGHT CORNER BRACKET
	0xFF64: 0xA4, // HALFWIDTH IDEOGRAPHIC COMMA

	// Halfwidth Katakana variants
	0xFF65: 0xA5, // HALFWIDTH KATAKANA MIDDLE DOT
	0xFF66: 0xA6, // HALFWIDTH KATAKANA LETTER WO
	0xFF67: 0xA7, // HALFWIDTH KATAKANA LETTER SMALL A
	0xFF68: 0xA8, // HALFWIDTH KATAKANA LETTER SMALL I
	0xFF69: 0xA9, // HALFWIDTH KATAKANA LETTER SMALL U
	0xFF6A: 0xAA, // HALFWIDTH KATAKANA LETTER SMALL E
	0xFF6B: 0xAB, // HALFWIDTH KATAKANA LETTER SMALL O
	0xFF6C: 0xAC, // HALFWIDTH KATAKANA LETTER SMALL YA
	0xFF6D: 0xAD, // HALFWIDTH KATAKANA LETTER SMALL YU
	0xFF6E: 0xAE, // HALFWIDTH KATAKANA LETTER SMALL YO
	0xFF6F: 0xAF, // HALFWIDTH KATAKANA LETTER SMALL TU
	0xFF70: 0xB0, // HALFWIDTH KATAKANA-HIRAGANA PROLONGED SOUND MARK
	0xFF71: 0xB1, // HALFWIDTH KATAKANA LETTER A
	0xFF72: 0xB2, // HALFWIDTH KATAKANA LETTER I
	0xFF73: 0xB3, // HALFWIDTH KATAKANA LETTER U
	0xFF74: 0xB4, // HALFWIDTH KATAKANA LETTER E
	0xFF75: 0xB5, // HALFWIDTH KATAKANA LETTER O
	0xFF76: 0xB6, // HALFWIDTH KATAKANA LETTER KA
	0xFF77: 0xB7, // HALFWIDTH KATAKANA LETTER KI
	0xFF78: 0xB8, // HALFWIDTH KATAKANA LETTER KU
	0xFF79: 0xB9, // HALFWIDTH KATAKANA LETTER KE
	0xFF7A: 0xBA, // HALFWIDTH KATAKANA LETTER KO
	0xFF7B: 0xBB, // HALFWIDTH KATAKANA LETTER SA
	0xFF7C: 0xBC, // HALFWIDTH KATAKANA LETTER SI
	0xFF7D: 0xBD, // HALFWIDTH KATAKANA LETTER SU
	0xFF7E: 0xBE, // HALFWIDTH KATAKANA LETTER SE
	0xFF7F: 0xBF, // HALFWIDTH KATAKANA LETTER SO
	0xFF80: 0xC0, // HALFWIDTH KATAKANA LETTER TA
	0xFF81: 0xC1, // HALFWIDTH KATAKANA LETTER TI
	0xFF82: 0xC2, // HALFWIDTH KATAKANA LETTER TU
	0xFF83: 0xC3, // HALFWIDTH KATAKANA LETTER TE
	0xFF84: 0xC4, // HALFWIDTH KATAKANA LETTER TO
	0xFF85: 0xC5, // HALFWIDTH KATAKANA LETTER NA
	0xFF86: 0xC6, // HALFWIDTH KATAKANA LETTER NI
	0xFF87: 0xC7, // HALFWIDTH KATAKANA LETTER NU
	0xFF88: 0xC8, // HALFWIDTH KATAKANA LETTER NE
	0xFF89: 0xC9, // HALFWIDTH KATAKANA LETTER NO
	0xFF8A: 0xCA, // HALFWIDTH KATAKANA LETTER HA
	0xFF8B: 0xCB, // HALFWIDTH KATAKANA LETTER HI
	0xFF8C: 0xCC, // HALFWIDTH KATAKANA LETTER HU
	0xFF8D: 0xCD, // HALFWIDTH KATAKANA LETTER HE
	0xFF8E: 0xCE, // HALFWIDTH KATAKANA LETTER HO
	0xFF8F: 0xCF, // HALFWIDTH KATAKANA LETTER MA
	0xFF90: 0xD0, // HALFWIDTH KATAKANA LETTER MI
	0xFF91: 0xD1, // HALFWIDTH KATAKANA LETTER MU
	0xFF92: 0xD2, // HALFWIDTH KATAKANA LETTER ME
	0xFF93: 0xD3, // HALFWIDTH KATAKANA LETTER MO
	0xFF94: 0xD4, // HALFWIDTH KATAKANA LETTER YA
	0xFF95: 0xD5, // HALFWIDTH KATAKANA LETTER YU
	0xFF96: 0xD6, // HALFWIDTH KATAKANA LETTER YO
	0xFF97: 0xD7, // HALFWIDTH KATAKANA LETTER RA
	0xFF98: 0xD8, // HALFWIDTH KATAKANA LETTER RI
	0xFF99: 0xD9, // HALFWIDTH KATAKANA LETTER RU
	0xFF9A: 0xDA, // HALFWIDTH KATAKANA LETTER RE
	0xFF9B: 0xDB, // HALFWIDTH KATAKANA LETTER RO
	0xFF9C: 0xDC, // HALFWIDTH KATAKANA LETTER WA
	0xFF9D: 0xDD, // HALFWIDTH KATAKANA LETTER N
	0xFF9E: 0xDE, // HALFWIDTH KATAKANA VOICED SOUND MARK
	0xFF9F: 0xDF, // HALFWIDTH KATAKANA SEMI-VOICED SOUND MARK
}
