package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/zellyn/wordle/solver"
)

func main() {
	colors := []*color.Color{
		color.New(color.FgYellow, color.Bold),
		color.New(color.FgGreen, color.Bold),
		color.New(color.FgHiBlack, color.Bold),
	}
	for i, r := range []rune(solver.BlockLetters) {
		colors[i%len(colors)].Set()
		fmt.Printf("%c ", r)
	}
	color.Unset()
	fmt.Println()
	fmt.Println("hello")
}
