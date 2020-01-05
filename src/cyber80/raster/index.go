package cyber80_raster

func Index(x int, y int) int {
	index := (Width * y) + x

	return index
}
