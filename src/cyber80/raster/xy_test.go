package cyber80_raster_test

import (
	"github.com/reiver/cyber80/src/cyber80/raster"

	"testing"
)

func TestXY(t *testing.T) {

	tests := []struct{
		I int
		ExpectedX1 int
		ExpectedX2 int
		ExpectedY int
	}{
		{
			I: 0,
			ExpectedX1: 0,
			ExpectedX2: 1,
			ExpectedY:  0,
		},
		{
			I: 1,
			ExpectedX1: 2,
			ExpectedX2: 3,
			ExpectedY:  0,
		},
		{
			I: 2,
			ExpectedX1: 4,
			ExpectedX2: 5,
			ExpectedY:  0,
		},
		{
			I: 3,
			ExpectedX1: 6,
			ExpectedX2: 7,
			ExpectedY:  0,
		},
		{
			I: 4,
			ExpectedX1: 8,
			ExpectedX2: 9,
			ExpectedY:  0,
		},
		{
			I: 5,
			ExpectedX1: 10,
			ExpectedX2: 11,
			ExpectedY:  0,
		},
		{
			I: 6,
			ExpectedX1: 12,
			ExpectedX2: 13,
			ExpectedY:  0,
		},
		{
			I: 7,
			ExpectedX1: 14,
			ExpectedX2: 15,
			ExpectedY:  0,
		},



		{
			I: 8,
			ExpectedX1: 16,
			ExpectedX2: 17,
			ExpectedY:  0,
		},
		{
			I: 9,
			ExpectedX1: 18,
			ExpectedX2: 19,
			ExpectedY:  0,
		},
		{
			I: 10,
			ExpectedX1: 20,
			ExpectedX2: 21,
			ExpectedY:  0,
		},
		{
			I: 11,
			ExpectedX1: 22,
			ExpectedX2: 23,
			ExpectedY:  0,
		},
		{
			I: 12,
			ExpectedX1: 24,
			ExpectedX2: 25,
			ExpectedY:  0,
		},
		{
			I: 13,
			ExpectedX1: 26,
			ExpectedX2: 27,
			ExpectedY:  0,
		},
		{
			I: 14,
			ExpectedX1: 28,
			ExpectedX2: 29,
			ExpectedY:  0,
		},
		{
			I: 15,
			ExpectedX1: 30,
			ExpectedX2: 31,
			ExpectedY:  0,
		},



		{
			I: 16,
			ExpectedX1: 32,
			ExpectedX2: 33,
			ExpectedY:  0,
		},
		{
			I: 17,
			ExpectedX1: 34,
			ExpectedX2: 35,
			ExpectedY:  0,
		},
		{
			I: 18,
			ExpectedX1: 36,
			ExpectedX2: 37,
			ExpectedY:  0,
		},
		{
			I: 19,
			ExpectedX1: 38,
			ExpectedX2: 39,
			ExpectedY:  0,
		},
		{
			I: 20,
			ExpectedX1: 40,
			ExpectedX2: 41,
			ExpectedY:  0,
		},
		{
			I: 21,
			ExpectedX1: 42,
			ExpectedX2: 43,
			ExpectedY:  0,
		},
		{
			I: 22,
			ExpectedX1: 44,
			ExpectedX2: 45,
			ExpectedY:  0,
		},
		{
			I: 23,
			ExpectedX1: 46,
			ExpectedX2: 47,
			ExpectedY:  0,
		},



		{
			I: 24,
			ExpectedX1: 48,
			ExpectedX2: 49,
			ExpectedY:  0,
		},
		{
			I: 25,
			ExpectedX1: 50,
			ExpectedX2: 51,
			ExpectedY:  0,
		},
		{
			I: 26,
			ExpectedX1: 52,
			ExpectedX2: 53,
			ExpectedY:  0,
		},
		{
			I: 27,
			ExpectedX1: 54,
			ExpectedX2: 55,
			ExpectedY:  0,
		},
		{
			I: 28,
			ExpectedX1: 56,
			ExpectedX2: 57,
			ExpectedY:  0,
		},
		{
			I: 29,
			ExpectedX1: 58,
			ExpectedX2: 59,
			ExpectedY:  0,
		},
		{
			I: 30,
			ExpectedX1: 60,
			ExpectedX2: 61,
			ExpectedY:  0,
		},
		{
			I: 31,
			ExpectedX1: 62,
			ExpectedX2: 63,
			ExpectedY:  0,
		},



		{
			I: 32,
			ExpectedX1: 64,
			ExpectedX2: 65,
			ExpectedY:  0,
		},
		{
			I: 33,
			ExpectedX1: 66,
			ExpectedX2: 67,
			ExpectedY:  0,
		},
		{
			I: 34,
			ExpectedX1: 68,
			ExpectedX2: 69,
			ExpectedY:  0,
		},
		{
			I: 35,
			ExpectedX1: 70,
			ExpectedX2: 71,
			ExpectedY:  0,
		},
		{
			I: 36,
			ExpectedX1: 72,
			ExpectedX2: 73,
			ExpectedY:  0,
		},
		{
			I: 37,
			ExpectedX1: 74,
			ExpectedX2: 75,
			ExpectedY:  0,
		},
		{
			I: 38,
			ExpectedX1: 76,
			ExpectedX2: 77,
			ExpectedY:  0,
		},
		{
			I: 39,
			ExpectedX1: 78,
			ExpectedX2: 79,
			ExpectedY:  0,
		},



		{
			I: 40,
			ExpectedX1: 80,
			ExpectedX2: 81,
			ExpectedY:  0,
		},
		{
			I: 41,
			ExpectedX1: 82,
			ExpectedX2: 83,
			ExpectedY:  0,
		},
		{
			I: 42,
			ExpectedX1: 84,
			ExpectedX2: 85,
			ExpectedY:  0,
		},
		{
			I: 43,
			ExpectedX1: 86,
			ExpectedX2: 87,
			ExpectedY:  0,
		},
		{
			I: 44,
			ExpectedX1: 88,
			ExpectedX2: 89,
			ExpectedY:  0,
		},
		{
			I: 45,
			ExpectedX1: 90,
			ExpectedX2: 91,
			ExpectedY:  0,
		},
		{
			I: 46,
			ExpectedX1: 92,
			ExpectedX2: 93,
			ExpectedY:  0,
		},
		{
			I: 47,
			ExpectedX1: 94,
			ExpectedX2: 95,
			ExpectedY:  0,
		},



		{
			I: 48,
			ExpectedX1: 96,
			ExpectedX2: 97,
			ExpectedY:  0,
		},
		{
			I: 49,
			ExpectedX1: 98,
			ExpectedX2: 99,
			ExpectedY:  0,
		},
		{
			I: 50,
			ExpectedX1: 100,
			ExpectedX2: 101,
			ExpectedY:  0,
		},
		{
			I: 51,
			ExpectedX1: 102,
			ExpectedX2: 103,
			ExpectedY:  0,
		},
		{
			I: 52,
			ExpectedX1: 104,
			ExpectedX2: 105,
			ExpectedY:  0,
		},
		{
			I: 53,
			ExpectedX1: 106,
			ExpectedX2: 107,
			ExpectedY:  0,
		},
		{
			I: 54,
			ExpectedX1: 108,
			ExpectedX2: 109,
			ExpectedY:  0,
		},
		{
			I: 55,
			ExpectedX1: 110,
			ExpectedX2: 111,
			ExpectedY:  0,
		},



		{
			I: 56,
			ExpectedX1: 112,
			ExpectedX2: 113,
			ExpectedY:  0,
		},
		{
			I: 57,
			ExpectedX1: 114,
			ExpectedX2: 115,
			ExpectedY:  0,
		},
		{
			I: 58,
			ExpectedX1: 116,
			ExpectedX2: 117,
			ExpectedY:  0,
		},
		{
			I: 59,
			ExpectedX1: 118,
			ExpectedX2: 119,
			ExpectedY:  0,
		},
		{
			I: 60,
			ExpectedX1: 120,
			ExpectedX2: 121,
			ExpectedY:  0,
		},
		{
			I: 61,
			ExpectedX1: 122,
			ExpectedX2: 123,
			ExpectedY:  0,
		},
		{
			I: 62,
			ExpectedX1: 124,
			ExpectedX2: 125,
			ExpectedY:  0,
		},
		{
			I: 63,
			ExpectedX1: 126,
			ExpectedX2: 127,
			ExpectedY:  0,
		},



		{
			I: 64,
			ExpectedX1: 0,
			ExpectedX2: 1,
			ExpectedY:  1,
		},
		{
			I: 65,
			ExpectedX1: 2,
			ExpectedX2: 3,
			ExpectedY:  1,
		},
		{
			I: 66,
			ExpectedX1: 4,
			ExpectedX2: 5,
			ExpectedY:  1,
		},
		{
			I: 67,
			ExpectedX1: 6,
			ExpectedX2: 7,
			ExpectedY:  1,
		},
		{
			I: 68,
			ExpectedX1: 8,
			ExpectedX2: 9,
			ExpectedY:  1,
		},
		{
			I: 69,
			ExpectedX1: 10,
			ExpectedX2: 11,
			ExpectedY:  1,
		},
		{
			I: 70,
			ExpectedX1: 12,
			ExpectedX2: 13,
			ExpectedY:  1,
		},
		{
			I: 71,
			ExpectedX1: 14,
			ExpectedX2: 15,
			ExpectedY:  1,
		},



		{
			I: 72,
			ExpectedX1: 16,
			ExpectedX2: 17,
			ExpectedY:  1,
		},
		{
			I: 73,
			ExpectedX1: 18,
			ExpectedX2: 19,
			ExpectedY:  1,
		},
		{
			I: 74,
			ExpectedX1: 20,
			ExpectedX2: 21,
			ExpectedY:  1,
		},
		{
			I: 75,
			ExpectedX1: 22,
			ExpectedX2: 23,
			ExpectedY:  1,
		},
		{
			I: 76,
			ExpectedX1: 24,
			ExpectedX2: 25,
			ExpectedY:  1,
		},
		{
			I: 77,
			ExpectedX1: 26,
			ExpectedX2: 27,
			ExpectedY:  1,
		},
		{
			I: 78,
			ExpectedX1: 28,
			ExpectedX2: 29,
			ExpectedY:  1,
		},
		{
			I: 79,
			ExpectedX1: 30,
			ExpectedX2: 31,
			ExpectedY:  1,
		},






		{
			I: (cyber80_raster.Size-8),
			ExpectedX1: (cyber80_raster.Width-16),
			ExpectedX2: (cyber80_raster.Width-15),
			ExpectedY:  (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-7),
			ExpectedX1: (cyber80_raster.Width-14),
			ExpectedX2: (cyber80_raster.Width-13),
			ExpectedY:  (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-6),
			ExpectedX1: (cyber80_raster.Width-12),
			ExpectedX2: (cyber80_raster.Width-11),
			ExpectedY:  (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-5),
			ExpectedX1: (cyber80_raster.Width-10),
			ExpectedX2: (cyber80_raster.Width-9),
			ExpectedY:  (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-4),
			ExpectedX1: (cyber80_raster.Width-8),
			ExpectedX2: (cyber80_raster.Width-7),
			ExpectedY:  (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-3),
			ExpectedX1: (cyber80_raster.Width-6),
			ExpectedX2: (cyber80_raster.Width-5),
			ExpectedY:  (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-2),
			ExpectedX1: (cyber80_raster.Width-4),
			ExpectedX2: (cyber80_raster.Width-3),
			ExpectedY:  (cyber80_raster.Height-1),
		},
		{
			I: (cyber80_raster.Size-1),
			ExpectedX1: (cyber80_raster.Width-2),
			ExpectedX2: (cyber80_raster.Width-1),
			ExpectedY:  (cyber80_raster.Height-1),
		},
	}

	for testNumber, test := range tests {

		actualX1, actualX2, actualY := cyber80_raster.XY(test.I)

		if expected, actual := test.ExpectedX1, actualX1; expected != actual {
			t.Errorf("For test #%d, actual x₁ is not what was expected.", testNumber)
			t.Logf("(i) = (%d)", test.I)
			t.Logf("EXPECTED (x₁, x₂, y) = (%d, %d, %d)", test.ExpectedX1, test.ExpectedX2, test.ExpectedY)
			t.Logf("ACTUAL:  (x₁, x₂, y) = (%d, %d, %d)", actualX1, actualX2, actualY)
			continue
		}

		if expected, actual := test.ExpectedX2, actualX2; expected != actual {
			t.Errorf("For test #%d, actual x₂ is not what was expected.", testNumber)
			t.Logf("(i) = (%d)", test.I)
			t.Logf("EXPECTED (x₁, x₂, y) = (%d, %d, %d)", test.ExpectedX1, test.ExpectedX2, test.ExpectedY)
			t.Logf("ACTUAL:  (x₁, x₂, y) = (%d, %d, %d)", actualX1, actualX2, actualY)
			continue
		}

		if expected, actual := test.ExpectedY, actualY; expected != actual {
			t.Errorf("For test #%d, actual y is not what was expected.", testNumber)
			t.Logf("(i) = (%d)", test.I)
			t.Logf("EXPECTED (x₁, x₂, y) = (%d, %d, %d)", test.ExpectedX1, test.ExpectedX2, test.ExpectedY)
			t.Logf("ACTUAL:  (x₁, x₂, y) = (%d, %d, %d)", actualX1, actualX2, actualY)
			continue
		}
	}
}
