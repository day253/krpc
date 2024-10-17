package asyncwriter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func BenchmarkBufferedWriteSyncer(b *testing.B) {
	b.Run("write file with buffer", func(b *testing.B) {
		file, err := os.CreateTemp(b.TempDir(), "test.log")
		require.NoError(b, err)

		defer func() {
			assert.NoError(b, file.Close())
		}()

		w := NewDefaultAsyncBufferWriteSyncer(zapcore.AddSync(file))
		defer func() {
			assert.NoError(b, w.Stop(), "failed to stop buffered write syncer")
		}()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := w.Write([]byte("foobarbazbabble")); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
}
