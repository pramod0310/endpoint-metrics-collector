package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"vmware/pkg/api"
	"vmware/pkg/conf"
	"vmware/pkg/prometheus"
)

func main() {

	config := conf.NewInstance()

	for _, httpEndPoint := range config.HttpEndpointConfigs {
		fmt.Println(httpEndPoint)
		httpClient := api.NewHttpClient(httpEndPoint.BaseURL, httpEndPoint.Scheme, httpEndPoint.TimeOut)
		for _, path := range httpEndPoint.Paths {
			go prometheus.CollectMetrics(httpClient, path)
		}
	}

	http.Handle(config.MetricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.MetricsPort), nil))
}
