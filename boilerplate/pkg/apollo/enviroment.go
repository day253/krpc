package apollo

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/expr-lang/expr/vm"
	"github.com/ishumei/krpc/objects"
	"github.com/sourcegraph/conc/pool"
	"gorm.io/gorm"
)

const (
	DefaultBatchSize = 500
)

type Enviroment struct {
	batchSize    int
	meta         []*HitResult
	program      []*vm.Program
	productIndex map[string]*ProductIndex
}

func (r *Enviroment) Query(product, organization, sceneId, appId, channel string) []int {
	if productIndex, exists := r.productIndex[product]; exists {
		return productIndex.query(organization, sceneId, appId, channel)
	}
	return []int{}
}

func (r *Enviroment) Add(ruleRecord *RuleRecord, count int) error {
	program, err := ruleRecord.Program()
	if err != nil {
		return err
	}
	r.addProgram(program)
	r.addMeta(ruleRecord.HitResult())
	r.addRuleRecord(ruleRecord, count)
	return nil
}

func (r *Enviroment) addRuleRecord(ruleRecord *RuleRecord, count int) {
	product := ruleRecord.Product
	if _, exists := r.productIndex[product]; !exists {
		r.productIndex[product] = NewProductIndex()
	}
	r.productIndex[product].add(ruleRecord.Organization, ruleRecord.SceneID, ruleRecord.AppID, ruleRecord.Channel, count)
}

func (r *Enviroment) addProgram(program *vm.Program) {
	r.program = append(r.program, program)
}

func (r *Enviroment) addMeta(meta *HitResult) {
	r.meta = append(r.meta, meta)
}

func (r *Enviroment) stat() map[string]interface{} {
	res := make(map[string]interface{})
	for product, productIndex := range r.productIndex {
		res[product] = productIndex.stat()
	}
	return res
}

func (r *Enviroment) Stat() string {
	return objects.String(r.stat())
}

func (r *Enviroment) Count() int {
	return len(r.program)
}

func (r *Enviroment) intervalEval(recallList []int, start, end int, features map[string]interface{}) []int {
	res := []int{}
	maxIndex := r.Count()
	for i := start; i < end; i++ {
		idx := recallList[i]
		if idx < maxIndex {
			program := r.program[idx]
			if result, err := vm.Run(program, features); err == nil && result.(bool) {
				res = append(res, idx)
			}
		}
	}
	return res
}

func (r *Enviroment) BatchSize() int {
	if r.batchSize == 0 {
		return DefaultBatchSize
	}
	return r.batchSize
}

func (r *Enviroment) batchEval(recallList []int, features map[string]interface{}) [][]int {
	p := pool.NewWithResults[[]int]()
	batchSize := r.BatchSize()
	maxIndex := len(recallList)
	for i := 0; i < maxIndex; i += batchSize {
		start := i
		end := i + batchSize
		if end > maxIndex {
			end = maxIndex
		}
		p.Go(func() []int {
			return r.intervalEval(recallList, start, end, features)
		})
	}
	return p.Wait()
}

func (r *Enviroment) HitMetas(results [][]int) []*HitResult {
	matchedMetas := make([]*HitResult, 0)
	for _, result := range results {
		for _, idx := range result {
			matchedMetas = append(matchedMetas, r.meta[idx])

		}
	}
	return matchedMetas
}

func (r *Enviroment) EvalFromFeatures(product, organization, sceneId, appId, channel string, features map[string]interface{}) (*Result, error) {
	return NewResult(r.HitMetas(r.batchEval(r.Query(product, organization, sceneId, appId, channel), features))), nil
}

type Option func(r *Enviroment)

func WithBatchSize(batchSize int) Option {
	return func(r *Enviroment) {
		r.batchSize = batchSize
	}
}

func NewEnviroment(db *gorm.DB, options ...Option) (*Enviroment, error) {
	rule := &Enviroment{
		productIndex: make(map[string]*ProductIndex),
	}

	for _, option := range options {
		option(rule)
	}

	query := db.Model(&RuleRecord{})

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}

	rule.meta = make([]*HitResult, 0, count)
	rule.program = make([]*vm.Program, 0, count)

	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var ruleRecord RuleRecord
		if err := db.ScanRows(rows, &ruleRecord); err != nil {
			klog.Error("failed to scan row at count ", i, ": ", err)
			continue
		}
		if err := rule.Add(&ruleRecord, i); err != nil {
			klog.Error("failed to add rule at count ", i, ": ", err)
			continue
		}
	}

	klog.Info("query: ", count, ", load ", rule.Count())
	return rule, nil
}
