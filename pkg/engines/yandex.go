package engines

import (
	"net/url"
	"strconv"
)

var Yandex_Selector = "#search-result a[href]"

func YandexURL(options *Options) string {

	var command string

	command = "site:" + options.Domain

	if options.Filetype != "" {
		command = command + " mime:" + options.Filetype
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

	return YandexURLEncode(options)
}

func YandexURLEncode(options *Options) string {

	values := url.Values{}

	values.Add("text", options.Command)
	values.Add("p", strconv.Itoa(options.Page))

	return "https://www.yandex.com/search/?" + values.Encode()
}
