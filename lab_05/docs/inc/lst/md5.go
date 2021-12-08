package md5conveyor

import (
	"sync"
	"time"
)

func (c *Conveyor) Md5All(root string) ([]Output, error) {
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles(done, root)

	out := make(chan dOutput)
	var wg sync.WaitGroup
	wg.Add(c.numDigesters)
	for i := 0; i < c.numDigesters; i++ {
		go func() {
			digester(done, paths, out)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	m := make([]Output, 0)
	for r := range out {
		if r.err != nil {
			return nil, r.err
		}
		m = append(m, Output{
			r.path,
			r.sum,
			r.filewalker,
			r.queue,
			r.processTime,
			time.Since(r.md5Start),
		})
	}

	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}
