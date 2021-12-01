package main

import (
	"fmt"
	"sort"

	"github.com/logrusorgru/aurora"

	"lab_07/dict"
)

func printRes(method string, res dict.Dict) {
	fmt.Printf("%v ", aurora.Yellow(method+":"))
	if res["username"] == nil {
		fmt.Printf("%v\n\n", aurora.Red("There is no dict such key"))
	} else {
		fmt.Printf("%v\n", aurora.Green("Dict with specified key is found"))
		fmt.Println(res.ToString())
	}
}

func main() {
	fmt.Printf("%v", aurora.Magenta("Search in array of dict's\n\n"))

	darr := dict.CreateArray(10)
	farr := darr.FreqAnalysis()

	fmt.Printf("%v\n\n", aurora.Yellow("Dictionary data:"))
	darr.Print()

	var gt string
	fmt.Printf("%v ", aurora.Green("Key to find:"))
	fmt.Scanln(&gt)
	fmt.Println()

	printRes("Brute algorythm", darr.Brute(gt))

	sort.Slice(darr, func(i, j int) bool {
		return darr[i]["username"].(string) < darr[j]["username"].(string)
	})
	printRes("Binary search algorythm", darr.Binary(gt))

	printRes("Freq Analysis algorythm", farr.Combined(gt))
}
