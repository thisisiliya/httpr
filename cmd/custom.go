package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thisisiliya/httpr/pkg/engines"
	"github.com/thisisiliya/httpr/pkg/request"
	"github.com/thisisiliya/httpr/pkg/utils"
	"golang.org/x/exp/slices"
)

var customCmd = &cobra.Command{
	Use:     "custom",
	Short:   "custom dork command to scrape",
	Long:    "engine page(s) scrape by custom dork commands",
	Example: `httpr custom --command "site:*.hackerone.com inurl:report" --target-host hackerone.com --engine Google --depth 1`,
	Run:     CustomEnum,
}

func CustomEnum(_ *cobra.Command, _ []string) {

	utils.Start(root_Silent)

	o.MinDelay = i.root_MinDelay
	o.MaxDelay = i.root_MaxDelay

	o.Browser = request.Browser(root_Proxy, root_Timeout, root_Chromium)
	defer o.Browser.MustClose()

	o.Dork.Domain = i.custom_TargetHost
	o.Dork.Command = strings.ReplaceAll(i.custom_Command, i.custom_SpiltBy, " ")

	switch strings.ToLower(i.custom_Engine) {

	case "google":
		o.Engines = []request.Engines{{Engine: engines.GoogleURLEncode, Selector: engines.Google_Selector}}

	case "bing":
		o.Engines = []request.Engines{{Engine: engines.BingURLEncode, Selector: engines.Bing_Selector}}

	case "yahoo":
		o.Engines = []request.Engines{{Engine: engines.YahooURLEncode, Selector: engines.Yahoo_Selector}}

	case "yandex":
		o.Engines = []request.Engines{{Engine: engines.YandexURLEncode, Selector: engines.Yandex_Selector}}

	default:
		o.Engines = []request.Engines{}
	}

	var data []request.Data

	for o.Dork.Page < i.custom_Depth {

		data = append(data, *request.Scrape(&o)...)
		o.Dork.Page++
	}

	for _, d := range data {

		if CustomValidate(&d) {

			switch true {

			case i.custom_ShowHost:
				fmt.Println(d.Host)

			case i.custom_ShowPath:
				fmt.Println(d.Path)

			case i.custom_ShowSub:
				fmt.Println(d.Subdomain)

			default:
				fmt.Println(d.URL)
			}
		}
	}
}

func CustomValidate(d *request.Data) bool {

	if d.URL != "" {

		if !slices.Contains(results, d.URL) {

			if i.root_Verify {

				if request.Verify(d.URL) {

					results = append(results, d.URL)
					return true
				}
			} else {

				results = append(results, d.URL)
				return true
			}
		}
	}

	return false
}

func init() {

	rootCmd.AddCommand(customCmd)

	customCmd.Flags().StringVarP(&i.custom_Command, "command", "c", "", "dork command to scrape")
	customCmd.Flags().StringVarP(&i.custom_Engine, "engine", "e", "Google", "target engine to scrape. available engines: Google, Bing, Yahoo")
	customCmd.Flags().StringVarP(&i.custom_TargetHost, "target-host", "t", "", "filter result by host")
	customCmd.Flags().StringVar(&i.custom_SpiltBy, "split-by", " ", "dork commands split character")
	customCmd.Flags().IntVar(&i.custom_Depth, "depth", 1, "number of pages to scrape")
	customCmd.Flags().BoolVar(&i.custom_ShowHost, "show-host", false, "show hosts as result")
	customCmd.Flags().BoolVar(&i.custom_ShowSub, "show-sub", false, "show subdomains as result")
	customCmd.Flags().BoolVar(&i.custom_ShowPath, "show-path", false, "show paths as result")

	customCmd.MarkFlagRequired("command")
	customCmd.MarkFlagsMutuallyExclusive("show-host", "show-sub", "show-path")
}
