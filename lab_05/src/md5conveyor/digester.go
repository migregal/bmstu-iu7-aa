package md5conveyor

import (
	"crypto/md5"
	"io/ioutil"
	"time"
)

func digester(done <-chan struct{}, paths <-chan fwOutput, c chan<- dOutput) {
	for path := range paths {
		start := time.Now()
		data, err := ioutil.ReadFile(path.path)
		select {
		case c <- dOutput{
			path.path,
			md5.Sum(data),
			err,
			path.processTime,
			start.Sub(path.queueStart),
			time.Since(start),
			time.Now(),
		}:
			start = time.Now()
		case <-done:
			return
		}
	}
}
