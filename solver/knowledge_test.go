package solver

import (
	"testing"

	"github.com/zellyn/wordle/data"
)

func TestContains(t *testing.T) {
	for _, word := range data.NYT_Words_14855 {
		if !fresh.contains(fromWord(word)) {
			t.Errorf("want fresh.contains(%q)==true", word)
		}
	}
}
