package deepcopy

import (
	"reflect"
	"testing"

	"github.com/huandu/go-clone"
	"github.com/stretchr/testify/assert"
)

type Basics struct {
	String      string
	Strings     []string
	StringArr   [4]string
	Bool        bool
	Bools       []bool
	Byte        byte
	Bytes       []byte
	Int         int
	Ints        []int
	Int8        int8
	Int8s       []int8
	Int16       int16
	Int16s      []int16
	Int32       int32
	Int32s      []int32
	Int64       int64
	Int64s      []int64
	Uint        uint
	Uints       []uint
	Uint8       uint8
	Uint8s      []uint8
	Uint16      uint16
	Uint16s     []uint16
	Uint32      uint32
	Uint32s     []uint32
	Uint64      uint64
	Uint64s     []uint64
	Float32     float32
	Float32s    []float32
	Float64     float64
	Float64s    []float64
	Complex64   complex64
	Complex64s  []complex64
	Complex128  complex128
	Complex128s []complex128
	Interface   interface{}
	Interfaces  []interface{}
	Maps        map[string]interface{}
}

var src = Basics{
	String:      "kimchi",
	Strings:     []string{"uni", "ika"},
	StringArr:   [4]string{"malort", "barenjager", "fernet", "salmiakki"},
	Bool:        true,
	Bools:       []bool{true, false, true},
	Byte:        'z',
	Bytes:       []byte("abc"),
	Int:         42,
	Ints:        []int{0, 1, 3, 4},
	Int8:        8,
	Int8s:       []int8{8, 9, 10},
	Int16:       16,
	Int16s:      []int16{16, 17, 18, 19},
	Int32:       32,
	Int32s:      []int32{32, 33},
	Int64:       64,
	Int64s:      []int64{64},
	Uint:        420,
	Uints:       []uint{11, 12, 13},
	Uint8:       81,
	Uint8s:      []uint8{81, 82},
	Uint16:      160,
	Uint16s:     []uint16{160, 161, 162, 163, 164},
	Uint32:      320,
	Uint32s:     []uint32{320, 321},
	Uint64:      640,
	Uint64s:     []uint64{6400, 6401, 6402, 6403},
	Float32:     32.32,
	Float32s:    []float32{32.32, 33},
	Float64:     64.1,
	Float64s:    []float64{64, 65, 66},
	Complex64:   complex64(-64 + 12i),
	Complex64s:  []complex64{complex64(-65 + 11i), complex64(66 + 10i)},
	Complex128:  complex128(-128 + 12i),
	Complex128s: []complex128{complex128(-128 + 11i), complex128(129 + 10i)},
	Interfaces:  []interface{}{42, true, "pan-galactic"},
	Maps: map[string]interface{}{
		"1": map[string]interface{}{
			"2": map[string]interface{}{
				"3": map[string]interface{}{
					"4": map[string]interface{}{
						"5": map[string]interface{}{
							"6": map[string]interface{}{
								"7": "8",
							},
						},
					},
				},
			},
		},
	},
}

func Test_DeepCopy(t *testing.T) {
	dst := &Basics{}
	assert.Nil(t, copyByCopier(dst, &src))
	assert.Equal(t, true, reflect.DeepEqual(&src, dst))
}

func Benchmark_DeepCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst Basics
		err := Copy(&dst, &src)
		if err != nil {
			b.Error(err)
		}
	}
}

func Test_GoClone(t *testing.T) {
	v := clone.Clone(&src).(*Basics)
	assert.Equal(t, true, reflect.DeepEqual(&src, v))
}

func Benchmark_GoClone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst := clone.Clone(&src).(*Basics)
		if len(dst.Maps) == 0 {
			b.Error("reflect deep copy failed")
		}
	}
}
