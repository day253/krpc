package strategies

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ishumei/dag"
	"github.com/ishumei/krpc/boilerplate/pkg/conf"
	"github.com/ishumei/krpc/boilerplate/pkg/models"
	"github.com/jinzhu/copier"
)

type StrategyDAG struct {
	Dag  *dag.DAG
	Root string
}

func (s *StrategyDAG) Execute(ctx context.Context) (map[string]interface{}, error) {
	result, err := s.Dag.DescendantsFlow(ctx, s.Root, nil, defaultCallback)
	if err != nil {
		return nil, err
	}
	if conf.Debug() {
		klog.CtxInfof(ctx, "graph: %s", s.Dag.String())
	}
	return func() (map[string]interface{}, error) {
		final := make(map[string]interface{})
		for _, r := range result {
			if r.Result == nil {
				continue
			}
			item := r.Result.(models.ModelResult).Result()
			_ = copier.Copy(&final, item)
		}
		return final, nil
	}()
}

func (s *StrategyDAG) String() string {
	return s.Dag.String()
}

func defaultCallback(ctx context.Context, d *dag.DAG, id string, parentResults []dag.FlowResult) (interface{}, error) {
	now := time.Now()
	v, _ := d.GetVertex(id)
	current, _ := v.(models.Model)
	defer func() {
		if r := recover(); r != nil {
			klog.CtxErrorf(ctx, "panic in callback: model:%s error=%v: stack=%s", current.Name(), r, debug.Stack())
		}
	}()
	input := NewFlowResultListInput(func() []FlowInput {
		mergeDeps := make([]FlowInput, len(parentResults))
		for _, r := range parentResults {
			p, _ := d.GetVertex(r.ID)
			if p == nil || r.Result == nil {
				continue
			}
			curDep := FlowInput{
				Model:  p.(models.Model),
				Result: r.Result.(models.ModelResult),
			}
			mergeDeps = append(mergeDeps, curDep)
			if conf.Debug() {
				klog.CtxInfof(ctx, "current: %s, depend: %s, result: %s", current.Name(), curDep.Model.Name(), curDep.Result.Json())
			}
		}
		return mergeDeps
	}())
	res, err := current.Run(ctx, input)
	elapsed := time.Since(now).Milliseconds()
	_ = Trace(&Entry{Context: ctx, ID: current.Name(), Elapsed: elapsed})
	klog.CtxInfof(ctx, "current: %s elapsed: %d", current.Name(), elapsed)
	return res, err
}

/*
	{
		"stages": [
			[
				"text-preprocess"
			],
			[
				"text-models",
				"text-list"
			],
			[
				"evaluation"
			]
		]
	}
placeholder0 // placeholder0起始节点，作为前驱连接stage0所有节点
stage0: text-preprocess
placeholder1 // placeholder1 作为stage0后继连接所有stage0节点，作为前驱连接stage1所有节点
stage1: text-models, text-list
placeholder2 // placeholder2 作为stage1后继连接所有stage1节点，作为前驱连接stage2所有节点
stage2: evaluation
placeholder3 // placeholder3 作为stage2后继连接所有stage2节点
*/
// Function to create a DAG from a Stages struct
func NewStrategyDAG(stages *Stages) (*StrategyDAG, error) {
	d := dag.NewDAG()
	placeholders := make([]string, len(stages.Stages)+1)
	for i := 0; i <= len(stages.Stages); i++ {
		placeholder, err := models.GetPlaceholder(fmt.Sprintf("placeholder%d", i))
		if err != nil {
			return nil, fmt.Errorf("failed to add placeholder vertex: %w", err)
		}
		placeholderVertex, err := d.AddVertex(placeholder)
		if err != nil {
			return nil, fmt.Errorf("failed to add placeholder vertex: %w", err)
		}
		placeholders[i] = placeholderVertex
	}
	for i, stage := range stages.Stages {
		for _, modelName := range stage {
			model, err := models.GetPredictor(modelName)
			if err != nil {
				return nil, fmt.Errorf("failed to add vertex %s: %w", modelName, err)
			}
			vertexID, err := d.AddVertex(model)
			if err != nil {
				return nil, fmt.Errorf("failed to add vertex %s: %w", modelName, err)
			}
			err = d.AddEdge(placeholders[i], vertexID)
			if err != nil {
				return nil, fmt.Errorf("failed to add edge from placeholder to %s: %w", modelName, err)
			}
			err = d.AddEdge(vertexID, placeholders[i+1])
			if err != nil {
				return nil, fmt.Errorf("failed to add edge from %s to placeholder: %w", modelName, err)
			}
		}
	}
	return &StrategyDAG{
		Dag:  d,
		Root: placeholders[0],
	}, nil
}
