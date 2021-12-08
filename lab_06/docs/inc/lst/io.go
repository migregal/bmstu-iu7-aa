package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetWeights(fn string) [][]int {
	w := make([][]int, 0)
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		str, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = strings.TrimSuffix(str, "\n")
		str = strings.TrimSuffix(str, "\r")
		str = strings.TrimRight(str, " ")
		cur := strings.Split(str, " ")

		path := make([]int, 0)
		for _, i := range cur {
			i, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println(err)
			}
			path = append(path, i)
		}
		w = append(w, path)
	}

	return w
}
