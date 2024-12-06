package data

import (
	"embed"
	_ "embed"
	"strings"
)

/*
    2309 wordle-nyt-answers-alphabetical.txt
    2315 wordle-answers-alphabetical.txt
   10638 wordle-nyt-allowed-guesses.txt
   10657 wordle-allowed-guesses.txt
   12546 wordle-nyt-allowed-guesses-update-12546.txt
   14855 wordle-nyt-words-14855.txt
*/

//go:embed *.txt
var f embed.FS

var Original_Answers_Alphabetical []string
var Original_Allowed_Guesses []string
var NYT_Answers_Alphabetical []string
var NYT_Allowed_Guesses []string
var NYT_Allowed_Guesses_Update_12546 []string
var NYT_Words_14855 []string

func init() {
	vars := []*[]string{
		&Original_Answers_Alphabetical,
		&Original_Allowed_Guesses,
		&NYT_Answers_Alphabetical,
		&NYT_Allowed_Guesses,
		&NYT_Allowed_Guesses_Update_12546,
		&NYT_Words_14855,
	}
	filenames := []string{
		"wordle-answers-alphabetical.txt",
		"wordle-allowed-guesses.txt",
		"wordle-nyt-answers-alphabetical.txt",
		"wordle-nyt-allowed-guesses.txt",
		"wordle-nyt-allowed-guesses-update-12546.txt",
		"wordle-nyt-words-14855.txt",
	}

	for i, variable := range vars {
		data, _ := f.ReadFile(filenames[i])
		*variable = strings.Split(strings.TrimSpace(string(data)), "\n")
	}
}
