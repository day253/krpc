package objects

import (
	json "github.com/bytedance/sonic"
)

func String[T any](t T) string {
	return string(Bytes(t))
}

func Bytes[T any](t T) []byte {
	result, _ := json.Marshal(t)
	return result
}

func StringIndent[T any](t T) string {
	return string(BytesIndent(t))
}

func BytesIndent[T any](t T) []byte {
	result, _ := json.MarshalIndent(t, "", "\t")
	return result
}
