package cyber80_raster_test

import (
	"github.com/reiver/cyber80/src/cyber80/raster"

	"testing"
)

func TestIndex(t *testing.T) {

	tests := []struct{
		X int
		Y int
		Expected int
	}{
		{
			X:          0,
			Y:          0,
			Expected:   0,
		},
		{
			X:          1,
			Y:          0,
			Expected:   1,
		},
		{
			X:          2,
			Y:          0,
			Expected:   2,
		},
		{
			X:          3,
			Y:          0,
			Expected:   3,
		},



		{
			X:        124,
			Y:          0,
			Expected: 124,
		},
		{
			X:        125,
			Y:        0,
			Expected: 125,
		},
		{
			X:        126,
			Y:        0,
			Expected: 126,
		},
		{
			X:        127,
			Y:        0,
			Expected: 127,
		},



		{
			X:        0,
			Y:        1,
			Expected: 128,
		},
		{
			X:        1,
			Y:        1,
			Expected: 129,
		},
		{
			X:        2,
			Y:        1,
			Expected: 130,
		},
		{
			X:        3,
			Y:        1,
			Expected: 131,
		},



		{
			X:        124,
			Y:        1,
			Expected: 252,
		},
		{
			X:        125,
			Y:        1,
			Expected: 253,
		},
		{
			X:        126,
			Y:        1,
			Expected: 254,
		},
		{
			X:        127,
			Y:        1,
			Expected: 255,
		},



		{
			X:        0,
			Y:        2,
			Expected: 256,
		},
		{
			X:        1,
			Y:        2,
			Expected: 257,
		},
		{
			X:        2,
			Y:        2,
			Expected: 258,
		},
		{
			X:        3,
			Y:        2,
			Expected: 259,
		},



		{
			X:                                    (cyber80_raster.Width-4),
			Y:        (cyber80_raster.Height-2),
			Expected: ((cyber80_raster.Width*(cyber80_raster.Height-1))-4),
		},
		{
			X:                                    (cyber80_raster.Width-3),
			Y:        (cyber80_raster.Height-2),
			Expected: ((cyber80_raster.Width*(cyber80_raster.Height-1))-3),
		},
		{
			X:                                    (cyber80_raster.Width-2),
			Y:        (cyber80_raster.Height-2),
			Expected: ((cyber80_raster.Width*(cyber80_raster.Height-1))-2),
		},
		{
			X:                                    (cyber80_raster.Width-1),
			Y:        (cyber80_raster.Height-2),
			Expected: ((cyber80_raster.Width*(cyber80_raster.Height-1))-1),
		},



		{
			X:                                                                           0,
			Y:        (cyber80_raster.Height-1),
			Expected: ((cyber80_raster.Width*cyber80_raster.Height)-cyber80_raster.Width+0),
		},
		{
			X:                                                                           1,
			Y:        (cyber80_raster.Height-1),
			Expected: ((cyber80_raster.Width*cyber80_raster.Height)-cyber80_raster.Width+1),
		},
		{
			X:                                                                           2,
			Y:        (cyber80_raster.Height-1),
			Expected: ((cyber80_raster.Width*cyber80_raster.Height)-cyber80_raster.Width+2),
		},
		{
			X:                                                                           3,
			Y:        (cyber80_raster.Height-1),
			Expected: ((cyber80_raster.Width*cyber80_raster.Height)-cyber80_raster.Width+3),
		},



		{
			X:                                (cyber80_raster.Width-4),
			Y:        (cyber80_raster.Height-1),
			Expected: ((cyber80_raster.Width*cyber80_raster.Height)-4),
		},
		{
			X:                                (cyber80_raster.Width-3),
			Y:        (cyber80_raster.Height-1),
			Expected: ((cyber80_raster.Width*cyber80_raster.Height)-3),
		},
		{
			X:                                (cyber80_raster.Width-2),
			Y:        (cyber80_raster.Height-1),
			Expected: ((cyber80_raster.Width*cyber80_raster.Height)-2),
		},
		{
			X:                                (cyber80_raster.Width-1),
			Y:        (cyber80_raster.Height-1),
			Expected: ((cyber80_raster.Width*cyber80_raster.Height)-1),
		},
	}

	for testNumber, test := range tests {

		index := cyber80_raster.Index(test.X, test.Y)

		if expected, actual := test.Expected, index; expected != actual {
			t.Errorf("For test #%d, the excepted ‘index’ was not what was actually gotten.", testNumber)
			t.Logf("EXPECTED: (x=%d, y=%d) -> %d", test.X, test.Y, expected)
			t.Logf("ACTUAL:   (x=%d, y=%d) -> %d", test.X, test.Y, actual)
			continue
		}
	}
}
