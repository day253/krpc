package apollo

import (
	json "github.com/bytedance/sonic"
)

type HitResult struct {
	RiskLabel1 string `json:"riskLabel1"`
	RiskLabel2 string `json:"riskLabel2"`
	RiskLabel3 string `json:"riskLabel3"`
	RiskLevel  string `json:"riskLevel"`
	RuleID     string `json:"ruleId"`
}

type Result struct {
	HitResults []*HitResult
}

func (r *Result) String() string {
	json, err := json.Marshal(r.HitResults)
	if err != nil {
		return ""
	}
	return string(json)
}

func NewResult(hitResults []*HitResult) *Result {
	return &Result{HitResults: hitResults}
}
