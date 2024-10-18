package strategies

import (
	"github.com/ishumei/krpc/boilerplate/pkg/models"
	"github.com/ishumei/krpc/objects"
	"github.com/jinzhu/copier"
)

type FlowInput struct {
	Model  models.Model
	Result models.ModelResult
}

type FlowResultListInput struct {
	data   []FlowInput
	merged map[string]interface{}
}

func (m *FlowResultListInput) Input() map[string]interface{} {
	return m.merged
}

func (m *FlowResultListInput) Json() string {
	return objects.StringIndent(m.Input())
}

func NewFlowResultListInput(data []FlowInput) *FlowResultListInput {
	merged := make(map[string]interface{})
	for _, r := range data {
		if r.Result == nil {
			continue
		}
		_ = copier.Copy(&merged, r.Result.Result())
	}
	return &FlowResultListInput{
		data:   data,
		merged: merged,
	}
}
