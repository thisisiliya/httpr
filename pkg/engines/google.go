package engines

import (
	"net/url"
	"strconv"
)

var Google_Selector = "#center_col a[href]"

func GoogleURL(options *Options) string {

	var command string

	switch true {

	case options.Wildcard:
		command = "site:*." + options.Domain

	default:
		command = "site:" + options.Domain
	}

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
	return GoogleURLEncode(options)
}

func GoogleURLEncode(options *Options) string {

	values := url.Values{}

	values.Add("q", options.Command)
	values.Add("start", strconv.Itoa(options.Page*10))

	return "https://www.google.com/search?" + values.Encode()
}
