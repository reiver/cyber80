package cyber80_raster

func XY(i int) (x int, y int) {
	x, y = i % Width, i / Width
	return
}
