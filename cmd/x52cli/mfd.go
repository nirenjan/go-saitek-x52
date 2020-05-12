package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"nirenjan.org/saitek-x52/x52/util"
)

var mfdCommand *cobra.Command

var mfdTextReplace bool
var mfdTextReplChar byte

func init() {
	mfdCommand = &cobra.Command{
		Use:   "mfd LINE TEXT",
		Short: "Set the text on the MFD",
		Long: `Set the text on the given MFD line.

LINE must be 1, 2 or 3. TEXT can be any string, but it must be
specified explicitly, even if empty. On most shells, an empty
string can be specified as '' or "".

The given string is translated into the character map of the X52
display. The options --replace and --replacement-byte control
whether characters that are not recognized by the display are shown
or not, and if shown, which byte in the character map to use.

If the translated string exceeds the line length, then it is
silently truncated.

`,
		Args: cobra.ExactArgs(2),
		RunE: setMFDText,
	}

	mfdCommand.Flags().BoolVar(&mfdTextReplace, "replace", false, "replace unknown characters")
	mfdCommand.Flags().Uint8Var(&mfdTextReplChar, "replacement-byte", util.ReplaceMissing, "replacement byte")
}

func setMFDText(_ *cobra.Command, args []string) error {
	// Get line index
	line, err := strconv.Atoi(args[0])
	if err != nil {
		e := err.(*strconv.NumError)
		return fmt.Errorf("parsing %q: %v", e.Num, e.Err)
	}

	// x52cli uses 1 based indexing
	if line < 1 || line > 3 {
		return fmt.Errorf("Line %v is outside the range [1, 3]", line)
	}

	ctx := connectToX52()
	defer ctx.Close()

	if cliVerbose {
		fmt.Printf("Setting MFD line %v to %q\n", line, args[1])
	}

	data := util.ConvertStringToX52Charmap(args[1], mfdTextReplace, mfdTextReplChar)
	ctx.SetMFDText(uint8(line), data)
	ctx.Update()

	return nil
}
