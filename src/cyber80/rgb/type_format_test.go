package cyber80_rgb_test

import (
	"github.com/reiver/cyber80/src/cyber80/rgb"

	"fmt"

	"testing"
)

func TestTypeFormat(t *testing.T) {

	tests := []struct{
		RGB      cyber80_rgb.Type
		Format  string
		Expected string
	}{
		{
			RGB: cyber80_rgb.Bytes([]byte{0x00,0x00,0x00}),
			Format:   "%z",
			Expected: "rgb(0,0,0)",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0x00,0x00,0x00}),
			Format:   "%x",
			Expected: "000000",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0x00,0x00,0x00}),
			Format:   "0x%x",
			Expected: "0x000000",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0x00,0x00,0x00}),
			Format:   "#%x",
			Expected: "#000000",
		},



		{
			RGB: cyber80_rgb.Bytes([]byte{0x01,0x23,0x45}),
			Format:   "%z",
			Expected: "rgb(1,35,69)",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0x01,0x23,0x45}),
			Format:   "%x",
			Expected: "012345",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0x01,0x23,0x45}),
			Format:   "0x%x",
			Expected: "0x012345",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0x01,0x23,0x45}),
			Format:   "#%x",
			Expected: "#012345",
		},



		{
			RGB: cyber80_rgb.Bytes([]byte{0xab,0xcd,0xef}),
			Format:   "%z",
			Expected: "rgb(171,205,239)",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0xab,0xcd,0xef}),
			Format:   "%x",
			Expected: "abcdef",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0xab,0xcd,0xef}),
			Format:   "0x%x",
			Expected: "0xabcdef",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0xab,0xcd,0xef}),
			Format:   "#%x",
			Expected: "#abcdef",
		},



		{
			RGB: cyber80_rgb.Bytes([]byte{0xff,0xff,0xff}),
			Format:   "%z",
			Expected: "rgb(255,255,255)",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0xff,0xff,0xff}),
			Format:   "%x",
			Expected: "ffffff",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0xff,0xff,0xff}),
			Format:   "0x%x",
			Expected: "0xffffff",
		},
		{
			RGB: cyber80_rgb.Bytes([]byte{0xff,0xff,0xff}),
			Format:   "#%x",
			Expected: "#ffffff",
		},
	}

	for testNumber, test := range tests {

		actual := fmt.Sprintf(test.Format, test.RGB)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, did not get formatted RGB that was actually expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
