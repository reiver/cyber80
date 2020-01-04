package cyber80_display

func XY(i int) (x1 int, x2 int, y int) {
	xs := i % (Width/2)
	y  = i / (Width/2)

	x1 = xs*2
	x2 = x1 + 1

	return
}
