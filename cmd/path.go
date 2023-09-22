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
	pathDomain   string
	pathDomains  string
	pathDepth    int
	pathShowPath bool

	pathCmd = &cobra.Command{
		Use:   "path",
		Short: "path enumeration for domains",
		Long: ("\npath enumeration for domain(s)" +
			"\nusage: -d www.google.com --depth 20"),
		Run: pathExt,
	}
)

func pathExt(_ *cobra.Command, _ []string) {

	start.Start(rootCmd)

	switch {

	case pathDomain != "":
		opt.Domain = pathDomain
		pathEnum()

	case pathDomains != "":
		for _, domain := range extract.ReadFile(pathDomains) {

			opt.Domain = domain
			pathEnum()
		}
	}
}

func pathEnum() {

	opt.Page = 0

	for opt.Page < pathDepth {

		wg.Add(1)

		go func() {

			defer wg.Done()

			URL = engines.GoogleURL(&opt)

			for _, r := range *extract.Scrape(URL, opt.Domain, rootRetry, extract.GoogleExt) {

				if pathValidate(r) {

					switch true {

					case pathShowPath:
						fmt.Println(r.Path)

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

func pathValidate(d extract.Data) bool {

	if d.Host == opt.Domain || d.Host == "www."+opt.Domain {

		if !slices.Contains(results, d.Path) {

			if rootVerify {

				if validate.Verify(d.URL) {

					results = append(results, d.Path)
					return true
				}
			} else {

				results = append(results, d.Path)
				return true
			}
		}
	}

	return false
}

func init() {

	rootCmd.AddCommand(pathCmd)

	pathCmd.Flags().StringVarP(&pathDomain, "domain", "d", "", "target domain to search")
	pathCmd.Flags().StringVar(&pathDomains, "domains", "", "target domains file path")
	pathCmd.Flags().IntVar(&pathDepth, "depth", 20, "number of pages to scrape per domain")
	pathCmd.Flags().BoolVar(&pathShowPath, "show-path", false, "show paths as result")

	pathCmd.MarkFlagsMutuallyExclusive("domain", "domains")
}
