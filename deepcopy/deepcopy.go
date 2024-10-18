package deepcopy

import (
	"github.com/huandu/go-clone"
	"github.com/jinzhu/copier"
)

func Copy(toValue, fromValue interface{}) error {
	return copyByCopier(toValue, fromValue)
}

func copyByCopier(toValue, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{DeepCopy: true})
}

func Clone[T any](t T) T {
	return clone.Clone(t).(T)
}
