package cyber80_ram

func (receiver *Type) Get(i int) (byte, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	value := receiver[i]

	return value, nil
}
