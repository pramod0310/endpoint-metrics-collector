package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"time"
	"vmware/pkg/api"
)

var (
	responseMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sample_external_url_response_ms",
			Help: "Checks response time in milliseconds",
		},
		[]string{"url"})

	healthStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sample_external_url_up",
			Help: "Checks if the Endpoint is up",

		},
		[]string{"url"})
)

func init(){
	prometheus.MustRegister(healthStatus)
	prometheus.MustRegister(responseMetric)
	fmt.Println("registered the metrics")
}

func CollectMetrics(httpClient api.Client, path string) {
	var response *api.Response
	var err error
	for {
		response, err = httpClient.Get(httpClient.GetURL(path))
		if err != nil {
			log.Println(err)
			requestedUrl := httpClient.GetURL(path)
			healthStatus.WithLabelValues(requestedUrl).Set(0)
			responseMetric.WithLabelValues(requestedUrl).Set(float64(httpClient.GetTimeout()))
			time.Sleep(10 *time.Second)
			continue
		}

		if response.StatusCode == 200 {
			healthStatus.WithLabelValues(response.URL).Set(1)
		}else {
			healthStatus.WithLabelValues(response.URL).Set(0)
		}

		responseMetric.WithLabelValues(response.URL).Set(float64(response.ResponseTime.Milliseconds()))
		time.Sleep(10 *time.Second)
	}
}