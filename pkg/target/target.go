package target

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartTarget(addr, step string, sets int) error {
	sleepTime, err := time.ParseDuration(step)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().Unix())

	reg := prometheus.NewRegistry()

	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_percent_used",
			Help: "CPU percent used.",
		},
		[]string{"host", "module", "set"})
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"host", "module", "set", "path", "status_code"})
	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Time (in seconds) spent serving HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"host", "module", "set", "path", "status_code"})

	summary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "response_size_bytes",
			Help:       "Response size in bytes.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"host", "module", "set", "path", "status_code"})

	reg.MustRegister(counter)
	reg.MustRegister(gauge)
	reg.MustRegister(histogram)
	reg.MustRegister(summary)

	hostValues := []string{"10.2.0.4", "10.2.0.5", "10.3.0.6"}
	moduleValues := []string{"server", "redis"}
	pathValues := []string{"/api", "/home", "/auth"}
	setValues := make([]string, sets)
	for i := range setValues {
		setValues[i] = fmt.Sprintf("%d", i+1)
	}
	cpuPercent := map[string]float64{}

	go func() {
		for {
			for h, host := range hostValues {
				for m, module := range moduleValues {
					for s, set := range setValues {
						path := pathValues[rand.Intn(len(pathValues))]
						status := generateStatusCode()

						cpuKey := fmt.Sprintf("%d.%d.%d", h, m, s)
						cpu, ok := cpuPercent[cpuKey]
						if !ok {
							cpu = (float64)(40 + rand.Intn(40))
						}
						// change cpu percent +/- 2 range
						cpu = math.Abs(cpu + (float64)(2-rand.Intn(5)))
						cpuPercent[cpuKey] = cpu

						gauge.WithLabelValues(host, module, set).Set(cpu)
						counter.WithLabelValues(host, module, set, path, status).Inc()
						histogram.WithLabelValues(host, module, set, path, status).Observe(generateRequestTime())
						summary.WithLabelValues(host, module, set, path, status).Observe(generateResponseSize())
					}
				}
			}
			time.Sleep(sleepTime)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	return http.ListenAndServe(addr, nil)
}

// return value based on probability density
func generateValue(cdf []int, values []interface{}) interface{} {
	r := rand.Intn(100)
	bucket := 0
	for r > cdf[bucket] {
		bucket++
	}
	return values[bucket]
}

func generateStatusCode() string {
	// 200: OK, 401: Unauthorized, 503: Service Unavailable
	status := generateValue([]int{80, 95, 100}, []interface{}{"200", "401", "503"})
	return status.(string)
}

func generateResponseSize() float64 {
	base := generateValue([]int{20, 60, 80, 90, 100}, []interface{}{200, 400, 600, 800, 1000})
	return float64(base.(int) + rand.Intn(200))
}

func generateRequestTime() float64 {
	base := generateValue([]int{80, 95, 100}, []interface{}{20, 500, 800})
	return time.Duration(time.Millisecond * time.Duration(base.(int)+rand.Intn(200))).Seconds()
}
