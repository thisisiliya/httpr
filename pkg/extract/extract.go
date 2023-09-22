package extract

import (
	"net/url"
	"strings"
)

type Data struct {
	URL       string
	Host      string
	Path      string
	Subdomain string
}

func GoogleExt(URL, target string) *Data {

	if parsedURL, err := url.Parse(URL); err == nil {

		if newURL := parsedURL.Query()["url"]; len(newURL) == 1 {

			if parsedNewURL, err := url.Parse(newURL[0]); err == nil {

				if strings.Contains(parsedNewURL.Host, target) {

					return &Data{
						URL:       newURL[0],
						Host:      parsedNewURL.Host,
						Path:      parsedNewURL.Path,
						Subdomain: SubdomainExt(parsedNewURL.Host),
					}
				}
			}
		}
	}
	return &Data{}
}

func SubdomainExt(host string) string {

	if domain := strings.Split(host, "."); len(domain) > 2 {

		return strings.Join(domain[0:len(domain)-2], ".")
	}

	return ""
}
