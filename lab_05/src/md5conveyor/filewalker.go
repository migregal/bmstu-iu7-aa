package md5conveyor

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

func walkFiles(done <-chan struct{}, root string) (<-chan filewalkerOutput, <-chan error) {
	paths := make(chan filewalkerOutput)
	errc := make(chan error, 1)
	go func() {
		defer close(paths)

		start := time.Now()
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- filewalkerOutput{path, time.Since(start), time.Now()}:
				start = time.Now()
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}
