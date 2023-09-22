package validate

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func Proxy(proxy string) bool {

	targetURL := "https://google.com"

	parsed_proxy, err := url.Parse("http://" + proxy)

	if err != nil {

		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}

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
