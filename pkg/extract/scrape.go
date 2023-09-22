package extract

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var PROXY string

func Scrape(URL, target string, retry bool, engine func(url, target string) *Data) *[]Data {

	var (
		newURL      string
		status_code int
		scraped     = []Data{}
		c           = colly.NewCollector()
	)

	extensions.RandomUserAgent(c)

	if PROXY != "" {

		c.SetProxy("http://" + PROXY)
	}

	c.OnResponse(func(r *colly.Response) {

		status_code = r.StatusCode
	})

	c.OnHTML("a[href]", func(r *colly.HTMLElement) {

		newURL = r.Request.AbsoluteURL(r.Attr("href"))
		scraped = append(scraped, *engine(newURL, target))
	})

	c.Visit(URL)

	if retry && status_code >= 500 {

		time.Sleep(1 * time.Minute)
		return Scrape(URL, target, retry, engine)
	} else {

		return &scraped
	}
}
