package sconfig

import (
	"os"
	"testing"

	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
)

func TestDefaultInjector(t *testing.T) {
	assert.Equal(
		t,
		do.MustInvoke[*FrameConfig](Injector).Address(),
		do.MustInvoke[*FrameConfig](Injector).Address(),
	)
}

func TestServiceName(t *testing.T) {
	os.Setenv("ENV_ROLE", "test")
	assert.Equal(
		t,
		"/request-engine/re-audio",
		do.MustInvoke[*FrameConfig](Injector).ServiceName,
	)
}

func TestSentinel(t *testing.T) {
	systemRules := do.MustInvoke[*FrameConfig](Injector).Sentinel.SystemRules
	assert.Greater(t, len(systemRules), 0)
	assert.Equal(t, "load", systemRules[0].MetricType.String())
	assert.Greater(t, systemRules[0].TriggerCount, float64(0))
	assert.Equal(t, "bbr", systemRules[0].Strategy.String())
}

func TestSentinelSystemRule_ToSystemRule1(t *testing.T) {
	testRule := &SentinelSystemRule{
		Rule: system.Rule{
			ID:           "rule1",
			MetricType:   0,
			TriggerCount: 0,
			Strategy:     1,
		},
		TriggerCountFactor: 0.5,
	}
	assert.Less(t, float64(0), testRule.ToSystemRule().TriggerCount)
}

func TestSentinelSystemRule_ToSystemRule2(t *testing.T) {
	testRule := &SentinelSystemRule{
		Rule: system.Rule{
			ID:           "rule1",
			MetricType:   1,
			TriggerCount: 0,
			Strategy:     1,
		},
		TriggerCountFactor: 0.5,
	}
	assert.Equal(t, float64(0), testRule.ToSystemRule().TriggerCount)
}

func TestSentinelSystemRule_ToSystemRule3(t *testing.T) {
	testRule := &SentinelSystemRule{
		Rule: system.Rule{
			ID:           "rule1",
			MetricType:   1,
			TriggerCount: 0,
			Strategy:     1,
		},
		TriggerCountFactor: 0,
	}
	assert.Equal(t, float64(0), testRule.ToSystemRule().TriggerCount)
}
