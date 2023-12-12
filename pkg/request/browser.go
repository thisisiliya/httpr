package request

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Browser(proxy string) *rod.Browser {

	if proxy != "" {

		url, _ := launcher.New().Set("proxy-server", proxy).Launch()
		browser := rod.New().Timeout(10 * time.Minute).ControlURL(url).MustConnect()

		return browser

	} else {

		browser := rod.New().Timeout(10 * time.Minute).MustConnect()

		return browser
	}
}
