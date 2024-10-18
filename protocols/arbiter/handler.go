package main

import (
	"context"

	service "github.com/ishumei/krpc/protocols/arbiter/kitex_gen/com/shumei/service"
)

// PredictorImpl implements the last service interface defined in the IDL.
type PredictorImpl struct{}

// Predict implements the PredictorImpl interface.
func (s *PredictorImpl) Predict(ctx context.Context, request *service.PredictRequest) (resp *service.PredictResult_, err error) {
	// TODO: Your code here...
	return
}

// Health implements the PredictorImpl interface.
func (s *PredictorImpl) Health(ctx context.Context) (resp bool, err error) {
	// TODO: Your code here...
	return
}
