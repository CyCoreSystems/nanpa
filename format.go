package nanpa

import (
	"fmt"
	"regexp"
	"strings"
)

// Format makes certain the destination number is 11 digits, starting with a 1
// and conforming to NANPA standards; if not, it returns the original string, unmodified
// and an error
func Format(input string) (string, error) {

	// Strip out anything other than digits and '+'
	stripper := regexp.MustCompile("[^+0-9]")
	output := stripper.ReplaceAllString(input, "")

	// If we have a +1 prefix, strip it, too
	output = strings.TrimPrefix(output, "+1")

	// If we have a 011-1 prefix, strip it, too
	output = strings.TrimPrefix(output, "0111")

	// Check for 10 digit length
	if len(output) != 10 {
		// If the length isn't 11, we can't do anything with it,
		// so just return the original string
		if len(output) != 11 {
			return input, fmt.Errorf("Unhandled number length")
		}
		// If the length is 11 but begins with something other
		// than "1", return original string
		if !strings.HasPrefix(output, "1") {
			return input, fmt.Errorf("Unhandled number length")
		}
	} else {
		// Length is 10, so prepend a 1
		output = "1" + output
	}

	// Verify 1NXXNXXXXXX format
	if output[1] == uint8('0') || output[1] == uint8('1') ||
		output[4] == uint8('0') || output[4] == uint8('1') {
		return input, fmt.Errorf("Non-NANPA number")
	}

	// Otherwise, return out newly-formatted number
	return output, nil
}
