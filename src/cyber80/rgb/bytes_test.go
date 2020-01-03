package cyber80_rgb

import (
	"testing"
)

func TestBytes(t *testing.T) {

	tests := []struct{
		Bytes []byte
		Expected Type
	}{
		{
			Bytes: []byte(nil),
			Expected: Type{},
		},
		{
			Bytes: []byte{},
			Expected: Type{},
		},









		{
			Bytes:    []byte{0x00},
			Expected: Type{r:0x00},
		},
		{
			Bytes:    []byte{0x01},
			Expected: Type{r:0x01},
		},
		{
			Bytes:    []byte{0x02},
			Expected: Type{r:0x02},
		},
		{
			Bytes:    []byte{0x03},
			Expected: Type{r:0x03},
		},
		{
			Bytes:    []byte{0x04},
			Expected: Type{r:0x04},
		},
		{
			Bytes:    []byte{0x05},
			Expected: Type{r:0x05},
		},
		{
			Bytes:    []byte{0x06},
			Expected: Type{r:0x06},
		},
		{
			Bytes:    []byte{0x07},
			Expected: Type{r:0x07},
		},
		{
			Bytes:    []byte{0x08},
			Expected: Type{r:0x08},
		},
		{
			Bytes:    []byte{0x09},
			Expected: Type{r:0x09},
		},
		{
			Bytes:    []byte{0x0a},
			Expected: Type{r:0x0a},
		},
		{
			Bytes:    []byte{0x0b},
			Expected: Type{r:0x0b},
		},
		{
			Bytes:    []byte{0x0c},
			Expected: Type{r:0x0c},
		},
		{
			Bytes:    []byte{0x0d},
			Expected: Type{r:0x0d},
		},
		{
			Bytes:    []byte{0x0e},
			Expected: Type{r:0x0e},
		},
		{
			Bytes:    []byte{0x0f},
			Expected: Type{r:0x0f},
		},
		{
			Bytes:    []byte{0x10},
			Expected: Type{r:0x10},
		},
		{
			Bytes:    []byte{0x11},
			Expected: Type{r:0x11},
		},
		{
			Bytes:    []byte{0x12},
			Expected: Type{r:0x12},
		},
		{
			Bytes:    []byte{0x13},
			Expected: Type{r:0x13},
		},
		{
			Bytes:    []byte{0x14},
			Expected: Type{r:0x14},
		},
		{
			Bytes:    []byte{0x15},
			Expected: Type{r:0x15},
		},
		{
			Bytes:    []byte{0x16},
			Expected: Type{r:0x16},
		},
		{
			Bytes:    []byte{0x17},
			Expected: Type{r:0x17},
		},
		{
			Bytes:    []byte{0x18},
			Expected: Type{r:0x18},
		},
		{
			Bytes:    []byte{0x19},
			Expected: Type{r:0x19},
		},
		{
			Bytes:    []byte{0x1a},
			Expected: Type{r:0x1a},
		},
		{
			Bytes:    []byte{0x1b},
			Expected: Type{r:0x1b},
		},
		{
			Bytes:    []byte{0x1c},
			Expected: Type{r:0x1c},
		},
		{
			Bytes:    []byte{0x1d},
			Expected: Type{r:0x1d},
		},
		{
			Bytes:    []byte{0x1e},
			Expected: Type{r:0x1e},
		},
		{
			Bytes:    []byte{0x1f},
			Expected: Type{r:0x1f},
		},
		{
			Bytes:    []byte{0x20},
			Expected: Type{r:0x20},
		},
		{
			Bytes:    []byte{0x21},
			Expected: Type{r:0x21},
		},
		{
			Bytes:    []byte{0x22},
			Expected: Type{r:0x22},
		},
		{
			Bytes:    []byte{0x23},
			Expected: Type{r:0x23},
		},
		{
			Bytes:    []byte{0x24},
			Expected: Type{r:0x24},
		},
		{
			Bytes:    []byte{0x25},
			Expected: Type{r:0x25},
		},
		{
			Bytes:    []byte{0x26},
			Expected: Type{r:0x26},
		},
		{
			Bytes:    []byte{0x27},
			Expected: Type{r:0x27},
		},
		{
			Bytes:    []byte{0x28},
			Expected: Type{r:0x28},
		},
		{
			Bytes:    []byte{0x29},
			Expected: Type{r:0x29},
		},
		{
			Bytes:    []byte{0x2a},
			Expected: Type{r:0x2a},
		},
		{
			Bytes:    []byte{0x2b},
			Expected: Type{r:0x2b},
		},
		{
			Bytes:    []byte{0x2c},
			Expected: Type{r:0x2c},
		},
		{
			Bytes:    []byte{0x2d},
			Expected: Type{r:0x2d},
		},
		{
			Bytes:    []byte{0x2e},
			Expected: Type{r:0x2e},
		},
		{
			Bytes:    []byte{0x2f},
			Expected: Type{r:0x2f},
		},
		{
			Bytes:    []byte{0x30},
			Expected: Type{r:0x30},
		},









		{
			Bytes:    []byte{0x00,   0x00},
			Expected: Type{r:0x00, g:0x00},
		},
		{
			Bytes:    []byte{0x01,   0x00},
			Expected: Type{r:0x01, g:0x00},
		},
		{
			Bytes:    []byte{0x02,   0x00},
			Expected: Type{r:0x02, g:0x00},
		},
		{
			Bytes:    []byte{0x03,   0x00},
			Expected: Type{r:0x03, g:0x00},
		},
		{
			Bytes:    []byte{0x04,   0x00},
			Expected: Type{r:0x04, g:0x00},
		},



		{
			Bytes:    []byte{0x00,   0x01},
			Expected: Type{r:0x00, g:0x01},
		},
		{
			Bytes:    []byte{0x01,   0x01},
			Expected: Type{r:0x01, g:0x01},
		},
		{
			Bytes:    []byte{0x02,   0x01},
			Expected: Type{r:0x02, g:0x01},
		},
		{
			Bytes:    []byte{0x03,   0x01},
			Expected: Type{r:0x03, g:0x01},
		},
		{
			Bytes:    []byte{0x04,   0x01},
			Expected: Type{r:0x04, g:0x01},
		},



		{
			Bytes:    []byte{0x00,   0x02},
			Expected: Type{r:0x00, g:0x02},
		},
		{
			Bytes:    []byte{0x01,   0x02},
			Expected: Type{r:0x01, g:0x02},
		},
		{
			Bytes:    []byte{0x02,   0x02},
			Expected: Type{r:0x02, g:0x02},
		},
		{
			Bytes:    []byte{0x03,   0x02},
			Expected: Type{r:0x03, g:0x02},
		},
		{
			Bytes:    []byte{0x04,   0x02},
			Expected: Type{r:0x04, g:0x02},
		},









		{
			Bytes:    []byte{0x00,   0x00,   0x00},
			Expected: Type{r:0x00, g:0x00, b:0x00},
		},
		{
			Bytes:    []byte{0x01,   0x00,   0x00},
			Expected: Type{r:0x01, g:0x00, b:0x00},
		},
		{
			Bytes:    []byte{0x02,   0x00,   0x00},
			Expected: Type{r:0x02, g:0x00, b:0x00},
		},
		{
			Bytes:    []byte{0x03,   0x00,   0x00},
			Expected: Type{r:0x03, g:0x00, b:0x00},
		},
		{
			Bytes:    []byte{0x04,   0x00,   0x00},
			Expected: Type{r:0x04, g:0x00, b:0x00},
		},



		{
			Bytes:    []byte{0x00,   0x01,   0x00},
			Expected: Type{r:0x00, g:0x01, b:0x00},
		},
		{
			Bytes:    []byte{0x01,   0x01,   0x00},
			Expected: Type{r:0x01, g:0x01, b:0x00},
		},
		{
			Bytes:    []byte{0x02,   0x01,   0x00},
			Expected: Type{r:0x02, g:0x01, b:0x00},
		},
		{
			Bytes:    []byte{0x03,   0x01,   0x00},
			Expected: Type{r:0x03, g:0x01, b:0x00},
		},
		{
			Bytes:    []byte{0x04,   0x01,   0x00},
			Expected: Type{r:0x04, g:0x01, b:0x00},
		},



		{
			Bytes:    []byte{0x00,   0x02,   0x00},
			Expected: Type{r:0x00, g:0x02, b:0x00},
		},
		{
			Bytes:    []byte{0x01,   0x02,   0x00},
			Expected: Type{r:0x01, g:0x02, b:0x00},
		},
		{
			Bytes:    []byte{0x02,   0x02,   0x00},
			Expected: Type{r:0x02, g:0x02, b:0x00},
		},
		{
			Bytes:    []byte{0x03,   0x02,   0x00},
			Expected: Type{r:0x03, g:0x02, b:0x00},
		},
		{
			Bytes:    []byte{0x04,   0x02,   0x00},
			Expected: Type{r:0x04, g:0x02, b:0x00},
		},



		{
			Bytes:    []byte{0x00,   0x03,   0x00},
			Expected: Type{r:0x00, g:0x03, b:0x00},
		},
		{
			Bytes:    []byte{0x01,   0x03,   0x00},
			Expected: Type{r:0x01, g:0x03, b:0x00},
		},
		{
			Bytes:    []byte{0x02,   0x03,   0x00},
			Expected: Type{r:0x02, g:0x03, b:0x00},
		},
		{
			Bytes:    []byte{0x03,   0x03,   0x00},
			Expected: Type{r:0x03, g:0x03, b:0x00},
		},
		{
			Bytes:    []byte{0x04,   0x03,   0x00},
			Expected: Type{r:0x04, g:0x03, b:0x00},
		},



		{
			Bytes:    []byte{0x00,   0x04,   0x00},
			Expected: Type{r:0x00, g:0x04, b:0x00},
		},
		{
			Bytes:    []byte{0x01,   0x04,   0x00},
			Expected: Type{r:0x01, g:0x04, b:0x00},
		},
		{
			Bytes:    []byte{0x02,   0x04,   0x00},
			Expected: Type{r:0x02, g:0x04, b:0x00},
		},
		{
			Bytes:    []byte{0x03,   0x04,   0x00},
			Expected: Type{r:0x03, g:0x04, b:0x00},
		},
		{
			Bytes:    []byte{0x04,   0x04,   0x00},
			Expected: Type{r:0x04, g:0x04, b:0x00},
		},



		{
			Bytes:    []byte{0xab,   0xcd,   0xef},
			Expected: Type{r:0xab, g:0xcd, b:0xef},
		},
	}

	for testNumber, test := range tests {

		actual := Bytes(test.Bytes)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, did not actually get what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL: %#v", actual)
			continue
		}
	}
}
