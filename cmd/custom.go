package cmd

import (
	"fmt"
	"strings"

	"github.com/thisisiliya/httpr/pkg/engines"
	"github.com/thisisiliya/httpr/pkg/extract"
	"github.com/thisisiliya/httpr/pkg/start"
	"github.com/thisisiliya/httpr/pkg/validate"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var (
	customCommand    string
	customTargetHost string
	customSpiltBy    string
	customDepth      int
	customShowHost   bool
	customShowPath   bool
	customShowSub    bool

	customCmd = &cobra.Command{
		Use:   "custom",
		Short: "custom dork command to scrape",
		Long: ("\ngoogle page(s) scrape by custom dork commands" +
			"\nusage: -c site:www.google.com,inurl:map -t google.com --depth 1"),
		Run: customEnum,
	}
)

func customEnum(_ *cobra.Command, _ []string) {

	start.Start(rootCmd)

	customCommand = strings.ReplaceAll(customCommand, customSpiltBy, " ")

	for page := 0; page < customDepth; page++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			URL = engines.EncodeGoogleURL(customCommand, page)

			for _, r := range *extract.Scrape(URL, customTargetHost, rootRetry, extract.GoogleExt) {

				if customValidate(r) {

					switch true {

					case customShowHost:
						fmt.Println(r.Host)

					case customShowPath:
						fmt.Println(r.Path)

					case customShowSub:
						fmt.Println(r.Subdomain)

					default:
						fmt.Println(r.URL)
					}
				}
			}
		}()

		start.Sleep(rootCmd)
	}

	wg.Wait()
}

func customValidate(d extract.Data) bool {

	if d.URL != "" {

		if !slices.Contains(results, d.URL) {

			if rootVerify {

				if validate.Verify(d.URL) {

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

	customCmd.Flags().StringVarP(&customCommand, "command", "c", "", "dork command to scrape")
	customCmd.Flags().StringVarP(&customTargetHost, "target-host", "t", "", "filter result by host")
	customCmd.Flags().StringVar(&customSpiltBy, "split-by", ",", "dork commands split character")
	customCmd.Flags().IntVar(&customDepth, "depth", 1, "number of pages to scrape")
	customCmd.Flags().BoolVar(&customShowHost, "show-host", false, "show hosts as result")
	customCmd.Flags().BoolVar(&customShowSub, "show-sub", false, "show subdomains as result")
	customCmd.Flags().BoolVar(&customShowPath, "show-path", false, "show paths as result")

	customCmd.MarkFlagRequired("command")
	customCmd.MarkFlagsMutuallyExclusive("show-host", "show-sub", "show-path")
}
