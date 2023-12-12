package request

import (
	"net/url"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
	"github.com/thisisiliya/httpr/pkg/engines"
)

func request(options *Engines, dork *engines.Options, browser *rod.Browser) *[]string {

	var result []string

	page := stealth.MustPage(browser)
	defer page.Close()

	page.MustNavigate(options.Engine(dork)).MustWaitStable()

	links := page.MustElements(options.Selector)

	for _, link := range links {

		href := link.MustAttribute("href")
		u, err := url.ParseRequestURI(*href)

		if err == nil && u.Scheme != "" {

			result = append(result, u.String())
		}
	}

	return &result
}
