package monitor

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type ArchClientLabel struct {
	Caller string
	Callee string
	Method string
	Status string
	Detail string
	Retry  int
}

func ArchClientThroughput(labels *ArchClientLabel) {
	if labels == nil {
		return
	}
	var l = make(prometheus.Labels)
	l[labelKeyCaller] = defaultValIfEmpty(labels.Caller, unknownLabelValue)
	l[labelKeyCallee] = defaultValIfEmpty(labels.Callee, unknownLabelValue)
	l[labelKeyMethod] = defaultValIfEmpty(labels.Method, unknownLabelValue)
	l[labelKeyStatus] = defaultValIfEmpty(labels.Status, unknownLabelValue)
	l[labelKeyDetail] = removeDynamicDetail(labels.Detail, unknownLabelValue)
	l[labelKeyRetry] = defaultValIfEmpty(strconv.Itoa(labels.Retry), unknownLabelValue)
	_ = counterAdd(globalClientHandledCounter, 1, l)
}

func ArchClientLatencyUs(value time.Duration, labels *ArchClientLabel) {
	if labels == nil {
		return
	}
	var l = make(prometheus.Labels)
	l[labelKeyCaller] = defaultValIfEmpty(labels.Caller, unknownLabelValue)
	l[labelKeyCallee] = defaultValIfEmpty(labels.Callee, unknownLabelValue)
	l[labelKeyMethod] = defaultValIfEmpty(labels.Method, unknownLabelValue)
	l[labelKeyStatus] = defaultValIfEmpty(labels.Status, unknownLabelValue)
	l[labelKeyDetail] = removeDynamicDetail(labels.Detail, unknownLabelValue)
	l[labelKeyRetry] = defaultValIfEmpty(strconv.Itoa(labels.Retry), unknownLabelValue)
	_ = histogramObserve(globalClientHandledHistogram, value, l)
}

type ArchServerLabel struct {
	Caller string
	Callee string
	Method string
	Status string
	Detail string
	Retry  int
}

func ArchServerThroughput(labels *ArchServerLabel) {
	if labels == nil {
		return
	}
	var l = make(prometheus.Labels)
	l[labelKeyCaller] = defaultValIfEmpty(labels.Caller, unknownLabelValue)
	l[labelKeyCallee] = defaultValIfEmpty(labels.Callee, unknownLabelValue)
	l[labelKeyMethod] = defaultValIfEmpty(labels.Method, unknownLabelValue)
	l[labelKeyStatus] = defaultValIfEmpty(labels.Status, unknownLabelValue)
	l[labelKeyDetail] = removeDynamicDetail(labels.Detail, unknownLabelValue)
	l[labelKeyRetry] = defaultValIfEmpty(strconv.Itoa(labels.Retry), unknownLabelValue)
	_ = counterAdd(globalServerHandledCounter, 1, l)
}

func ArchServerLatencyUs(value time.Duration, labels *ArchServerLabel) {
	if labels == nil {
		return
	}
	var l = make(prometheus.Labels)
	l[labelKeyCaller] = defaultValIfEmpty(labels.Caller, unknownLabelValue)
	l[labelKeyCallee] = defaultValIfEmpty(labels.Callee, unknownLabelValue)
	l[labelKeyMethod] = defaultValIfEmpty(labels.Method, unknownLabelValue)
	l[labelKeyStatus] = defaultValIfEmpty(labels.Status, unknownLabelValue)
	l[labelKeyDetail] = removeDynamicDetail(labels.Detail, unknownLabelValue)
	l[labelKeyRetry] = defaultValIfEmpty(strconv.Itoa(labels.Retry), unknownLabelValue)
	_ = histogramObserve(globalServerHandledHistogram, value, l)
}
