//go:build go1.18
// +build go1.18

package objects

import (
	"encoding/json"

	"github.com/bytedance/sonic"
)

func String[T any](t T) string {
	return string(Bytes(t))
}

func Bytes[T any](t T) []byte {
	result, _ := sonic.Marshal(t)
	return result
}

func StringIndent[T any](t T) string {
	return string(BytesIndent(t))
}

func BytesIndent[T any](t T) []byte {
	result, _ := json.MarshalIndent(t, "", "\t")
	return result
}
