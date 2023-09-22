package engines

import (
	"net/url"
	"strconv"
)

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

	return EncodeGoogleURL(command, options.Page)
}

func EncodeGoogleURL(command string, page int) string {

	values := url.Values{}

	values.Add("q", command)
	values.Add("start", strconv.Itoa(page*10))

	return "https://www.google.com/search?" + values.Encode()
}
