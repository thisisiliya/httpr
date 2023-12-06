package request

import (
	"context"
	"net/url"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/thisisiliya/go_utils/errors"
)

func request(URL string, ctx *context.Context) *[]string {

	var result []string
	var nodes  []*cdp.Node

	err := chromedp.Run(

		*ctx,
		chromedp.Navigate(URL),
		chromedp.Nodes("a[href]", &nodes),
	)

	errors.Check(err)

	for _, n := range nodes {

		u, err := url.ParseRequestURI(n.AttributeValue("href"))

		if err == nil && u.Scheme != "" {

			result = append(result, u.String())
		}
	}

	return &result
}
