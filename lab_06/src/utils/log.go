package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

func LogTime(h string, a []time.Duration) {
	fmt.Printf("%20v\n", aurora.BgRed(h))
	fmt.Printf("%v%18v%v\n", "+", strings.Repeat("-", 18), "+")
	fmt.Printf("|%3v|%14v|\n", aurora.Green("N"), aurora.Green("Time"))
	fmt.Printf("%v%18v%v\n", "+", strings.Repeat("-", 18), "+")
	for i := 0; i < len(a); i++ {
		fmt.Printf("|%3v|%14v|\n", i+2, a[i])
	}
	fmt.Printf("%v%18v%v\n", "+", strings.Repeat("-", 18), "+")
}
