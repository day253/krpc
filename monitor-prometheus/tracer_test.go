package monitor_prometheus

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

func TestPrometheusReporter(t *testing.T) {
	registry := prom.NewRegistry()
	http.Handle("/prometheus", promhttp.HandlerFor(registry, promhttp.HandlerOpts{ErrorHandling: promhttp.ContinueOnError}))
	go func() {
		if err := http.ListenAndServe(":9090", nil); err != nil {
			t.Error("Unable to start a promhttp server, err: " + err.Error())
			return
		}
	}()

	counter := prom.NewCounterVec(
		prom.CounterOpts{
			Name:        "test_counter",
			ConstLabels: prom.Labels{"service": "prometheus-test"},
		},
		[]string{"test1", "test2"},
	)
	registry.MustRegister(counter)

	histogram := prom.NewHistogramVec(
		prom.HistogramOpts{
			Name:        "test_histogram",
			ConstLabels: prom.Labels{"service": "prometheus-test"},
			Buckets:     prom.DefBuckets,
		},
		[]string{"test1", "test2"},
	)
	registry.MustRegister(histogram)

	labels := prom.Labels{
		"test1": "abc",
		"test2": "def",
	}

	assert.True(t, counterAdd(counter, 6, labels) == nil)
	assert.True(t, histogramObserve(histogram, time.Second, labels) == nil)

	promServerResp, err := http.Get("http://localhost:9090/prometheus")
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, promServerResp.StatusCode == http.StatusOK)

	bodyBytes, err := io.ReadAll(promServerResp.Body)
	assert.True(t, err == nil)
	respStr := string(bodyBytes)
	assert.True(t, strings.Contains(respStr, `test_counter{service="prometheus-test",test1="abc",test2="def"} 6`))
	assert.True(t, strings.Contains(respStr, `test_histogram_sum{service="prometheus-test",test1="abc",test2="def"} 1e+06`))
}

func Test_removeDynamicDetail(t *testing.T) {
	type args struct {
		val string
		def string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_removeDynamicDetail",
			args: args{
				val: "dail 127.0.0.1:7000 failed",
				def: "",
			},
			want: `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`,
		},
		{
			name: "test_removeDynamicDetail",
			args: args{
				val: "dail instance failed",
				def: "",
			},
			want: "dail instance failed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDynamicDetail(tt.args.val, tt.args.def); got != tt.want {
				t.Errorf("removeDynamicDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
