package main

import (
	"context"
	re "github.com/ishumei/krpc/protocols/text/kitex_gen/shumei/strategy/re"
)

// TextPredictorImpl implements the last service interface defined in the IDL.
type TextPredictorImpl struct{}

// Predict implements the TextPredictorImpl interface.
func (s *TextPredictorImpl) Predict(ctx context.Context, request *re.TextPredictRequest) (resp *re.TextPredictResult_, err error) {
	// TODO: Your code here...
	return
}

// Health implements the TextPredictorImpl interface.
func (s *TextPredictorImpl) Health(ctx context.Context) (resp bool, err error) {
	// TODO: Your code here...
	return
}
