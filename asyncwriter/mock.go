package asyncwriter

import (
	"bytes"
	"sync"
)

// threadSafeBuffer 封装了 bytes.Buffer 和一个互斥锁
type threadSafeBuffer struct {
	buf bytes.Buffer
	mu  sync.Mutex
}

// Write 向缓冲区写入数据，确保线程安全
func (tsb *threadSafeBuffer) Write(p []byte) (n int, err error) {
	tsb.mu.Lock()
	defer tsb.mu.Unlock()
	return tsb.buf.Write(p)
}

// Read 从缓冲区读取数据，确保线程安全
func (tsb *threadSafeBuffer) Read(p []byte) (n int, err error) {
	tsb.mu.Lock()
	defer tsb.mu.Unlock()
	return tsb.buf.Read(p)
}

// Bytes 返回缓冲区的一个副本，确保线程安全
func (tsb *threadSafeBuffer) Bytes() []byte {
	tsb.mu.Lock()
	defer tsb.mu.Unlock()
	return tsb.buf.Bytes()
}

// String 返回缓冲区内容的字符串形式，确保线程安全
func (tsb *threadSafeBuffer) String() string {
	tsb.mu.Lock()
	defer tsb.mu.Unlock()
	return tsb.buf.String()
}

// Len 返回缓冲区中的字节数，确保线程安全
func (tsb *threadSafeBuffer) Len() int {
	tsb.mu.Lock()
	defer tsb.mu.Unlock()
	return tsb.buf.Len()
}

// Cap 返回缓冲区的容量，确保线程安全
func (tsb *threadSafeBuffer) Cap() int {
	tsb.mu.Lock()
	defer tsb.mu.Unlock()
	return tsb.buf.Cap()
}

// Reset 重置缓冲区，确保线程安全
func (tsb *threadSafeBuffer) Reset() {
	tsb.mu.Lock()
	defer tsb.mu.Unlock()
	tsb.buf.Reset()
}

// Grow 增加缓冲区的容量，确保线程安全
func (tsb *threadSafeBuffer) Grow(n int) {
	tsb.mu.Lock()
	defer tsb.mu.Unlock()
	tsb.buf.Grow(n)
}

// 模拟 io.Writer 来检测写入行为
type mockWriter struct {
	buf bytes.Buffer
	err error
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	if m.err != nil {
		return 0, m.err
	}
	return m.buf.Write(p)
}
