package cyber80_ram

func (receiver *Type) Or(i int, value byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver[i] |= value

	return nil
}
