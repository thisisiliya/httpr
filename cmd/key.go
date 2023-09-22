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
	keyDomain   string
	keyDomains  string
	keyKeyword  string
	keyKeywords string
	keyDepth    int
	keyShowHost bool
	keyShowPath bool
	keyShowSub  bool

	keyCmd = &cobra.Command{
		Use:   "key",
		Short: "keywords enumeration for domains",
		Long: ("\nkeyword(s) enumeration for domain(s)" +
			"\nusage: -d www.google.com -k exploit --depth 3"),
		Run: keyExt,
	}
)

func keyExt(_ *cobra.Command, _ []string) {

	start.Start(rootCmd)

	switch {

	case keyDomain != "":
		opt.Domain = keyDomain

		switch {

		case keyKeyword != "":
			opt.Key = keyKeyword
			keyEnum()

		case keyKeywords != "":
			for _, key := range extract.ReadFile(keyKeywords) {

				opt.Key = key
				keyEnum()
			}
		}

	case keyDomains != "":
		switch {

		case keyKeyword != "":
			opt.Key = keyKeyword

			for _, domain := range extract.ReadFile(keyDomains) {

				opt.Domain = domain
				keyEnum()
			}

		case keyKeywords != "":
			for _, domain := range extract.ReadFile(keyDomains) {

				opt.Domain = domain

				for _, key := range extract.ReadFile(keyKeywords) {

					opt.Key = key
					keyEnum()
				}
			}
		}
	}
}

func keyEnum() {

	opt.Page = 0

	for opt.Page < keyDepth {

		wg.Add(1)

		go func() {

			defer wg.Done()

			URL = engines.GoogleURL(&opt)

			for _, r := range *extract.Scrape(URL, opt.Domain, rootRetry, extract.GoogleExt) {

				if keyValidate(r) {

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

		opt.Page++
		start.Sleep(rootCmd)
	}

	wg.Wait()
}

func keyValidate(d extract.Data) bool {

	if d.Host == opt.Domain || d.Host == "www."+opt.Domain {

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

	rootCmd.AddCommand(keyCmd)

	keyCmd.Flags().StringVarP(&keyDomain, "domain", "d", "", "target domain to search keyword(s)")
	keyCmd.Flags().StringVar(&keyDomains, "domains", "", "target domains file path")
	keyCmd.Flags().StringVarP(&keyKeyword, "keyword", "k", "", "target keyword to search")
	keyCmd.Flags().StringVar(&keyKeywords, "keywords", "", "target keywords path")
	keyCmd.Flags().IntVar(&keyDepth, "depth", 3, "number of pages to scrape per key")
	keyCmd.Flags().BoolVar(&keyShowHost, "show-host", false, "show hosts as result")
	keyCmd.Flags().BoolVar(&keyShowSub, "show-sub", false, "show subdomains as result")
	keyCmd.Flags().BoolVar(&keyShowPath, "show-path", false, "show paths as result")

	keyCmd.MarkFlagsMutuallyExclusive("domain", "domains")
	keyCmd.MarkFlagsMutuallyExclusive("keyword", "keywords")
}
