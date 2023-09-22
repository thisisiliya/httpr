package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/thisisiliya/httpr/pkg/engines"

	"github.com/spf13/cobra"
)

var (
	rootRetry  bool
	rootVerify bool

	URL     string
	results []string

	wg sync.WaitGroup

	opt = engines.Options{}

	rootCmd = &cobra.Command{
		Use:  "httpr",
		Long: "\nHTTPR is an OSINT tool to Scrape the Undisclosed Data via Search Engines",
	}
)

func Execute() {

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {

		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}

func init() {

	rootCmd.PersistentFlags().StringP("proxy", "p", "", "proxy for scraping (ip:port)")
	rootCmd.PersistentFlags().Int("delay", 3, "max delay per requests")
	rootCmd.PersistentFlags().BoolVarP(&rootRetry, "retry", "r", false, "retry for status code above 500")
	rootCmd.PersistentFlags().BoolVar(&rootVerify, "verify", false, "verify the result by request")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "disable printing banner")
}
