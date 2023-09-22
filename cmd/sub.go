package cmd

import (
	"fmt"

	"github.com/thisisiliya/httpr/pkg/engines"
	"github.com/thisisiliya/httpr/pkg/extract"
	"github.com/thisisiliya/httpr/pkg/start"
	"github.com/thisisiliya/httpr/pkg/validate"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var (
	subDomain  string
	subDomains string
	subDepth   int
	subAll     bool
	subShowSub bool
	subShowURL bool

	blockKeys   [][]string
	dividedKeys [][]string

	subCmd = &cobra.Command{
		Use:   "sub",
		Short: "algorithmic subdomain enumeration for domains",
		Long: ("\nalgorithmic subdomain enumeration for domain(s)" +
			"\nusage: -d google.com"),
		Run: subExt,
	}
)

func subExt(_ *cobra.Command, _ []string) {

	start.Start(rootCmd)
	opt.Wildcard = true

	switch {

	case subDomain != "":
		opt.Domain = subDomain
		subEnum()

	case subDomains != "":
		for _, domain := range extract.ReadFile(subDomains) {

			opt.Domain = domain
			subEnum()
		}
	}
}

func subEnum() {

	blockKeys = [][]string{}
	dividedKeys = [][]string{}

	subScrape()

	for {

		dividedKeys = divideKeys(results)

		if len(blockKeys) != len(dividedKeys) {

			for len(dividedKeys) != len(blockKeys) {

				blockKeys = append(blockKeys, []string{})
			}

			for i, v := range dividedKeys {

				if len(blockKeys[i]) != len(v) {

					opt.Block = v
					subScrape()
				}
			}

			blockKeys = dividedKeys
		} else {

			break
		}
	}
}

func subScrape() {

	opt.Page = 0

	for opt.Page < subDepth {

		wg.Add(1)

		go func() {

			defer wg.Done()

			URL = engines.GoogleURL(&opt)

			for _, r := range *extract.Scrape(URL, opt.Domain, rootRetry, extract.GoogleExt) {

				if subValidate(r) {

					switch true {

					case subShowSub:
						fmt.Println(r.Subdomain)

					case subShowURL:
						fmt.Println(r.URL)

					default:
						fmt.Println(r.Host)
					}
				}
			}
		}()

		opt.Page++
		start.Sleep(rootCmd)
	}

	wg.Wait()
}

func divideKeys(list []string) [][]string {

	n := 5
	r := [][]string{}
	c := len(list)/n + 1

	for i := 0; i < c; i++ {

		r = append(r, []string{})
	}

	for i := 0; i < len(list); i++ {

		r[i%c] = append(r[i%c], list[i])
	}

	return r
}

func subValidate(d extract.Data) bool {

	if d.Subdomain != "" {

		if !slices.Contains(results, d.Subdomain) {

			if rootVerify {

				if validate.Verify(d.URL) {

					results = append(results, d.Subdomain)
					return true
				}
			} else {

				results = append(results, d.Subdomain)
				return true
			}
		}
	}

	return false
}

func init() {

	rootCmd.AddCommand(subCmd)

	subCmd.Flags().StringVarP(&subDomain, "domain", "d", "", "target domain to search subdomains")
	subCmd.Flags().StringVar(&subDomains, "domains", "", "target domains file path")
	subCmd.Flags().IntVar(&subDepth, "depth", 5, "number of pages to scrape per result")
	subCmd.Flags().BoolVarP(&subAll, "all", "a", false, "redo the process for the result")
	subCmd.Flags().BoolVar(&subShowSub, "show-sub", false, "show subdomains as result")
	subCmd.Flags().BoolVar(&subShowURL, "show-url", false, "show URLs as result")

	subCmd.MarkFlagsMutuallyExclusive("domain", "domains")
	subCmd.MarkFlagsMutuallyExclusive("show-sub", "show-url")
}
