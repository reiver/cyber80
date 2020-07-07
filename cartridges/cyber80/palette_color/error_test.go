package cyber80_palette_color

import (
	"testing"
)

func TestError(t *testing.T) {

	tests := []struct{
		Actual Type
		Expected  Type
	}{
		{
			Actual: Error(),
			Expected: Type{
				loaded: true,
				errored: true,
			},
		},
	}

	for testNumber, test := range tests {

		if expected, actual := test.Expected, test.Actual; expected != actual {
			t.Errorf("For test #%d, what was expected was not what was actually gotten.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}

	}
}
