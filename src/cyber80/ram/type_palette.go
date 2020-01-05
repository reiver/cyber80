package cyber80_ram

import (
	"github.com/reiver/cyber80/src/cyber80/palette"
	"github.com/reiver/cyber80/src/cyber80/raster"
)

func (receiver *Type) Palette() cyber80_palette.Type {
	if nil == receiver {
		return cyber80_palette.Type{
			0x00,0x00,0x00, // 0 — black
			0x00,0x00,0x7f, // 1 — dark-blue
			0x7f,0x00,0x7f, // 2 — dark-purple
			0x00,0x7f,0x00, // 3 — dark-green
			0x7f,0x3f,0x3f, // 4 — brown
			0x3f,0x3f,0x3f, // 5 — dark-gray
			0x7f,0x7f,0x7f, // 6 — light-gray
			0xff,0xff,0xff, // 7 — white
			0xff,0x00,0x00, // 8 — red
			0xff,0x7f,0x00, // 9 — orange
			0xff,0xff,0x00, // 10 — yellow
			0x00,0xff,0x00, // 11 — red
			0x00,0x00,0xff, // 12 — blue
		}
	}

	start := cyber80_raster.Size

	p := Value[start:]

	var buffer cyber80_palette.Type

	copy(buffer[:], p)

	return buffer
}
