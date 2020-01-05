package cyber80_raster

func Index(x int, y int, bp int) int {
	index := (Width * y) + x + bp

	return index
}
