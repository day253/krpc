package strategies

import (
	json "github.com/bytedance/sonic"
	"github.com/samber/do"
)

var Injector = do.New()

type Stages struct {
	TimeoutMs int        `json:"timeout_ms"`
	Stages    [][]string `json:"stages"`
}

func NewStages(jsonStr []byte) (*Stages, error) {
	var stages Stages
	err := json.Unmarshal(jsonStr, &stages)
	if err != nil {
		return nil, err
	}
	return &stages, nil
}
