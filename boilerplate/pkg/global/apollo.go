package global

import (
	json "github.com/bytedance/sonic"
	"github.com/day253/krpc/deepcopy"
	"github.com/day253/krpc/objects"
	"github.com/day253/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
)

type ApolloStrategyContext struct {
	featuresMap map[string]interface{}
	request     *service.PredictRequest
}

func (s *ApolloStrategyContext) getModelFeatures(toMerge map[string]interface{}) map[string]interface{} {
	// 此处的输入toMerge已经经过move逻辑的数据
	result := make(map[string]interface{})
	// 类似SendAllFeatures
	_ = deepcopy.Copy(&result, s.featuresMap)
	_ = copier.Copy(&result, toMerge)
	return result
}

// 目前并发控制不安全，因此每次调用都是复制出来的新的
func (s *ApolloStrategyContext) GetServicePredictRequest(toMerge map[string]interface{}) *service.PredictRequest {
	return &service.PredictRequest{
		RequestId:    s.request.RequestId,
		Organization: s.request.Organization,
		Data:         lo.ToPtr(objects.String(s.getModelFeatures(toMerge))),
	}
}

func NewApolloStrategyContext(request *service.PredictRequest) (*ApolloStrategyContext, error) {
	var featuresMap map[string]interface{}
	if err := json.Unmarshal([]byte(request.GetData()), &featuresMap); err != nil {
		return nil, err
	}
	return &ApolloStrategyContext{
		featuresMap: featuresMap,
		request:     request,
	}, nil
}
