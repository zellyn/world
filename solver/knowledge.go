package solver

import (
	"fmt"
	"math/bits"
	"strings"
)

type knowledge [6]uint32

func fromWord(word string) knowledge {
	word = strings.ToLower(word)
	var k knowledge
	for i, r := range word {
		if r < 'a' || r > 'z' {
			panic(fmt.Sprintf("weird word: %q", word))
		}
		k[i] = 1 << (r - 'a')
		k[5] |= 1 << (r - 'a')
	}
	return k
}

var fresh knowledge

func init() {
	for i := range 5 {
		fresh[i] = 1<<26 - 1
	}
}

func (k knowledge) contains(other knowledge) bool {
	for i := range 5 {
		if k[i]|other[i] != k[i] {
			return false
		}
	}
	if k[5]|other[5] != other[5] {
		return false
	}
	return true
}

func (k knowledge) containsWord(word string) bool {
	return k.contains(fromWord(word))
}

func fromGuess(answer, guess knowledge) knowledge {
	var k knowledge

	for i := range 5 {
		g := guess[i]
		if g == answer[i] {
			k[i] = g
			k[5] |= g
			continue
		}
		k[i] = (1<<26 - 1) &^ g
		if g|answer[5] != 0 {
			k[5] |= g
		}
	}

	if bits.OnesCount32(k[5]) == 5 {
		for i := range 5 {
			k[i] &= k[5]
		}
	}

	return k
}

func fromGuessWords(answer, guess string) knowledge {
	return fromGuess(fromWord(answer), fromWord(guess))
}
