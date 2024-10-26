package monitor

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// counterAdd wraps Add of prometheus.Counter.
func counterAdd(counterVec *prometheus.CounterVec, value int, labels prometheus.Labels) error {
	counter, err := counterVec.GetMetricWith(labels)
	if err != nil {
		return err
	}
	counter.Add(float64(value))
	return nil
}

// histogramObserve wraps Observe of prometheus.Observer.
func histogramObserve(histogramVec *prometheus.HistogramVec, value time.Duration, labels prometheus.Labels) error {
	histogram, err := histogramVec.GetMetricWith(labels)
	if err != nil {
		return err
	}
	histogram.Observe(float64(value.Microseconds()))
	return nil
}
