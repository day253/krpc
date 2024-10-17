package asyncwriter

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBufferWriterSyncer_Write 测试正常写入
func TestBufferWriterSyncer_Write(t *testing.T) {
	assert := assert.New(t)
	w := &threadSafeBuffer{}
	size := 10
	syncer := newBufferWriterSyncer(w, size)

	data := "Hello"
	n, err := syncer.Write([]byte(data))
	assert.NoError(err, "Write should not return an error")
	assert.Equal(len(data), n, "Write should return correct byte count")

	assert.Empty(w.String(), "Buffer should be empty before flushing")

	err = syncer.Sync()
	assert.NoError(err, "Flush should not return an error")
	assert.Equal(data, w.String(), "Buffer should contain the written data after flushing")
}

// TestBufferWriterSyncer_BufferOverflow 测试缓冲区溢出处理
func TestBufferWriterSyncer_BufferOverflow(t *testing.T) {
	assert := assert.New(t)
	w := &threadSafeBuffer{}
	size := 5
	syncer := newBufferWriterSyncer(w, size)

	data := "HelloWorld"
	_, err := syncer.Write([]byte(data))
	assert.NoError(err, "Write should not return an error")
	assert.Equal(data, w.String(), "Buffer should contain the written data")
}

// TestBufferWriterSyncer_Sync 测试同步写入
func TestBufferWriterSyncer_Sync(t *testing.T) {
	assert := assert.New(t)
	w := &threadSafeBuffer{}
	syncer := newBufferWriterSyncer(w, 10)

	data := "Hello"
	_, _ = syncer.Write([]byte(data))

	err := syncer.Sync()
	assert.NoError(err, "Sync should not return an error")
	assert.Equal(data, w.String(), "Buffer should contain the written data after sync")
}

// TestBufferWriterSyncer_ConcurrentWrites 测试并发写入
func TestBufferWriterSyncer_ConcurrentWrites(t *testing.T) {
	assert := assert.New(t)
	w := &threadSafeBuffer{}
	syncer := newBufferWriterSyncer(w, 20)

	var wg sync.WaitGroup
	writeData := "Hello"
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = syncer.Write([]byte(writeData))
		}()
	}
	wg.Wait()

	err := syncer.Sync()
	assert.NoError(err, "Sync should not return an error")
	assert.Len(w.String(), 5*len(writeData), "Buffer should contain the correct amount of written data")
}

// TestBufferWriterSyncer_ErrorHandling 测试错误处理
func TestBufferWriterSyncer_ErrorHandling(t *testing.T) {
	assert := assert.New(t)
	mockErr := errors.New("mock error")
	mw := &mockWriter{err: mockErr}
	syncer := newBufferWriterSyncer(mw, 3)

	_, err := syncer.Write([]byte("Hello"))
	assert.Equal(mockErr, err, "Expected mock error to be returned")
}

// TestBufferWriterSyncer_ErrorHandling 测试错误处理
func TestBufferWriterSyncer_ErrorHandling2(t *testing.T) {
	assert := assert.New(t)
	mockErr := errors.New("mock error")
	mw := &mockWriter{err: mockErr}
	syncer := newBufferWriterSyncer(mw, 10)

	_, err := syncer.Write([]byte("Hello"))
	assert.NotEqual(mockErr, err, "Expected mock error to be returned")
}
