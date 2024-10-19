package main

import (
	"context"
	re "github.com/ishumei/krpc/protocols/event/kitex_gen/shumei/strategy/re"
)

// EventPredictorImpl implements the last service interface defined in the IDL.
type EventPredictorImpl struct{}

// Predict implements the EventPredictorImpl interface.
func (s *EventPredictorImpl) Predict(ctx context.Context, request *re.EventPredictRequest) (resp *re.EventPredictResult_, err error) {
	// TODO: Your code here...
	return
}

// Health implements the EventPredictorImpl interface.
func (s *EventPredictorImpl) Health(ctx context.Context) (resp bool, err error) {
	// TODO: Your code here...
	return
}
