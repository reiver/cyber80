package cyber80_palette_color

func Something(value byte) Type {
	if value < 0 || 15 < value{
		return Error()
	}

	return Type{
		loaded: true,
		value: value,
	}
}
