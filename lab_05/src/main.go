package main

import (
	"fmt"
	"lab_05/md5conveyor"
	"os"
	"time"
)

func main() {
	conveyor := md5conveyor.NewConveyor(30)
	start := time.Now()
	data, err := conveyor.Md5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	process(data, time.Since(start).Nanoseconds())
}
