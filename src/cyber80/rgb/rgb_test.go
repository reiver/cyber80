package cyber80_rgb

import (
	"testing"
)

func TestRGB(t *testing.T) {

	tests := []struct{
		R byte
		G byte
		B byte
		Actual   Type

		Expected Type
	}{
		{
			               R:0x00,G:0x00,B:0x00,
			Expected: Type{r:0x00,g:0x00,b:0x00},
		},



		{
			               R:0x01,G:0x00,B:0x00,
			Expected: Type{r:0x01,g:0x00,b:0x00},
		},



		{
			               R:0x00,G:0x01,B:0x00,
			Expected: Type{r:0x00,g:0x01,b:0x00},
		},



		{
			               R:0x00,G:0x00,B:0x01,
			Expected: Type{r:0x00,g:0x00,b:0x01},
		},



		{
			               R:0xab,G:0xcd,B:0xef,
			Expected: Type{r:0xab,g:0xcd,b:0xef},
		},
	}

	for testNumber, test := range tests {

		actual := RGB(test.R, test.G, test.B)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL: %#v", actual)
			continue
		}
	}
}
