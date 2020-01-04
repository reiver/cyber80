package cyber80_ram

import (
	"github.com/reiver/cyber80/src/cyber80/display"
)

const (
	rasterDisplaySize = (cyber80_display.Width * cyber80_display.Height) / 2
	palleteSize       = 24 * 16
)

const (
	Size = rasterDisplaySize + palleteSize
)

