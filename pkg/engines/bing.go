package engines

import (
	"net/url"
	"strconv"
)

var Bing_Selector = "#b_results a[href]"

func BingURL(options *Options) string {

	var command string

	command = "site:" + options.Domain

	if options.Filetype != "" {
		command = command + " filetype:" + options.Filetype
	}

	if options.Key != "" {
		command = command + ` "` + options.Key + `"`
	}

	if len(options.Block) != 0 {
		for _, Block := range options.Block {
			command = command + " -" + Block
		}
	}

	options.Command = command

	return BingURLEncode(options)
}

func BingURLEncode(options *Options) string {

	values := url.Values{}

	values.Add("q", options.Command)
	values.Add("first", strconv.Itoa((options.Page*10)+1))

	return "https://www.bing.com/search?" + values.Encode()
}
