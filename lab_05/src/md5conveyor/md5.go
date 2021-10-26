package md5conveyor

import (
	"crypto/md5"
	"sync"
)

func (c *Conveyor) Md5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles(done, root)

	out := make(chan result)
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

	m := make(map[string][md5.Size]byte)
	for r := range out {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}
