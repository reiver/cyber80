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
			Expected: Type{value:[3]byte{0x00,  0x00,  0x00}},
		},



		{
			                           R:0x01,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x01,  0x00,  0x00}},
		},
		{
			                           R:0x02,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x02,  0x00,  0x00}},
		},
		{
			                           R:0x03,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x03,  0x00,  0x00}},
		},
		{
			                           R:0x04,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x04,  0x00,  0x00}},
		},
		{
			                           R:0x05,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x05,  0x00,  0x00}},
		},
		{
			                           R:0x06,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x06,  0x00,  0x00}},
		},
		{
			                           R:0x07,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x07,  0x00,  0x00}},
		},
		{
			                           R:0x08,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x08,  0x00,  0x00}},
		},
		{
			                           R:0x09,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x09,  0x00,  0x00}},
		},
		{
			                           R:0x0a,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x0a,  0x00,  0x00}},
		},
		{
			                           R:0x0b,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x0b,  0x00,  0x00}},
		},
		{
			                           R:0x0c,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x0c,  0x00,  0x00}},
		},
		{
			                           R:0x0d,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x0d,  0x00,  0x00}},
		},
		{
			                           R:0x0e,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x0e,  0x00,  0x00}},
		},
		{
			                           R:0x0f,G:0x00,B:0x00,
			Expected: Type{value:[3]byte{0x0f,  0x00,  0x00}},
		},



		{
			                           R:0x00,G:0x01,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x01,  0x00}},
		},
		{
			                           R:0x00,G:0x02,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x02,  0x00}},
		},
		{
			                           R:0x00,G:0x03,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x03,  0x00}},
		},
		{
			                           R:0x00,G:0x04,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x04,  0x00}},
		},
		{
			                           R:0x00,G:0x05,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x05,  0x00}},
		},
		{
			                           R:0x00,G:0x06,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x06,  0x00}},
		},
		{
			                           R:0x00,G:0x07,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x07,  0x00}},
		},
		{
			                           R:0x00,G:0x08,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x08,  0x00}},
		},
		{
			                           R:0x00,G:0x09,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x09,  0x00}},
		},
		{
			                           R:0x00,G:0x0a,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x0a,  0x00}},
		},
		{
			                           R:0x00,G:0x0b,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x0b,  0x00}},
		},
		{
			                           R:0x00,G:0x0c,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x0c,  0x00}},
		},
		{
			                           R:0x00,G:0x0d,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x0d,  0x00}},
		},
		{
			                           R:0x00,G:0x0e,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x0e,  0x00}},
		},
		{
			                           R:0x00,G:0x0f,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x0f,  0x00}},
		},
		{
			                           R:0x00,G:0x10,B:0x00,
			Expected: Type{value:[3]byte{0x00,  0x10,  0x00}},
		},



		{
			                           R:0x00,G:0x00,B:0x01,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x01}},
		},
		{
			                           R:0x00,G:0x00,B:0x02,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x02}},
		},
		{
			                           R:0x00,G:0x00,B:0x03,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x03}},
		},
		{
			                           R:0x00,G:0x00,B:0x04,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x04}},
		},
		{
			                           R:0x00,G:0x00,B:0x05,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x05}},
		},
		{
			                           R:0x00,G:0x00,B:0x06,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x06}},
		},
		{
			                           R:0x00,G:0x00,B:0x07,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x07}},
		},
		{
			                           R:0x00,G:0x00,B:0x08,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x08}},
		},
		{
			                           R:0x00,G:0x00,B:0x09,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x09}},
		},
		{
			                           R:0x00,G:0x00,B:0x0a,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x0a}},
		},
		{
			                           R:0x00,G:0x00,B:0x0b,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x0b}},
		},
		{
			                           R:0x00,G:0x00,B:0x0c,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x0c}},
		},
		{
			                           R:0x00,G:0x00,B:0x0d,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x0d}},
		},
		{
			                           R:0x00,G:0x00,B:0x0e,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x0e}},
		},
		{
			                           R:0x00,G:0x00,B:0x0f,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x0f}},
		},
		{
			                           R:0x00,G:0x00,B:0x10,
			Expected: Type{value:[3]byte{0x00,  0x00,  0x10}},
		},



		{
			                           R:0xab,G:0xcd,B:0xef,
			Expected: Type{value:[3]byte{0xab,  0xcd,  0xef}},
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
