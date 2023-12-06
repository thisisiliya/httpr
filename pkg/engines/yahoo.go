package engines

import (
	"net/url"
	"strconv"
)

func YahooURL(options *Options) string {

	var command string

	command = "site:" + options.Domain

	if options.Key != "" {
		command = command + ` "` + options.Key + `"`
	}

	if len(options.Block) != 0 {
		for _, Block := range options.Block {
			command = command + " -" + Block
		}
	}

	options.Command = command

	return YahooURLEncode(options)
}

func YahooURLEncode(options *Options) string {

	values := url.Values{}

	values.Add("p", options.Command)
	values.Add("b", strconv.Itoa((options.Page*7)+1))

	return "https://search.yahoo.com/search?" + values.Encode()
}
