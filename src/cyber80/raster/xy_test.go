package cyber80_raster_test

import (
	"github.com/reiver/cyber80/src/cyber80/raster"

	"testing"
)

func TestXY(t *testing.T) {

	tests := []struct{
		I int
		ExpectedX int
		ExpectedY int
	}{
		{
			I:         0,
			ExpectedX: 0,
			ExpectedY: 0,
		},
		{
			I:         1,
			ExpectedX: 1,
			ExpectedY: 0,
		},
		{
			I:         2,
			ExpectedX: 2,
			ExpectedY: 0,
		},
		{
			I:         3,
			ExpectedX: 3,
			ExpectedY: 0,
		},



		{
			I:         124,
			ExpectedX: 124,
			ExpectedY: 0,
		},
		{
			I:         125,
			ExpectedX: 125,
			ExpectedY: 0,
		},
		{
			I:         126,
			ExpectedX: 126,
			ExpectedY: 0,
		},
		{
			I:         127,
			ExpectedX: 127,
			ExpectedY: 0,
		},



		{
			I:         128,
			ExpectedX: 0,
			ExpectedY: 1,
		},
		{
			I:         129,
			ExpectedX: 1,
			ExpectedY: 1,
		},
		{
			I:         130,
			ExpectedX: 2,
			ExpectedY: 1,
		},
		{
			I:         131,
			ExpectedX: 3,
			ExpectedY: 1,
		},



		{
			I:         252,
			ExpectedX: 124,
			ExpectedY: 1,
		},
		{
			I:         253,
			ExpectedX: 125,
			ExpectedY: 1,
		},
		{
			I:         254,
			ExpectedX: 126,
			ExpectedY: 1,
		},
		{
			I:         255,
			ExpectedX: 127,
			ExpectedY: 1,
		},



		{
			I:         256,
			ExpectedX: 0,
			ExpectedY: 2,
		},
		{
			I:         257,
			ExpectedX: 1,
			ExpectedY: 2,
		},
		{
			I:         258,
			ExpectedX: 2,
			ExpectedY: 2,
		},
		{
			I:         259,
			ExpectedX: 3,
			ExpectedY: 2,
		},



		{
			I: (cyber80_raster.Size-cyber80_raster.Width-4),
			ExpectedX:             (cyber80_raster.Width-4),
			ExpectedY: (cyber80_raster.Height-2),
		},
		{
			I: (cyber80_raster.Size-cyber80_raster.Width-3),
			ExpectedX:             (cyber80_raster.Width-3),
			ExpectedY: (cyber80_raster.Height-2),
		},
		{
			I: (cyber80_raster.Size-cyber80_raster.Width-2),
			ExpectedX:             (cyber80_raster.Width-2),
			ExpectedY: (cyber80_raster.Height-2),
		},
		{
			I: (cyber80_raster.Size-cyber80_raster.Width-1),
			ExpectedX:             (cyber80_raster.Width-1),
			ExpectedY: (cyber80_raster.Height-2),
		},



		{
			I: (cyber80_raster.Size-cyber80_raster.Width+0),
			ExpectedX:                                   0,
			ExpectedY: (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-cyber80_raster.Width+1),
			ExpectedX:                                   1,
			ExpectedY: (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-cyber80_raster.Width+2),
			ExpectedX:                                   2,
			ExpectedY: (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-cyber80_raster.Width+3),
			ExpectedX:                                   3,
			ExpectedY: (cyber80_raster.Height-1),
		},



		{
			I:          (cyber80_raster.Size-4),
			ExpectedX: (cyber80_raster.Width-4),
			ExpectedY: (cyber80_raster.Height-1),
		},
		{
			I:          (cyber80_raster.Size-3),
			ExpectedX: (cyber80_raster.Width-3),
			ExpectedY: (cyber80_raster.Height-1),
		},
		{
			I:          (cyber80_raster.Size-2),
			ExpectedX: (cyber80_raster.Width-2),
			ExpectedY: (cyber80_raster.Height-1),
		},
		{
			I:          (cyber80_raster.Size-1),
			ExpectedX: (cyber80_raster.Width-1),
			ExpectedY: (cyber80_raster.Height-1),
		},
	}

	{
		w := cyber80_raster.Width
		h := cyber80_raster.Height

		for row:=0; row<h; row++ {

			for i:=(row*w); i<((row+1)*w); i++ {
				test := struct{
					I int
					ExpectedX int
					ExpectedY int
				}{
					I:         i,
					ExpectedX: i - (row*w),
					ExpectedY: row,
				}

				tests = append(tests, test)
			}

		}
	}

	for testNumber, test := range tests {

		actualX, actualY := cyber80_raster.XY(test.I)

		if expected, actual := test.ExpectedX, actualX; expected != actual {
			t.Errorf("For test #%d, actual ‘x’ is not what was expected.", testNumber)
			t.Logf("(i) = (%d)", test.I)
			t.Logf("EXPECTED (x, y) = (%d, %d)", test.ExpectedX, test.ExpectedY)
			t.Logf("ACTUAL:  (x, y) = (%d, %d)", actualX, actualY)
			continue
		}

		if expected, actual := test.ExpectedY, actualY; expected != actual {
			t.Errorf("For test #%d, actual y is not what was expected.", testNumber)
			t.Logf("(i) = (%d)", test.I)
			t.Logf("EXPECTED (x, y) = (%d, %d)", test.ExpectedX, test.ExpectedY)
			t.Logf("ACTUAL:  (x, y) = (%d, %d)", actualX, actualY)
			continue
		}
	}
}
