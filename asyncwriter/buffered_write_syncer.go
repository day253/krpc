package asyncwriter

import (
	"bufio"
	"io"
	"sync"
)

type bufferWriterSyncer struct {
	w *bufio.Writer
	sync.Mutex
}

func newBufferWriterSyncer(w io.Writer, size int) *bufferWriterSyncer {
	return &bufferWriterSyncer{
		w: bufio.NewWriterSize(w, size),
	}
}

func (s *bufferWriterSyncer) Write(bs []byte) (int, error) {
	s.Lock()
	defer s.Unlock()

	// there are some logic internal for bufio.Writer here:
	// 1. when the buffer is enough, data would not be flushed.
	// 2. when the buffer is not enough, data would be flushed as soon as the buffer fills up.
	// this would lead to log spliting, which is not acceptable for log collector
	// so we need to flush bufferWriter before writing the data into bufferWriter
	if len(bs) > s.w.Available() && s.w.Buffered() > 0 {
		if err := s.w.Flush(); err != nil {
			return 0, err
		}
	}

	return s.w.Write(bs)
}

func (s *bufferWriterSyncer) Sync() error {
	s.Lock()
	defer s.Unlock()

	return s.w.Flush()
}
