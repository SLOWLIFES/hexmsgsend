package hexmsgsend

type Transmitter struct {
	CheckSum func(instruction []byte, isAppend bool) []byte
}

func (t *Transmitter) Data(data []byte) []byte {
	sendRaw := t.CheckSum(data, true)
	return sendRaw
}
