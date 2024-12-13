package models

import (
	"github.com/day253/krpc/objects"
	"github.com/wI2L/jsondiff"
)

func diff(source, target []byte) string {
	patch, _ := jsondiff.CompareJSON(source, target)
	return objects.StringIndent(patch)
}
