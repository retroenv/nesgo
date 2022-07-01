package gamegenie

var translationCharToValue = map[byte]byte{
	'A': 0x0,
	'P': 0x1,
	'Z': 0x2,
	'L': 0x3,
	'G': 0x4,
	'I': 0x5,
	'T': 0x6,
	'Y': 0x7,
	'E': 0x8,
	'O': 0x9,
	'X': 0xA,
	'U': 0xB,
	'K': 0xC,
	'S': 0xD,
	'V': 0xE,
	'N': 0xF,
}

var translationValueToChar = []byte{
	'A',
	'P',
	'Z',
	'L',
	'G',
	'I',
	'T',
	'Y',
	'E',
	'O',
	'X',
	'U',
	'K',
	'S',
	'V',
	'N',
}
