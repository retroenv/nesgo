// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 2
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 33: // ['!','!']
			return 3
		case r == 34: // ['"','"']
			return 4
		case r == 38: // ['&','&']
			return 5
		case r == 40: // ['(','(']
			return 6
		case r == 41: // [')',')']
			return 7
		case r == 42: // ['*','*']
			return 8
		case r == 43: // ['+','+']
			return 9
		case r == 44: // [',',',']
			return 10
		case r == 45: // ['-','-']
			return 11
		case r == 46: // ['.','.']
			return 12
		case r == 47: // ['/','/']
			return 13
		case r == 48: // ['0','0']
			return 14
		case 49 <= r && r <= 57: // ['1','9']
			return 15
		case r == 58: // [':',':']
			return 16
		case r == 59: // [';',';']
			return 17
		case r == 60: // ['<','<']
			return 18
		case r == 61: // ['=','=']
			return 19
		case r == 62: // ['>','>']
			return 20
		case 65 <= r && r <= 72: // ['A','H']
			return 21
		case r == 73: // ['I','I']
			return 22
		case 74 <= r && r <= 77: // ['J','M']
			return 21
		case r == 78: // ['N','N']
			return 23
		case 79 <= r && r <= 90: // ['O','Z']
			return 21
		case r == 94: // ['^','^']
			return 5
		case r == 95: // ['_','_']
			return 21
		case r == 96: // ['`','`']
			return 24
		case r == 97: // ['a','a']
			return 21
		case r == 98: // ['b','b']
			return 25
		case r == 99: // ['c','c']
			return 26
		case 100 <= r && r <= 101: // ['d','e']
			return 21
		case r == 102: // ['f','f']
			return 27
		case r == 103: // ['g','g']
			return 28
		case r == 104: // ['h','h']
			return 21
		case r == 105: // ['i','i']
			return 29
		case 106 <= r && r <= 111: // ['j','o']
			return 21
		case r == 112: // ['p','p']
			return 30
		case r == 113: // ['q','q']
			return 21
		case r == 114: // ['r','r']
			return 31
		case r == 115: // ['s','s']
			return 21
		case r == 116: // ['t','t']
			return 32
		case r == 117: // ['u','u']
			return 33
		case r == 118: // ['v','v']
			return 34
		case 119 <= r && r <= 122: // ['w','z']
			return 21
		case r == 123: // ['{','{']
			return 35
		case r == 124: // ['|','|']
			return 5
		case r == 125: // ['}','}']
			return 36
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 37
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 38
		case r == 92: // ['\','\']
			return 39
		default:
			return 4
		}
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 40
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 40
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 41
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 42
		case r == 47: // ['/','/']
			return 43
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 44
		case r == 66: // ['B','B']
			return 45
		case r == 88: // ['X','X']
			return 46
		case r == 98: // ['b','b']
			return 45
		case r == 120: // ['x','x']
			return 46
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 47
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 2
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 60: // ['<','<']
			return 5
		case r == 61: // ['=','=']
			return 37
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 37
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 37
		case r == 62: // ['>','>']
			return 5
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 109: // ['a','m']
			return 21
		case r == 110: // ['n','n']
			return 50
		case 111 <= r && r <= 122: // ['o','z']
			return 21
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 100: // ['a','d']
			return 21
		case r == 101: // ['e','e']
			return 51
		case 102 <= r && r <= 122: // ['f','z']
			return 21
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case r == 96: // ['`','`']
			return 52
		default:
			return 24
		}
	},
	// S25
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 113: // ['a','q']
			return 21
		case r == 114: // ['r','r']
			return 53
		case 115 <= r && r <= 122: // ['s','z']
			return 21
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 110: // ['a','n']
			return 21
		case r == 111: // ['o','o']
			return 54
		case 112 <= r && r <= 122: // ['p','z']
			return 21
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 110: // ['a','n']
			return 21
		case r == 111: // ['o','o']
			return 55
		case 112 <= r && r <= 116: // ['p','t']
			return 21
		case r == 117: // ['u','u']
			return 56
		case 118 <= r && r <= 122: // ['v','z']
			return 21
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 110: // ['a','n']
			return 21
		case r == 111: // ['o','o']
			return 57
		case 112 <= r && r <= 122: // ['p','z']
			return 21
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 101: // ['a','e']
			return 21
		case r == 102: // ['f','f']
			return 58
		case 103 <= r && r <= 108: // ['g','l']
			return 21
		case r == 109: // ['m','m']
			return 59
		case r == 110: // ['n','n']
			return 60
		case 111 <= r && r <= 122: // ['o','z']
			return 21
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case r == 97: // ['a','a']
			return 61
		case 98 <= r && r <= 122: // ['b','z']
			return 21
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 100: // ['a','d']
			return 21
		case r == 101: // ['e','e']
			return 62
		case 102 <= r && r <= 122: // ['f','z']
			return 21
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 120: // ['a','x']
			return 21
		case r == 121: // ['y','y']
			return 63
		case r == 122: // ['z','z']
			return 21
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 104: // ['a','h']
			return 21
		case r == 105: // ['i','i']
			return 64
		case 106 <= r && r <= 122: // ['j','z']
			return 21
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case r == 97: // ['a','a']
			return 65
		case 98 <= r && r <= 122: // ['b','z']
			return 21
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 4
		case r == 110: // ['n','n']
			return 66
		case r == 114: // ['r','r']
			return 66
		case r == 116: // ['t','t']
			return 66
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 67
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 68
		default:
			return 42
		}
	},
	// S43
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 69
		default:
			return 43
		}
	},
	// S44
	func(r rune) int {
		switch {
		case 48 <= r && r <= 55: // ['0','7']
			return 44
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		case 48 <= r && r <= 49: // ['0','1']
			return 70
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 71
		case 65 <= r && r <= 70: // ['A','F']
			return 72
		case 97 <= r && r <= 102: // ['a','f']
			return 72
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 47
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 107: // ['a','k']
			return 21
		case r == 108: // ['l','l']
			return 73
		case 109 <= r && r <= 122: // ['m','z']
			return 21
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 118: // ['a','v']
			return 21
		case r == 119: // ['w','w']
			return 74
		case 120 <= r && r <= 122: // ['x','z']
			return 21
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 100: // ['a','d']
			return 21
		case r == 101: // ['e','e']
			return 75
		case 102 <= r && r <= 122: // ['f','z']
			return 21
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 109: // ['a','m']
			return 21
		case r == 110: // ['n','n']
			return 76
		case 111 <= r && r <= 122: // ['o','z']
			return 21
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 113: // ['a','q']
			return 21
		case r == 114: // ['r','r']
			return 77
		case 115 <= r && r <= 122: // ['s','z']
			return 21
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 109: // ['a','m']
			return 21
		case r == 110: // ['n','n']
			return 78
		case 111 <= r && r <= 122: // ['o','z']
			return 21
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 115: // ['a','s']
			return 21
		case r == 116: // ['t','t']
			return 79
		case 117 <= r && r <= 122: // ['u','z']
			return 21
		}
		return NoState
	},
	// S58
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S59
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 111: // ['a','o']
			return 21
		case r == 112: // ['p','p']
			return 80
		case 113 <= r && r <= 122: // ['q','z']
			return 21
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 115: // ['a','s']
			return 21
		case r == 116: // ['t','t']
			return 81
		case 117 <= r && r <= 122: // ['u','z']
			return 21
		}
		return NoState
	},
	// S61
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 98: // ['a','b']
			return 21
		case r == 99: // ['c','c']
			return 82
		case 100 <= r && r <= 122: // ['d','z']
			return 21
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 115: // ['a','s']
			return 21
		case r == 116: // ['t','t']
			return 83
		case 117 <= r && r <= 122: // ['u','z']
			return 21
		}
		return NoState
	},
	// S63
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 111: // ['a','o']
			return 21
		case r == 112: // ['p','p']
			return 84
		case 113 <= r && r <= 122: // ['q','z']
			return 21
		}
		return NoState
	},
	// S64
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 109: // ['a','m']
			return 21
		case r == 110: // ['n','n']
			return 85
		case 111 <= r && r <= 122: // ['o','z']
			return 21
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 113: // ['a','q']
			return 21
		case r == 114: // ['r','r']
			return 86
		case 115 <= r && r <= 122: // ['s','z']
			return 21
		}
		return NoState
	},
	// S66
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 38
		case r == 92: // ['\','\']
			return 39
		default:
			return 4
		}
	},
	// S67
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 68
		case r == 47: // ['/','/']
			return 87
		default:
			return 42
		}
	},
	// S69
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S70
	func(r rune) int {
		switch {
		case 48 <= r && r <= 49: // ['0','1']
			return 70
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 71
		case 65 <= r && r <= 70: // ['A','F']
			return 72
		case 97 <= r && r <= 102: // ['a','f']
			return 72
		}
		return NoState
	},
	// S72
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 71
		case 65 <= r && r <= 70: // ['A','F']
			return 72
		case 97 <= r && r <= 102: // ['a','f']
			return 72
		}
		return NoState
	},
	// S73
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 104: // ['a','h']
			return 21
		case r == 105: // ['i','i']
			return 88
		case 106 <= r && r <= 122: // ['j','z']
			return 21
		}
		return NoState
	},
	// S74
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 72: // ['A','H']
			return 21
		case r == 73: // ['I','I']
			return 89
		case 74 <= r && r <= 84: // ['J','T']
			return 21
		case r == 85: // ['U','U']
			return 90
		case 86 <= r && r <= 90: // ['V','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S75
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case r == 97: // ['a','a']
			return 91
		case 98 <= r && r <= 122: // ['b','z']
			return 21
		}
		return NoState
	},
	// S76
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 114: // ['a','r']
			return 21
		case r == 115: // ['s','s']
			return 92
		case 116 <= r && r <= 122: // ['t','z']
			return 21
		}
		return NoState
	},
	// S77
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S78
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 98: // ['a','b']
			return 21
		case r == 99: // ['c','c']
			return 93
		case 100 <= r && r <= 122: // ['d','z']
			return 21
		}
		return NoState
	},
	// S79
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 110: // ['a','n']
			return 21
		case r == 111: // ['o','o']
			return 94
		case 112 <= r && r <= 122: // ['p','z']
			return 21
		}
		return NoState
	},
	// S80
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 110: // ['a','n']
			return 21
		case r == 111: // ['o','o']
			return 95
		case 112 <= r && r <= 122: // ['p','z']
			return 21
		}
		return NoState
	},
	// S81
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 55: // ['0','7']
			return 49
		case r == 56: // ['8','8']
			return 96
		case r == 57: // ['9','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 100: // ['a','d']
			return 21
		case r == 101: // ['e','e']
			return 97
		case 102 <= r && r <= 122: // ['f','z']
			return 21
		}
		return NoState
	},
	// S82
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 106: // ['a','j']
			return 21
		case r == 107: // ['k','k']
			return 98
		case 108 <= r && r <= 122: // ['l','z']
			return 21
		}
		return NoState
	},
	// S83
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 116: // ['a','t']
			return 21
		case r == 117: // ['u','u']
			return 99
		case 118 <= r && r <= 122: // ['v','z']
			return 21
		}
		return NoState
	},
	// S84
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 100: // ['a','d']
			return 21
		case r == 101: // ['e','e']
			return 100
		case 102 <= r && r <= 122: // ['f','z']
			return 21
		}
		return NoState
	},
	// S85
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 115: // ['a','s']
			return 21
		case r == 116: // ['t','t']
			return 101
		case 117 <= r && r <= 122: // ['u','z']
			return 21
		}
		return NoState
	},
	// S86
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S87
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S88
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 109: // ['a','m']
			return 21
		case r == 110: // ['n','n']
			return 102
		case 111 <= r && r <= 122: // ['o','z']
			return 21
		}
		return NoState
	},
	// S89
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 109: // ['a','m']
			return 21
		case r == 110: // ['n','n']
			return 103
		case 111 <= r && r <= 122: // ['o','z']
			return 21
		}
		return NoState
	},
	// S90
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 104: // ['a','h']
			return 21
		case r == 105: // ['i','i']
			return 104
		case 106 <= r && r <= 122: // ['j','z']
			return 21
		}
		return NoState
	},
	// S91
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 106: // ['a','j']
			return 21
		case r == 107: // ['k','k']
			return 105
		case 108 <= r && r <= 122: // ['l','z']
			return 21
		}
		return NoState
	},
	// S92
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 115: // ['a','s']
			return 21
		case r == 116: // ['t','t']
			return 106
		case 117 <= r && r <= 122: // ['u','z']
			return 21
		}
		return NoState
	},
	// S93
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S94
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S95
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 113: // ['a','q']
			return 21
		case r == 114: // ['r','r']
			return 107
		case 115 <= r && r <= 122: // ['s','z']
			return 21
		}
		return NoState
	},
	// S96
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S97
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 113: // ['a','q']
			return 21
		case r == 114: // ['r','r']
			return 108
		case 115 <= r && r <= 122: // ['s','z']
			return 21
		}
		return NoState
	},
	// S98
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case r == 97: // ['a','a']
			return 109
		case 98 <= r && r <= 122: // ['b','z']
			return 21
		}
		return NoState
	},
	// S99
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 113: // ['a','q']
			return 21
		case r == 114: // ['r','r']
			return 110
		case 115 <= r && r <= 122: // ['s','z']
			return 21
		}
		return NoState
	},
	// S100
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S101
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case r == 48: // ['0','0']
			return 49
		case r == 49: // ['1','1']
			return 111
		case 50 <= r && r <= 55: // ['2','7']
			return 49
		case r == 56: // ['8','8']
			return 96
		case r == 57: // ['9','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S102
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 100: // ['a','d']
			return 21
		case r == 101: // ['e','e']
			return 112
		case 102 <= r && r <= 122: // ['f','z']
			return 21
		}
		return NoState
	},
	// S103
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 115: // ['a','s']
			return 21
		case r == 116: // ['t','t']
			return 113
		case 117 <= r && r <= 122: // ['u','z']
			return 21
		}
		return NoState
	},
	// S104
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 109: // ['a','m']
			return 21
		case r == 110: // ['n','n']
			return 114
		case 111 <= r && r <= 122: // ['o','z']
			return 21
		}
		return NoState
	},
	// S105
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S106
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S107
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 115: // ['a','s']
			return 21
		case r == 116: // ['t','t']
			return 115
		case 117 <= r && r <= 122: // ['u','z']
			return 21
		}
		return NoState
	},
	// S108
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 101: // ['a','e']
			return 21
		case r == 102: // ['f','f']
			return 116
		case 103 <= r && r <= 122: // ['g','z']
			return 21
		}
		return NoState
	},
	// S109
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 102: // ['a','f']
			return 21
		case r == 103: // ['g','g']
			return 117
		case 104 <= r && r <= 122: // ['h','z']
			return 21
		}
		return NoState
	},
	// S110
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 109: // ['a','m']
			return 21
		case r == 110: // ['n','n']
			return 118
		case 111 <= r && r <= 122: // ['o','z']
			return 21
		}
		return NoState
	},
	// S111
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 53: // ['0','5']
			return 49
		case r == 54: // ['6','6']
			return 96
		case 55 <= r && r <= 57: // ['7','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S112
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S113
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 55: // ['0','7']
			return 49
		case r == 56: // ['8','8']
			return 119
		case r == 57: // ['9','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S114
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 115: // ['a','s']
			return 21
		case r == 116: // ['t','t']
			return 120
		case 117 <= r && r <= 122: // ['u','z']
			return 21
		}
		return NoState
	},
	// S115
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S116
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case r == 97: // ['a','a']
			return 121
		case 98 <= r && r <= 122: // ['b','z']
			return 21
		}
		return NoState
	},
	// S117
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 100: // ['a','d']
			return 21
		case r == 101: // ['e','e']
			return 122
		case 102 <= r && r <= 122: // ['f','z']
			return 21
		}
		return NoState
	},
	// S118
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S119
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S120
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case r == 48: // ['0','0']
			return 49
		case r == 49: // ['1','1']
			return 123
		case 50 <= r && r <= 55: // ['2','7']
			return 49
		case r == 56: // ['8','8']
			return 119
		case r == 57: // ['9','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S121
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 98: // ['a','b']
			return 21
		case r == 99: // ['c','c']
			return 124
		case 100 <= r && r <= 122: // ['d','z']
			return 21
		}
		return NoState
	},
	// S122
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S123
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 53: // ['0','5']
			return 49
		case r == 54: // ['6','6']
			return 119
		case 55 <= r && r <= 57: // ['7','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		}
		return NoState
	},
	// S124
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 100: // ['a','d']
			return 21
		case r == 101: // ['e','e']
			return 125
		case 102 <= r && r <= 122: // ['f','z']
			return 21
		}
		return NoState
	},
	// S125
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 48
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		case r == 123: // ['{','{']
			return 126
		}
		return NoState
	},
	// S126
	func(r rune) int {
		switch {
		case r == 125: // ['}','}']
			return 127
		}
		return NoState
	},
	// S127
	func(r rune) int {
		switch {
		}
		return NoState
	},
}
