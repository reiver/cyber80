package cyber80_palette

import (
	"github.com/reiver/cyber80/src/cyber80/rgb"
)

func (receiver Type) RGB(index int) cyber80_rgb.Type {
	i := index * colorSize
	if len(receiver) <= i {
		return cyber80_rgb.Type{}
	}

	p := receiver[i:]

	return cyber80_rgb.Bytes(p)
}
