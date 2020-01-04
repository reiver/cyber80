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

func TestSomethingError(t *testing.T) {

	tests := []struct{
		Actual Type
	}{
		{
			Actual: Something(16),
		},
		{
			Actual: Something(17),
		},
		{
			Actual: Something(18),
		},
		{
			Actual: Something(19),
		},
		{
			Actual: Something(20),
		},
		{
			Actual: Something(21),
		},
		{
			Actual: Something(22),
		},
		{
			Actual: Something(23),
		},
		{
			Actual: Something(24),
		},
		{
			Actual: Something(25),
		},
		{
			Actual: Something(26),
		},
		{
			Actual: Something(27),
		},
		{
			Actual: Something(28),
		},
		{
			Actual: Something(29),
		},
		{
			Actual: Something(30),
		},
		{
			Actual: Something(31),
		},
		{
			Actual: Something(32),
		},
		{
			Actual: Something(33),
		},
		{
			Actual: Something(34),
		},
		{
			Actual: Something(35),
		},
		{
			Actual: Something(36),
		},
		{
			Actual: Something(37),
		},
		{
			Actual: Something(38),
		},
		{
			Actual: Something(39),
		},
		{
			Actual: Something(40),
		},
		{
			Actual: Something(41),
		},
		{
			Actual: Something(42),
		},
		{
			Actual: Something(43),
		},
		{
			Actual: Something(44),
		},
		{
			Actual: Something(45),
		},
		{
			Actual: Something(46),
		},
		{
			Actual: Something(47),
		},
		{
			Actual: Something(48),
		},
		{
			Actual: Something(49),
		},
		{
			Actual: Something(50),
		},
		{
			Actual: Something(51),
		},
		{
			Actual: Something(52),
		},
		{
			Actual: Something(53),
		},
		{
			Actual: Something(54),
		},
		{
			Actual: Something(55),
		},
		{
			Actual: Something(56),
		},
		{
			Actual: Something(57),
		},
		{
			Actual: Something(58),
		},
		{
			Actual: Something(59),
		},
		{
			Actual: Something(60),
		},
		{
			Actual: Something(61),
		},
		{
			Actual: Something(62),
		},
		{
			Actual: Something(63),
		},
		{
			Actual: Something(64),
		},
		{
			Actual: Something(65),
		},
		{
			Actual: Something(66),
		},
		{
			Actual: Something(67),
		},
		{
			Actual: Something(68),
		},
		{
			Actual: Something(69),
		},
		{
			Actual: Something(70),
		},
		{
			Actual: Something(71),
		},
		{
			Actual: Something(72),
		},
		{
			Actual: Something(73),
		},
		{
			Actual: Something(74),
		},
		{
			Actual: Something(75),
		},
		{
			Actual: Something(76),
		},
		{
			Actual: Something(77),
		},
		{
			Actual: Something(78),
		},
		{
			Actual: Something(79),
		},
		{
			Actual: Something(80),
		},
		{
			Actual: Something(81),
		},
		{
			Actual: Something(82),
		},
		{
			Actual: Something(83),
		},
		{
			Actual: Something(84),
		},
		{
			Actual: Something(85),
		},
		{
			Actual: Something(86),
		},
		{
			Actual: Something(87),
		},
		{
			Actual: Something(88),
		},
		{
			Actual: Something(89),
		},
		{
			Actual: Something(90),
		},
		{
			Actual: Something(91),
		},
		{
			Actual: Something(92),
		},
		{
			Actual: Something(93),
		},
		{
			Actual: Something(94),
		},
		{
			Actual: Something(95),
		},
		{
			Actual: Something(96),
		},
		{
			Actual: Something(97),
		},
		{
			Actual: Something(98),
		},
		{
			Actual: Something(99),
		},
		{
			Actual: Something(100),
		},
	}

	for testNumber, test := range tests {

		errored := Type{
			loaded: true,
			errored: true,
		}

		if expected, actual := errored, test.Actual; expected != actual {
			t.Errorf("For test #%d, what was expected was not what was actually gotten.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			continue
		}

	}
}
