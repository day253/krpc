package asyncwriter

import (
	"bytes"
	"context"
	"log"
	"sync"
	"time"

	"github.com/day253/krpc/bufferpool"
	"go.uber.org/zap/zapcore"
)

type AsyncWriterSyncer struct {
	ws      zapcore.WriteSyncer
	wg      sync.WaitGroup
	pool    *bufferpool.BytesBufferPool
	bufChan chan *bytes.Buffer
	cancel  context.CancelFunc
}

func (s *AsyncWriterSyncer) Stop() error {
	s.wg.Wait()
	s.cancel()
	s.cleanup()
	return s.Sync()
}

func (s *AsyncWriterSyncer) cleanup() {
	timeout := 100 * time.Millisecond
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	var buf *bytes.Buffer
	var ok bool
	for {
		select {
		case buf, ok = <-s.bufChan:
			if !ok {
				return
			}
			_, _ = s.write(buf)
		case <-timer.C:
			return
		}
		timer.Reset(timeout)
	}
}

func (s *AsyncWriterSyncer) consume(ctx context.Context) {
	var readBuf *bytes.Buffer
	var ok bool
	for {
		select {
		case readBuf, ok = <-s.bufChan:
			if !ok {
				log.Println("log channel 0 close, async consume exit")
				return
			}
			_, _ = s.write(readBuf)
		case <-ctx.Done():
			log.Println("log context done, async consume exit")
			return
		}
	}
}

// Write取一个对象，然后write时候归还
func (s *AsyncWriterSyncer) Write(bs []byte) (int, error) {
	writeBuf := s.pool.Get()
	writeBuf.Write(bs)
	select {
	case s.bufChan <- writeBuf:
	default:
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			// TODO 阻塞久了可能因为写日志OOM
			_, _ = s.write(writeBuf)
		}()
	}
	return len(bs), nil
}

func (s *AsyncWriterSyncer) write(bsBuf *bytes.Buffer) (int, error) {
	buf := bsBuf.Bytes()
	defer s.pool.Put(bsBuf)
	if len(buf) == 0 {
		return 0, nil
	}
	n, err := s.ws.Write(buf)
	return n, err
}

func (s *AsyncWriterSyncer) Sync() error {
	return s.ws.Sync()
}

func NewDefaultAsyncBufferWriteSyncer(ws zapcore.WriteSyncer) *AsyncWriterSyncer {
	return NewAsyncBufferWriteSyncer(ws, 0, 0, 0)
}

// Buffer wraps a WriteSyncer in a buffer to improve performance,
// if bufferSize = 0, we set it to defaultBufferSize
// if flushInterval = 0, we set it to defaultFlushInterval
func NewAsyncBufferWriteSyncer(ws zapcore.WriteSyncer, bufferSize int, asyncBufferSize int, flushInterval time.Duration) *AsyncWriterSyncer {
	if bufferSize <= 0 {
		bufferSize = defaultBufferSize
	}

	if asyncBufferSize <= 0 {
		asyncBufferSize = defaultAsyncBufferSize
	}

	if flushInterval == 0 {
		flushInterval = defaultFlushInterval
	}

	ctx, cancel := context.WithCancel(context.Background())
	l := &AsyncWriterSyncer{
		ws:      newBufferWriterSyncer(ws, bufferSize),
		pool:    bufferpool.NewBytesBufferPool(1024),
		bufChan: make(chan *bytes.Buffer, asyncBufferSize),
		cancel:  cancel,
	}

	go l.consume(ctx)

	// flush buffer every interval
	// 不需要同步但是要取消
	go func() {
		ticker := time.NewTicker(flushInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := l.Sync(); err != nil {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return l
}
