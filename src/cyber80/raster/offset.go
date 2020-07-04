package cyber80_raster

func Offset(x int, y int, bp int) int {
	offset := (Width * y) + x + bp

	return offset
}
