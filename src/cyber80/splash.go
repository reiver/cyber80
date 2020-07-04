package main

import (
	"github.com/reiver/cyber80/src/cyber80/ram"
	"github.com/reiver/cyber80/src/cyber80/raster"
)

// splash creates a (specific) splash-screen on the cyber80's raster display.
func splash() {
	const bp = cyber80_ram.BPRaster

	const top = 40
	const left = 16

	for y:=top;y<(top+64/2); y++ {
		for x:=left;x<(127-16);x++ {
			cyber80_ram.Value[cyber80_raster.Offset(x,y,bp)] = 0x0c
		}
	}

	face(left+4, top+9)
	man(left+23, top+18)
	botLeft(left+53, top+18)
	bot(left+63, top+18)
	bot(left+72, top+18)
	key(left+82, top+18)
}

// face is used by splash().
func face(u int, v int) {
	const bp = cyber80_ram.BPRaster

	//                                        0
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+0, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+0, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+0, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+0, bp)] = 0x07
	//                                        6
	//                                        7

	//                                        0
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+1, bp)] = 0x07
	//                                        7

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+2, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+2, bp)] = 0x07
	//                                        2
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+2, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+2, bp)] = 0x07
	//                                        5
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+2, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+2, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+3, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+4, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+5, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+5, bp)] = 0x07
	//                                        2
	//                                        3
	//                                        4
	//                                        5
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+5, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+5, bp)] = 0x07

	//                                        0
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+6, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+6, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+6, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+6, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+6, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+6, bp)] = 0x07
	//                                        7

	//                                        0
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+7, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+7, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+7, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+7, bp)] = 0x07
	//                                        6
	//                                        7
}

// bot is used by splash().
func bot(u int, v int) {
	const bp = cyber80_ram.BPRaster

	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+0, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+1, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+2, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+2, bp)] = 0x07
	//                                        3
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+2, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+2, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+3, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+4, bp)] = 0x07
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+4, bp)] = 0x07
	//                                        5
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+4, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+5, bp)] = 0x07
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+5, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+5, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+5, bp)] = 0x07
	//                                        5
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+5, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+6, bp)] = 0x07
	//                                        3
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+6, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+7, bp)] = 0x07
	//                                        3
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+7, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+8, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+8, bp)] = 0x07
	//                                        3
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+8, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+8, bp)] = 0x07
}

// botLeft is used by splash().
func botLeft(u int, v int) {
	const bp = cyber80_ram.BPRaster

	//                                        0
	//                                        1
	//                                        2
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+0, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+0, bp)] = 0x07
	//                                        5
	//                                        6
	//                                        7

	//                                        0
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+1, bp)] = 0x07
	//                                        6
	//                                        7

	//                                        0
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+2, bp)] = 0x07
	//                                        2
	//                                        3
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+2, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+2, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+2, bp)] = 0x07
	//                                        7

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+3, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+4, bp)] = 0x07
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+4, bp)] = 0x07
	//                                        6
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+4, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+5, bp)] = 0x07
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+5, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+5, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+5, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+5, bp)] = 0x07
	//                                        6
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+5, bp)] = 0x07

	//                                        0
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+6, bp)] = 0x07
	//                                        3
	//                                        4
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+6, bp)] = 0x07
	//                                        6
	//                                        7

	//                                        0
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+7, bp)] = 0x07
	//                                        3
	//                                        4
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+7, bp)] = 0x07
	//                                        6
	//                                        7

	//                                        0
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+8, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+8, bp)] = 0x07
	//                                        3
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+8, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+8, bp)] = 0x07
	//                                        6
	//                                        7
}

// key is used by splash().
func key(u int, v int) {
	const bp = cyber80_ram.BPRaster

	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+6, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+6, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+6, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+7, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+7, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+7, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+7, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+7, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+7, bp)] = 0x07
	//                                        6
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+7, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+8, bp)] = 0x07
	//                                        1
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+8, bp)] = 0x07
	//                                        3
	//                                        4
	cyber80_ram.Value[cyber80_raster.Offset(u+5, v+8, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+6, v+8, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+8, bp)] = 0x07
}

// man is used by splash().
func man(u int, v int) {
	const bp = cyber80_ram.BPRaster

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+0, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+0, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+1, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+1, bp)] = 0x07

	//                                             2

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+3, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+4, v+3, bp)] = 0x07
	//                                        5
	//                                        6
	cyber80_ram.Value[cyber80_raster.Offset(u+7, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+8, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+9, v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+10,v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+11,v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+12,v+3, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+13,v+3, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+4, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+4, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+5, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+5, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+6, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+6, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+7, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+7, bp)] = 0x07

	cyber80_ram.Value[cyber80_raster.Offset(u+0, v+8, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+1, v+8, bp)] = 0x07
	cyber80_ram.Value[cyber80_raster.Offset(u+2, v+8, bp)] = 0x07
}
