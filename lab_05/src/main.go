package main

import (
	"fmt"
	"lab_05/md5conveyor"
	"os"
	"sort"
)

func main() {
	conveyor := md5conveyor.NewConveyor(40)
	m, err := conveyor.Md5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
