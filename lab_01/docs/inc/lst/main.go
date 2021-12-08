package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"lab_01/levenstein"
)

func main() {
	fmt.Println(aurora.Magenta("Levinstein distance\n"))

	fst, snd := getWords()

	printMethodRes("Recursive method",
		levenshtein.NewLevenstein(fst, snd).Recursive())

	printMethodRes("Recursive Damerau-Levenstein method",
		levenshtein.NewLevenstein(fst, snd).DamerauLevenshteinRec())

	printMethodRes("Iterative method with cash",
		levenshtein.NewLevenstein(fst, snd).IterativeCache())

	printMethodRes("Damerau-Levenstein with cash",
		levenshtein.NewLevenstein(fst, snd).DamerauLevenshteinCache())
}
