package cyber80_rgb

func Bytes(p []byte) Type {

	var datum Type

	copy(datum.value[:], p)

	return datum
}
