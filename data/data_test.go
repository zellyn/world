package data

import "testing"

func TestCounts(t *testing.T) {
	testdata := []struct {
		filename string
		data     []string
		want     int
	}{

		{
			filename: "wordle-answers-alphabetical.txt",
			data:     Original_Answers_Alphabetical,
			want:     2315,
		},
		{
			filename: "wordle-allowed-guesses.txt",
			data:     Original_Allowed_Guesses,
			want:     10657,
		},
		{
			filename: "wordle-nyt-answers-alphabetical.txt",
			data:     NYT_Answers_Alphabetical,
			want:     2309,
		},
		{
			filename: "wordle-nyt-allowed-guesses.txt",
			data:     NYT_Allowed_Guesses,
			want:     10638,
		},
		{
			filename: "wordle-nyt-allowed-guesses-update-12546.txt",
			data:     NYT_Allowed_Guesses_Update_12546,
			want:     12546,
		},
		{
			filename: "wordle-nyt-words-14855.txt",
			data:     NYT_Words_14855,
			want:     14855,
		},
	}

	for _, tt := range testdata {
		if got := len(tt.data); got != tt.want {
			t.Errorf("Want %q to have %d lines; got %d", tt.filename, tt.want, got)
		}
	}
}
