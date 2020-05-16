package target

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
	Address     string
	MetricsPath string
	JobName     string
	Sets        int
	UpdateRate  string
	TlsEnabled  bool
	TlsKeyPath  string
	TlsCertPath string
}

func StartTarget(cfg *Config) error {
	sleepTime, err := time.ParseDuration(cfg.UpdateRate)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().Unix())

	reg := prometheus.NewRegistry()

	gaugeLabelNames := []string{"host", "module", "set"}
	labelNames := []string{"host", "module", "set", "path", "status_code"}
	if len(cfg.JobName) > 0 {
		gaugeLabelNames = append(gaugeLabelNames, "job")
		labelNames = append(labelNames, "job")
	}
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_percent_used",
			Help: "CPU percent used.",
		}, gaugeLabelNames)
	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		}, labelNames)
	histogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Time (in seconds) spent serving HTTP requests.",
			Buckets: prometheus.DefBuckets,
		}, labelNames)
	summary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "response_size_bytes",
			Help:       "Response size in bytes.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, labelNames)

	reg.MustRegister(counter)
	reg.MustRegister(gauge)
	reg.MustRegister(histogram)
	reg.MustRegister(summary)

	moduleValues := []string{"frontend", "backend", "api"}
	pathValues := []string{"/api", "/home", "/auth"}
	setValues := make([]string, cfg.Sets)
	for i := range setValues {
		setValues[i] = fmt.Sprintf("%d", i+1)
	}
	cpuPercent := map[string]float64{}

	go func() {
		for {
			host := cfg.Address
			for m, module := range moduleValues {
				for s, set := range setValues {
					path := pathValues[rand.Intn(len(pathValues))]
					status := generateStatusCode()

					cpuKey := fmt.Sprintf("%d.%d", m, s)
					cpu, ok := cpuPercent[cpuKey]
					if !ok {
						cpu = (float64)(40 + rand.Intn(40))
					}
					// change cpu percent +/- 2 range
					cpu = math.Abs(cpu + (float64)(2-rand.Intn(5)))
					cpuPercent[cpuKey] = cpu

					gaugeLabelValues := []string{host, module, set}
					labelValues := []string{host, module, set, path, status}
					if len(cfg.JobName) > 0 {
						gaugeLabelValues = append(gaugeLabelValues, cfg.JobName)
						labelValues = append(labelValues, cfg.JobName)
					}
					gauge.WithLabelValues(gaugeLabelValues...).Set(cpu)
					counter.WithLabelValues(labelValues...).Inc()
					histogram.WithLabelValues(labelValues...).Observe(generateRequestTime())
					summary.WithLabelValues(labelValues...).Observe(generateResponseSize())
				}
			}

			time.Sleep(sleepTime)
		}
	}()

	if !strings.HasPrefix(cfg.MetricsPath, "/") {
		cfg.MetricsPath = "/" + cfg.MetricsPath
	}
	http.Handle(cfg.MetricsPath, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	if cfg.TlsEnabled {
		return http.ListenAndServeTLS(cfg.Address, cfg.TlsCertPath, cfg.TlsKeyPath, nil)
	}
	return http.ListenAndServe(cfg.Address, nil)
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
