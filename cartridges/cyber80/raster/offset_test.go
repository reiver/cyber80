package cyber80_raster_test

import (
	"github.com/reiver/cyber80/src/cyber80/raster"

	"testing"
)

func TestOffset(t *testing.T) {

	tests := []struct{
		X  int
		Y  int
		BP int
		Expected int
	}{
		{
			X:          0,
			Y:          0,
			BP:         0,
			Expected:   0,
		},
		{
			X:          1,
			Y:          0,
			BP:         0,
			Expected:   1,
		},
		{
			X:          2,
			Y:          0,
			BP:         0,
			Expected:   2,
		},
		{
			X:          3,
			Y:          0,
			BP:         0,
			Expected:   3,
		},



		{
			X:          0,
			Y:          0,
			BP:         1,
			Expected:   1,
		},
		{
			X:          1,
			Y:          0,
			BP:         1,
			Expected:   2,
		},
		{
			X:          2,
			Y:          0,
			BP:         1,
			Expected:   3,
		},
		{
			X:          3,
			Y:          0,
			BP:         1,
			Expected:   4,
		},



		{
			X:          0,
			Y:          0,
			BP:       200,
			Expected: 200,
		},
		{
			X:          1,
			Y:          0,
			BP:       200,
			Expected: 201,
		},
		{
			X:          2,
			Y:          0,
			BP:       200,
			Expected: 202,
		},
		{
			X:          3,
			Y:          0,
			BP:       200,
			Expected: 203,
		},









		{
			X:        124,
			Y:          0,
			BP:         0,
			Expected: 124,
		},
		{
			X:        125,
			Y:          0,
			BP:         0,
			Expected: 125,
		},
		{
			X:        126,
			Y:          0,
			BP:         0,
			Expected: 126,
		},
		{
			X:        127,
			Y:          0,
			BP:         0,
			Expected: 127,
		},



		{
			X:        124,
			Y:          0,
			BP:         1,
			Expected: 125,
		},
		{
			X:        125,
			Y:          0,
			BP:         1,
			Expected: 126,
		},
		{
			X:        126,
			Y:          0,
			BP:         1,
			Expected: 127,
		},
		{
			X:        127,
			Y:          0,
			BP:         1,
			Expected: 128,
		},



		{
			X:        124,
			Y:          0,
			BP:       200,
			Expected: 324,
		},
		{
			X:        125,
			Y:          0,
			BP:       200,
			Expected: 325,
		},
		{
			X:        126,
			Y:          0,
			BP:       200,
			Expected: 326,
		},
		{
			X:        127,
			Y:          0,
			BP:       200,
			Expected: 327,
		},









		{
			X:          0,
			Y:          1,
			BP:         0,
			Expected: 128,
		},
		{
			X:          1,
			Y:          1,
			BP:         0,
			Expected: 129,
		},
		{
			X:          2,
			Y:          1,
			BP:         0,
			Expected: 130,
		},
		{
			X:          3,
			Y:          1,
			BP:         0,
			Expected: 131,
		},



		{
			X:        124,
			Y:          1,
			BP:         0,
			Expected: 252,
		},
		{
			X:        125,
			Y:          1,
			BP:         0,
			Expected: 253,
		},
		{
			X:        126,
			Y:          1,
			BP:         0,
			Expected: 254,
		},
		{
			X:        127,
			Y:          1,
			BP:         0,
			Expected: 255,
		},



		{
			X:          0,
			Y:          2,
			BP:         0,
			Expected: 256,
		},
		{
			X:          1,
			Y:          2,
			BP:         0,
			Expected: 257,
		},
		{
			X:          2,
			Y:          2,
			BP:         0,
			Expected: 258,
		},
		{
			X:          3,
			Y:          2,
			BP:         0,
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

		offset := cyber80_raster.Offset(test.X, test.Y, test.BP)

		if expected, actual := test.Expected, offset; expected != actual {
			t.Errorf("For test #%d, the excepted ‘offset’ was not what was actually gotten.", testNumber)
			t.Logf("EXPECTED: (x=%d, y=%d) -> %d", test.X, test.Y, expected)
			t.Logf("ACTUAL:   (x=%d, y=%d) -> %d", test.X, test.Y, actual)
			continue
		}
	}
}
