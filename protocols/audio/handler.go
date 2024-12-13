package main

import (
	"context"

	re "github.com/day253/krpc/protocols/audio/kitex_gen/shumei/strategy/re"
)

// AudioPredictorImpl implements the last service interface defined in the IDL.
type AudioPredictorImpl struct{}

// Predict implements the AudioPredictorImpl interface.
func (s *AudioPredictorImpl) Predict(ctx context.Context, request *re.AudioPredictRequest) (resp *re.AudioPredictResult_, err error) {
	// TODO: Your code here...
	return
}

// Health implements the AudioPredictorImpl interface.
func (s *AudioPredictorImpl) Health(ctx context.Context) (resp bool, err error) {
	// TODO: Your code here...
	return
}
