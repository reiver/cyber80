package cyber80_ram_test

import (
	"github.com/reiver/cyber80/src/cyber80/ram"
	"github.com/reiver/cyber80/src/cyber80/rgb"

	"reflect"

	"testing"
)

func TestTypePaletteNilReceiver(t *testing.T) {

	var receiver *cyber80_ram.Type

	palette := receiver.Palette()

	if expected, actual := cyber80_rgb.RGB(0,0,0), palette.RGB(0); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Did not get expected color.")
		t.Logf("EXPECTED: %z", expected)
		t.Logf("ACTUAL:   %z", actual)
		return
	}
	if expected, actual := cyber80_rgb.RGB(0,0,127), palette.RGB(1); !reflect.DeepEqual(expected, actual) {
		t.Errorf("Did not get expected color.")
		t.Logf("EXPECTED: %z", expected)
		t.Logf("ACTUAL:   %z", actual)
		return
	}
}
