package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func getWords() (string, string) {
	fmt.Printf("Input first word: ")
	fst := readWord()
	fmt.Printf("Input second word: ")
	snd := readWord()
	fmt.Println()

	return fst, snd
}

func readWord() string {
	var word string
	fmt.Scanln(&word)

	return word
}

func printMethodRes(method string, dist int) {
	fmt.Printf("%s\nDistance: %d\n\n",
		aurora.Yellow(method),
		aurora.Green(dist))
}
