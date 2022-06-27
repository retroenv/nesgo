package cartridge

func highNibble(b byte) byte {
	return b >> 4
}

func mergeNibbles(highNibble byte, lowNibble byte) byte {
	return highNibble<<4 | lowNibble
}

// ControlBytes returns the 2 control bytes of the iNES header based on the cartridge configuration.
func ControlBytes(battery, mirror, mapper byte, hasTrainer bool) (byte, byte) {
	var control1, control2 byte
	control1 |= (battery & 1) << 1

	control1 |= mirror & 1
	control1 |= ((mirror >> 1) & 1) << 3

	control1 |= mergeNibbles(mapper, control1)
	control2 |= mergeNibbles(highNibble(mapper), control2)

	if hasTrainer {
		control1 |= trainerFlag
	}
	return control1, control2
}
