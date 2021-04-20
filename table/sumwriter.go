package table

import "io"

// sumWriter is a wrapper around an io.Writer which keeps a running sum of bytes written.
type sumWriter struct {
	Sum int
	Dst io.Writer
}

func newSumWriter(w io.Writer) *sumWriter {
	return &sumWriter{
		Sum: 0,
		Dst: w,
	}
}

func (w *sumWriter) Write(buf []byte) (int, error) {
	count, err := w.Dst.Write(buf)
	w.Sum += count
	return count, err
}
