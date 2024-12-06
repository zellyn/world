package solver

const blockLetterStart = 127344

var BlockLetters string

func init() {
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		BlockLetters += string(blockLetter(r))
	}
}

func blockLetter(r rune) rune {
	if r < 'a' || r > 'z' {
		return '?'
	}
	return (r - 'a') + blockLetterStart
}
