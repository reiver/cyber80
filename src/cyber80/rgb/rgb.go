package cyber80_rgb

func RGB(red byte, green byte, blue byte) Type {

	var datum Type

	datum.value[r] = red
	datum.value[g] = green
	datum.value[b] = blue

	return datum
}
