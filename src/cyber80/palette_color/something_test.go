package cyber80_palette_color

import (
	"testing"
)

func TestSomething(t *testing.T) {

	tests := []struct{
		Actual Type
		Expected  Type
	}{
		{
			Actual: Something(0),
			Expected: Type{
				loaded: true,
				value: 0,
			},
		},
		{
			Actual: Something(1),
			Expected: Type{
				loaded: true,
				value: 1,
			},
		},
		{
			Actual: Something(2),
			Expected: Type{
				loaded: true,
				value: 2,
			},
		},
		{
			Actual: Something(3),
			Expected: Type{
				loaded: true,
				value: 3,
			},
		},
		{
			Actual: Something(4),
			Expected: Type{
				loaded: true,
				value: 4,
			},
		},
		{
			Actual: Something(5),
			Expected: Type{
				loaded: true,
				value: 5,
			},
		},
		{
			Actual: Something(6),
			Expected: Type{
				loaded: true,
				value: 6,
			},
		},
		{
			Actual: Something(7),
			Expected: Type{
				loaded: true,
				value: 7,
			},
		},
		{
			Actual: Something(8),
			Expected: Type{
				loaded: true,
				value: 8,
			},
		},
		{
			Actual: Something(9),
			Expected: Type{
				loaded: true,
				value: 9,
			},
		},
		{
			Actual: Something(10),
			Expected: Type{
				loaded: true,
				value: 10,
			},
		},
		{
			Actual: Something(11),
			Expected: Type{
				loaded: true,
				value: 11,
			},
		},
		{
			Actual: Something(12),
			Expected: Type{
				loaded: true,
				value: 12,
			},
		},
		{
			Actual: Something(13),
			Expected: Type{
				loaded: true,
				value: 13,
			},
		},
		{
			Actual: Something(14),
			Expected: Type{
				loaded: true,
				value: 14,
			},
		},
		{
			Actual: Something(15),
			Expected: Type{
				loaded: true,
				value: 15,
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
