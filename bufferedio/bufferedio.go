package bufferedio

import (
	"io"
)

type BufWriter struct {
	cap     int
	buf     []byte
	wrapped io.Writer
	pos     int
}

func New(wrapped io.Writer, cap int) *BufWriter {
	return &BufWriter{
		cap:     cap,
		buf:     make([]byte, cap),
		wrapped: wrapped,
	}
}

func (bw *BufWriter) Write(data []byte) {
	for _, b := range data {
		bw.buf[bw.pos] = b
		bw.pos++
		if bw.pos == bw.cap {
			bw.Flush()
		}
	}
}

func (bw *BufWriter) Flush() {
	if bw.pos > 0 {
		bw.wrapped.Write(bw.buf)
		bw.pos = 0
		bw.buf = make([]byte, bw.cap)
	}
}
