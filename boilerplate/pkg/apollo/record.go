package apollo

import (
	"os"
	"reflect"

	json "github.com/bytedance/sonic"
	"github.com/expr-lang/expr/ast"
	"github.com/expr-lang/expr/compiler"
	"github.com/expr-lang/expr/conf"
	"github.com/expr-lang/expr/file"
	"github.com/expr-lang/expr/optimizer"
	"github.com/expr-lang/expr/parser"
	"github.com/expr-lang/expr/vm"
)

type Operand struct {
	Operands []Operand `json:"operands"`
	Operator string    `json:"operator"`
	Type     string    `json:"type"`
	Values   []string  `json:"values,omitempty"`
}

func (o *Operand) Node() (*ast.BinaryNode, error) {
	return parseOperand(o), nil
}

func (o *Operand) Program() (*vm.Program, error) {
	astNode, err := o.Node()
	if err != nil {
		return nil, err
	}
	config := &conf.Config{Expect: reflect.Bool, ExpectAny: true, Strict: false}
	tree := &parser.Tree{
		Node:   astNode,
		Source: file.NewSource(""),
	}
	if err := optimizer.Optimize(&tree.Node, config); err != nil {
		return nil, err
	}
	program, err := compiler.Compile(tree, config)
	if err != nil {
		return nil, err
	}
	return program, nil
}

type RuleRecord struct {
	Condition    Operand `json:"_condition"`
	Product      string  `json:"product"`
	Organization string  `json:"organization"`
	SceneID      string  `json:"sceneId"`
	AppID        string  `json:"appId"`
	Channel      string  `json:"channel"`
	RiskLabel1   string  `json:"riskLabel1"`
	RiskLabel2   string  `json:"riskLabel2"`
	RiskLabel3   string  `json:"riskLabel3"`
	RiskLevel    string  `json:"riskLevel"`
	RuleID       string  `json:"ruleId"`
}

func (m *RuleRecord) Program() (*vm.Program, error) {
	return m.Condition.Program()
}

func (m *RuleRecord) HitResult() *HitResult {
	return &HitResult{
		RiskLabel1: m.RiskLabel1,
		RiskLabel2: m.RiskLabel2,
		RiskLabel3: m.RiskLabel3,
		RiskLevel:  m.RiskLevel,
		RuleID:     m.RuleID,
	}
}

func NewMetaFromFile(filePath string) (*RuleRecord, error) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return NewMetaFromJSON(jsonData)
}

func NewMetaFromJSON(jsonData []byte) (*RuleRecord, error) {
	var metaMap map[string]interface{}

	if err := json.Unmarshal(jsonData, &metaMap); err != nil {
		return nil, err
	}

	modifiedJsonData, err := json.Marshal(metaMap)
	if err != nil {
		return nil, err
	}

	var meta RuleRecord
	if err := json.Unmarshal(modifiedJsonData, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}
