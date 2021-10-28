package main

import (
	"fmt"
	"lab_05/md5conveyor"
)

func process(data []md5conveyor.Output, realTime int64) {
	var (
		total int64 = 0

		filewalker    int64
		digesterQueue int64
		digester      int64
		queue         int64
	)

	for _, res := range data {
		total += 1

		filewalker += res.Filewalker.Milliseconds()
		digesterQueue += res.DigesterQueue.Milliseconds()
		digester += res.Digester.Milliseconds()
		queue += res.Queue.Milliseconds()
	}

	fmt.Println("Total files:            ", total)
	fmt.Println("Real time spent:        ", realTime)
	fmt.Println("Total time spent:       ", filewalker+digesterQueue+digester+queue)

	fmt.Println("Avg filewarker time:    ", float64(filewalker)/float64(total))
	fmt.Println("Avg digester queue time:", float64(digesterQueue)/float64(total))
	fmt.Println("Avg digester time:      ", float64(digester)/float64(total))
	fmt.Println("Avg queue time:         ", float64(queue)/float64(total))
}
