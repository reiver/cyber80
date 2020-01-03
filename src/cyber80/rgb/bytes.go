package cyber80_rgb

func Bytes(p []byte) Type {

	var r byte
	var g byte
	var b byte

	{
		length := len(p)

		if 1 <= length {
			r = p[0]

			if 2 <= length {
				g = p[1]

				if 3 <= length {
					b = p[2]
				}
			}
		}
	}

	return RGB(r, g, b)
}
