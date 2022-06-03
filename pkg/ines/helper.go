package ines

func highNibble(b byte) byte {
	return b >> 4
}

func mergeNibbles(highNibble byte, lowNibble byte) byte {
	return highNibble<<4 | lowNibble
}
