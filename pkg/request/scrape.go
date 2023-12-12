package request

import (
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

func Scrape(options *Options) *[]Data {

	var result []Data
	var wg_task sync.WaitGroup
	var wg_time sync.WaitGroup

	wg_time.Add(1)

	go func() {

		time.RandomSleep(options.MinDelay, options.MaxDelay)
		wg_time.Done()
	}()

	for _, opt := range options.Engines {

		wg_task.Add(1)
		opt := opt

		go func() {

			defer wg_task.Done()

			URLs := request(&opt, &options.Dork, options.Browser)

			for _, newURL := range *URLs {

				if parsed_URL, err := url.Parse(newURL); err == nil {

					if strings.Contains(parsed_URL.Host, options.Dork.Domain) {

						result = append(result, Data{
							URL:       newURL,
							Host:      parsed_URL.Host,
							Path:      parsed_URL.Path,
							Subdomain: subdomainExt(parsed_URL.Host),
						})
					}
				}
			}
		}()
	}

	wg_task.Wait()
	wg_time.Wait()

	return &result
}

func subdomainExt(host string) string {

	if domain := strings.Split(host, "."); len(domain) > 2 {

		return strings.Join(domain[0:len(domain)-2], ".")
	}

	return ""
}
