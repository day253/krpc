package main

import (
	"context"

	re "github.com/day253/krpc/protocols/image/kitex_gen/shumei/strategy/re"
)

// ImagePredictorImpl implements the last service interface defined in the IDL.
type ImagePredictorImpl struct{}

// Predict implements the ImagePredictorImpl interface.
func (s *ImagePredictorImpl) Predict(ctx context.Context, request *re.ImagePredictRequest) (resp *re.ImagePredictResult_, err error) {
	// TODO: Your code here...
	return
}

// Health implements the ImagePredictorImpl interface.
func (s *ImagePredictorImpl) Health(ctx context.Context) (resp bool, err error) {
	// TODO: Your code here...
	return
}
