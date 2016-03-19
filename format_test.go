package nanpa

import "testing"

type testData struct {
	Input  string
	Output string
	Error  error
}

var tests = []struct {
	Input  string
	Output string
	Error  string
}{
	{"2342355678", "12342355678", ""},
	{"2349115678", "12349115678", ""}, // NOTE: wikipedia specifies this is invalid
	{"2812345678", "12812345678", ""},

	{"0", "0", "Unhandled number length"},
	{"3141592653", "3141592653", "Non-NANPA number"},
	{"1232345678", "1232345678", "Non-NANPA number"},
	{"32812345678", "32812345678", "Unhandled number length"},
}

func TestFormat(t *testing.T) {
	for _, d := range tests {
		out, err := Format(d.Input)

		errS := ""
		if err != nil {
			errS = err.Error()
		}

		if errS != d.Error || out != d.Output {
			t.Errorf("Expected nanpa.Format(%s) to return (%s, %s), returns (%s, %s)",
				d.Input, d.Output, d.Error,
				out, errS)
		}
	}
}
