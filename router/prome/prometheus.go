/**********************************************
** @Des:
** @Author: lg1024
** @Last Modified time: 2019/12/19 下午2:34
***********************************************/

package prome

import "github.com/prometheus/client_golang/prometheus"

var collectorContainer []prometheus.Collector

var (
//ActivitySigninSendVipTotal     *prometheus.CounterVec
//ActivityDailyTaskSendCodeTotal *prometheus.CounterVec
//ProcessingTime                 *prometheus.HistogramVec
//RequestDurations               *prometheus.SummaryVec
)

// InitPrometheus ... initialize prometheus
func InitPrometheus() {
	prometheus.MustRegister(collectorContainer...)
}

// PushRegister ... Push collectores to prometheus before initializing
func PushRegister(c ...prometheus.Collector) {
	collectorContainer = append(collectorContainer, c...)
}

/*
func InitMetrics() {
	ActivitySigninSendVipTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "game",
			Subsystem: "activity",
			Name:      "signin_send_vip_total",
			Help:      "Total signin send vip count",
		},
		[]string{"worker_id", "type"},
	)

	RequestDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "request_time_cost",
			Help:       "req latency distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"module", "req_name", "status_code"},
	)

	ProcessingTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "request",
			Subsystem: "jobs",
			Name:      "process_time_seconds",
			Help:      "Amount of time spent processing jobs",
		},
		[]string{"worker_id", "type"},
	)

	PushRegister(ProcessingTime, ActivitySigninSendVipTotal, ActivityDailyTaskSendCodeTotal, RequestDurations)
}*/
