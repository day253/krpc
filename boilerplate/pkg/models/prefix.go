package models

import (
	"context"
)

type ObjectPrefix struct {
	VariablePrefix string
}

func (s *ObjectPrefix) Run(ctx context.Context, details map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	result[s.VariablePrefix] = details
	return result
}
