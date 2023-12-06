package validate

import (
	"net/http"
	"net/url"

	"github.com/thisisiliya/go_utils/errors"
)

func Proxy(proxy string) bool {

	targetURL := "https://google.com"

	parsed_proxy, err := url.Parse(proxy)
	errors.Check(err)

	client := &http.Client{

		Transport: &http.Transport{

			Proxy: http.ProxyURL(parsed_proxy),
		},
	}

	if resp, err := client.Head(targetURL); err == nil {

		if resp.StatusCode == 200 {

			return true
		}
	}

	return false
}
