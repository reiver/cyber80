package cyber80_palette_color

func (receiver Type) Unwrap() (byte, error) {

	if ! receiver.loaded {
		return 0, errNotLoaded
	}
	if receiver.errored {
		return 0, errError
	}

	return receiver.value, nil
}
