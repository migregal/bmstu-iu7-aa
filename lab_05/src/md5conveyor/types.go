package md5conveyor

import (
	"crypto/md5"
	"time"
)

type fwOutput struct {
	path string

	processTime time.Duration
	queueStart  time.Time
}

type dOutput struct {
	path string
	sum  [md5.Size]byte
	err  error

	filewalker  time.Duration
	queue       time.Duration
	processTime time.Duration
	md5Start    time.Time
}

type Output struct {
	Path string
	Sum  [md5.Size]byte

	Filewalker    time.Duration
	DigesterQueue time.Duration
	Digester      time.Duration
	Queue         time.Duration
}
