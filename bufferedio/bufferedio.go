package bufferedio

import (
	"errors"
	"io"
)

var (
	EofError = errors.New("end of file")
)

type BufWriter struct {
	buf     []int
	wrapped io.Writer
}

func New(cap int, wrapped io.Writer) *BufWriter {
	return &BufWriter{
		buf:     make([]int, cap),
		wrapped: wrapped,
	}
}

func (b *BufWriter) Read(p []byte) (n int, err error) {
	return -1, EofError
}
