package prometheus

import (
	"errors"
	"github.com/golang/mock/gomock"
	promtestutil "github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
	"vmware/mock"
	"vmware/pkg/api"
)

const metadataUrlUp = `
		# HELP sample_external_url_up Checks if the Endpoint is up
		# TYPE sample_external_url_up gauge
	`
const metadataResponseMS = `
		# HELP sample_external_url_response_ms Checks response time in milliseconds
        # TYPE sample_external_url_response_ms gauge
	`
func TestCollectMetrics(t *testing.T) {
	ctl := gomock.NewController(t)
	apiMock := mock.NewMockClient(ctl)

	apiMock.EXPECT().Get(gomock.Any()).Return(&api.Response{
		URL: "https://testurl.com/200",
		ResponseTime: 2*time.Millisecond,
		StatusCode: 200,

	},nil)

	apiMock.EXPECT().GetURL(gomock.Any()).Return("https://testurl.com/200")

	go CollectMetrics(apiMock,"200")

	time.Sleep(1 *time.Second)
	expectedUrlUp := `

		sample_external_url_up{url="https://testurl.com/200"} 1
	`
	expectedResponseMS := `

		sample_external_url_response_ms{url="https://testurl.com/200"} 2
	`



	err := promtestutil.CollectAndCompare(healthStatus,strings.NewReader(metadataUrlUp+expectedUrlUp), "sample_external_url_up")

	assert.Nil(t, err, "prometheus get values error")

	err = promtestutil.CollectAndCompare(responseMetric,strings.NewReader(metadataResponseMS+expectedResponseMS), "sample_external_url_response_ms")

	assert.Nil(t, err, "prometheus get values error")


}

func TestCollectMetrics_Error(t *testing.T) {
	ctl := gomock.NewController(t)
	apiMock := mock.NewMockClient(ctl)

	apiMock.EXPECT().Get(gomock.Any()).Return(nil,errors.New("response timed out"))


	firstCall := apiMock.EXPECT().GetURL(gomock.Any()).Return("https://testurl.com/200")
	secondCall := apiMock.EXPECT().GetURL(gomock.Any()).Return("https://testurl.com/200")

	gomock.InOrder(
		firstCall,
		secondCall,
	)

	apiMock.EXPECT().GetTimeout().Return(time.Duration(10))

	go CollectMetrics(apiMock,"200")

	time.Sleep(1 *time.Second)
	expectedUrlUp := `
	
		sample_external_url_up{url="https://testurl.com/200"} 0
	`
	expectedResponseMS := `
	
		sample_external_url_response_ms{url="https://testurl.com/200"} 10
	`
	err := promtestutil.CollectAndCompare(healthStatus,strings.NewReader(metadataUrlUp+expectedUrlUp), "sample_external_url_up")

	assert.Nil(t, err, "prometheus get values error")

	err = promtestutil.CollectAndCompare(responseMetric,strings.NewReader(metadataResponseMS+expectedResponseMS), "sample_external_url_response_ms")

	assert.Nil(t, err, "prometheus get values error")


}
