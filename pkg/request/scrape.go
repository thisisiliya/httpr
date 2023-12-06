package request

import (
	"context"
	"net/url"
	"strings"
	"sync"

	"github.com/thisisiliya/go_utils/time"
)

type Data struct {
	URL       string
	Host      string
	Path      string
	Subdomain string
}

func Scrape(options *Options, result *[]Data, wg *sync.WaitGroup, ctx *context.Context) {

	for _, e := range options.Engines {

		wg.Add(1)
		engine := e

		go func() {

			defer wg.Done()

			URL := engine(&options.Dork)
			URLs := request(URL, ctx)

			for _, newURL := range *URLs {

				if parsedURL, err := url.Parse(newURL); err == nil {

					if strings.Contains(parsedURL.Host, options.Dork.Domain) {

						*result = append(*result, Data{
							URL:       newURL,
							Host:      parsedURL.Host,
							Path:      parsedURL.Path,
							Subdomain: subdomainExt(parsedURL.Host),
						})
					}
				}
			}
		}()

		time.RandomSleep(options.MinDelay, options.MaxDelay)
	}
}

func subdomainExt(host string) string {

	if domain := strings.Split(host, "."); len(domain) > 2 {

		return strings.Join(domain[0:len(domain)-2], ".")
	}

	return ""
}
