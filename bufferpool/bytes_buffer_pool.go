package bufferpool

import (
	"bytes"
	"sync"
)

type BytesBufferPool struct {
	p *sync.Pool
}

// NewPool constructs a new Pool.
func NewBytesBufferPool(bytesCap int) *BytesBufferPool {
	return &BytesBufferPool{p: &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, bytesCap))
		},
	}}
}

// Get retrieves a Buffer from the pool, creating one if necessary.
func (this *BytesBufferPool) Get() *bytes.Buffer {
	buf := this.p.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func (this *BytesBufferPool) Put(buf *bytes.Buffer) {
	this.p.Put(buf)
}
