package http

import (
	"sync"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"github.com/hashicorp/go-retryablehttp"
)

var httpClient *retryablehttp.Client
var once sync.Once

func GetInstance() *retryablehttp.Client {
	once.Do(
		func() {

			httpLogMessage := logger.HttpLogMessage{}

			httpClient = retryablehttp.NewClient()

			httpClient.Logger = httpLogMessage
		},
	)

	return httpClient
}
