package request

import (
	"github.com/thisisiliya/httpr/pkg/engines"
)

type Options struct {
	Dork     engines.Options
	Engines  []func(options *engines.Options) string
	Timeout  int
	MinDelay int
	MaxDelay int
}
