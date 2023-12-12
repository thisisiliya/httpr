package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thisisiliya/httpr/pkg/engines"
	"github.com/thisisiliya/httpr/pkg/request"
)

var (
	URL     string
	results []string

	root_Proxy  string
	root_Silent bool

	i Options         // User inputs
	o request.Options // Request options

	rootCmd = &cobra.Command{
		Use: "httpr",
		Long: ("\nHTTPR is an OSINT tool to Scrape the Undisclosed Data via Search Engines" +
			"\nfor more information visit https://github.com/thisisiliya/httpr"),
	}
)

func Execute() {

	o.Engines = []request.Engines{
		{Engine: engines.GoogleURL, Selector: engines.Google_Selector},
		{Engine: engines.BingURL, Selector: engines.Bing_Selector},
		{Engine: engines.YahooURL, Selector: engines.Yahoo_Selector},
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {

		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&root_Proxy, "proxy", "p", "", "proxy or scraping (ip:port)")
	rootCmd.PersistentFlags().IntVar(&i.root_MinDelay, "min-delay", 1, "min delay per request in sec")
	rootCmd.PersistentFlags().IntVar(&i.root_MaxDelay, "max-delay", 10, "max delay per request in sec")
	rootCmd.PersistentFlags().BoolVarP(&i.root_Verify, "verify", "v", false, "verify the result by a request")
	rootCmd.PersistentFlags().BoolVarP(&root_Silent, "silent", "s", false, "disable printing banner")
}
