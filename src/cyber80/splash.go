package main

import (
	"github.com/reiver/cyber80/src/cyber80/ram"
	"github.com/reiver/cyber80/src/cyber80/raster"
)

func splash() {
	const top = 40
	const left = 16

	for y:=top;y<(top+64/2); y++ {
		for x:=left;x<(127-16);x++ {
			cyber80_ram.Value[cyber80_raster.Index(x,y)] = 0x0c
		}
	}

	face(left+4, top+9)
	man(left+23, top+18)
	botLeft(left+53, top+18)
	bot(left+63, top+18)
	bot(left+72, top+18)
	key(left+82, top+18)
}

func face(u int, v int) {

	//                                       0
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+0)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+0)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+0)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+0)] = 0x07
	//                                       6
	//                                       7

	//                                       0
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+1)] = 0x07
	//                                       7

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+2)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+2)] = 0x07
	//                                       2
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+2)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+2)] = 0x07
	//                                       5
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+2)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+2)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+3)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+4)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+5)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+5)] = 0x07
	//                                       2
	//                                       3
	//                                       4
	//                                       5
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+5)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+5)] = 0x07

	//                                       0
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+6)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+6)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+6)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+6)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+6)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+6)] = 0x07
	//                                       7

	//                                       0
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+7)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+7)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+7)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+7)] = 0x07
	//                                       6
	//                                       7

}

func bot(u int, v int) {
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+0)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+2, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+1)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+1, v+2)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+2)] = 0x07
	//                                       3
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+2)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+2)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+3)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+4)] = 0x07
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+4)] = 0x07
	//                                       5
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+4)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+5)] = 0x07
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+5)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+5)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+5)] = 0x07
	//                                       5
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+5)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+2, v+6)] = 0x07
	//                                       3
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+6)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+2, v+7)] = 0x07
	//                                       3
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+7)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+1, v+8)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+8)] = 0x07
	//                                       3
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+8)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+8)] = 0x07
}

func botLeft(u int, v int) {
	//                                       0
	//                                       1
	//                                       2
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+0)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+0)] = 0x07
	//                                       5
	//                                       6
	//                                       7

	//                                       0
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+1)] = 0x07
	//                                       6
	//                                       7

	//                                       0
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+2)] = 0x07
	//                                       2
	//                                       3
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+2)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+2)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+2)] = 0x07
	//                                       7

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+3)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+4)] = 0x07
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+4)] = 0x07
	//                                       6
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+4)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+5)] = 0x07
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+5)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+5)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+5)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+5)] = 0x07
	//                                       6
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+5)] = 0x07

	//                                       0
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+6)] = 0x07
	//                                       3
	//                                       4
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+6)] = 0x07
	//                                       6
	//                                       7

	//                                       0
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+7)] = 0x07
	//                                       3
	//                                       4
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+7)] = 0x07
	//                                       6
	//                                       7

	//                                       0
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+8)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+8)] = 0x07
	//                                       3
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+8)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+8)] = 0x07
	//                                       6
	//                                       7
}

func key(u int, v int) {

	cyber80_ram.Value[cyber80_raster.Index(u+5, v+6)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+6)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+6)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+7)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+7)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+7)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+7)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+7)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+7)] = 0x07
	//                                       6
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+7)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+8)] = 0x07
	//                                       1
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+8)] = 0x07
	//                                       3
	//                                       4
	cyber80_ram.Value[cyber80_raster.Index(u+5, v+8)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+6, v+8)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+8)] = 0x07
}

func man(u int, v int) {

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+0)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+0)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+1)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+1)] = 0x07

	//                                            2

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+3, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+4, v+3)] = 0x07
	//                                       5
	//                                       6
	cyber80_ram.Value[cyber80_raster.Index(u+7, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+8, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+9, v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+10,v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+11,v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+12,v+3)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+13,v+3)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+4)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+4)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+5)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+5)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+6)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+6)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+7)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+7)] = 0x07

	cyber80_ram.Value[cyber80_raster.Index(u+0, v+8)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+1, v+8)] = 0x07
	cyber80_ram.Value[cyber80_raster.Index(u+2, v+8)] = 0x07

}
