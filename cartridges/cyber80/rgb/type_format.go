package cyber80_rgb

import (
	"fmt"
)

func (receiver Type) Format(f fmt.State, c rune) {
	switch c {
	case 'x':
		fmt.Fprintf(f, "%02x%02x%02x", receiver.value[r], receiver.value[g], receiver.value[b])
	case 'X':
		fmt.Fprintf(f, "%02X%02X%02X", receiver.value[r], receiver.value[g], receiver.value[b])
	default:
		fmt.Fprintf(f, "rgb(%d,%d,%d)", receiver.value[r], receiver.value[g], receiver.value[b])
	}
}
