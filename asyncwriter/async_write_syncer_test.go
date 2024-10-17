package asyncwriter

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func requireWriteWorks(t testing.TB, ws zapcore.WriteSyncer) {
	n, err := ws.Write([]byte("foo"))
	require.NoError(t, err, "Unexpected error writing to WriteSyncer.")
	require.Equal(t, 3, n, "Wrote an unexpected number of bytes.")
}

func TestAsyncWriterSyncer(t *testing.T) {
	// If we pass a plain io.Writer, make sure that we still get a WriteSyncer
	// with a no-op Sync.
	buf := &threadSafeBuffer{}
	ws := NewDefaultAsyncBufferWriteSyncer(zapcore.AddSync(buf))

	requireWriteWorks(t, ws)
	assert.Empty(t, buf.String(), "Unexpected log calling a no-op Write method.")

	assert.NoError(t, ws.Sync(), "Unexpected error calling a no-op Sync method.")
	// assert.Equal(t, "foo", buf.String(), "Unexpected log string")

	assert.NoError(t, ws.Stop())
	assert.Equal(t, "foo", buf.String(), "Unexpected log string")
}

func TestAsyncWriterSyncer_Concurrency(t *testing.T) {
	buf := &threadSafeBuffer{}
	asyncWriter := NewDefaultAsyncBufferWriteSyncer(zapcore.AddSync(buf))

	var wg sync.WaitGroup
	testData := []byte("concurrent data")

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := asyncWriter.Write(testData)
			assert.NoError(t, err)
		}()
	}

	wg.Wait()

	err := asyncWriter.Sync()
	require.NoError(t, err, "Failed to sync data")

	assert.LessOrEqual(t, len(buf.String()), 100*len(testData), "Data written not equal to data received")

	err = asyncWriter.Stop()
	require.NoError(t, err, "Failed to stop AsyncWriterSyncer")

	assert.Equal(t, len(buf.String()), 100*len(testData), "Data written not equal to data received")
}

func TestAsyncWriterSyncer_StopTwicefunc(t *testing.T) {
	buf := &threadSafeBuffer{}
	ws := NewDefaultAsyncBufferWriteSyncer(zapcore.AddSync(buf))
	requireWriteWorks(t, ws)
	assert.Empty(t, buf.String(), "Unexpected log calling a no-op Write method.")
	assert.NoError(t, ws.Stop())
	assert.Equal(t, "foo", buf.String(), "Unexpected log string")
	assert.NoError(t, ws.Stop())
	assert.Equal(t, "foo", buf.String(), "Unexpected log string")
}
