package graph

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func genData(fn string, n int) {
	os.Remove(fn)
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				str := fmt.Sprintf("%d ", rand.Intn(10)+1)
				f.WriteString(str)
			} else {
				str := fmt.Sprintf("%d ", 0)
				f.WriteString(str)
			}
		}
		f.WriteString("\n")
	}
}
