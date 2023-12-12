package request

import (
	"github.com/go-rod/rod"
	"github.com/thisisiliya/httpr/pkg/engines"
)

type Options struct {
	Dork     engines.Options
	Engines  []Engines
	Browser  *rod.Browser
	Timeout  int
	MinDelay int
	MaxDelay int
}

type Engines struct {
	Engine   func(options *engines.Options) string
	Selector string
}
