package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thisisiliya/go_utils/errors"
	"github.com/thisisiliya/go_utils/file"
	"github.com/thisisiliya/httpr/pkg/request"
	"github.com/thisisiliya/httpr/pkg/utils"
	"golang.org/x/exp/slices"
)

var pathCmd = &cobra.Command{
	Use:     "path",
	Short:   "path enumeration for domains",
	Long:    "path enumeration for domain(s)",
	Example: "usage: --domain www.google.com --depth 10",
	Run:     PathExt,
}

func PathExt(_ *cobra.Command, _ []string) {

	utils.Start(root_Silent)

	o.MinDelay = i.root_MinDelay
	o.MaxDelay = i.root_MaxDelay

	o.Browser = request.Browser(root_Proxy, root_Timeout, root_Chromium)
	defer o.Browser.MustClose()

	switch {

	case i.path_Domain != "":
		o.Dork.Domain = i.path_Domain
		PathEnum()

	case i.path_Domains != "":
		path_domains, err := file.ReadByLine(i.path_Domains)
		errors.Check(err)

		for _, domain := range *path_domains {

			o.Dork.Domain = domain
			PathEnum()
		}
	}
}

func PathEnum() {

	o.Dork.Page = 0
	var data []request.Data

	for o.Dork.Page < i.path_Depth {

		data = append(data, *request.Scrape(&o)...)
		o.Dork.Page++
	}

	for _, d := range data {

		if PathValidate(&d) {

			switch true {

			case i.path_ShowPath:
				fmt.Println(d.Path)

			default:
				fmt.Println(d.URL)
			}
		}
	}
}

func PathValidate(data *request.Data) bool {

	if data.Host == o.Dork.Domain || data.Host == "www."+o.Dork.Domain {

		if !slices.Contains(results, data.Path) {

			if i.root_Verify {

				if request.Verify(data.URL) {

					results = append(results, data.Path)
					return true
				}
			} else {

				results = append(results, data.Path)
				return true
			}
		}
	}

	return false
}

func init() {

	rootCmd.AddCommand(pathCmd)

	pathCmd.Flags().StringVarP(&i.path_Domain, "domain", "d", "", "target domain to search")
	pathCmd.Flags().StringVar(&i.path_Domains, "domains", "", "target domains file path")
	pathCmd.Flags().IntVar(&i.path_Depth, "depth", 10, "number of pages to scrape per domain")
	pathCmd.Flags().BoolVar(&i.path_ShowPath, "show-path", false, "show paths as result")

	pathCmd.MarkFlagsOneRequired("domain", "domains")
	pathCmd.MarkFlagsMutuallyExclusive("domain", "domains")
	pathCmd.MarkFlagDirname("domains")
}
