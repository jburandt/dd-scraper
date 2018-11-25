package cmd

import (
	"github.com/zorkian/go-datadog-api"
	"unicode"
)

type ClientLogin struct {
	Datadog *datadog.Client
}

func DatadogClient(apiKey, appKey string) (ClientLogin, error) {
	client := ClientLogin{
		Datadog: datadog.NewClient(apiKey, appKey),
	}

	return client, nil
}

func IntCheck(s string) bool {
	for _, i := range s {
		if !unicode.IsDigit(i) {
			return false
		}
	}
	return true
}
