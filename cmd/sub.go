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

var (
	mainDomain    string
	checkedResult []string

	subCmd = &cobra.Command{
		Use:     "sub",
		Short:   "algorithmic subdomain enumeration for domains",
		Long:    "algorithmic subdomain enumeration for domain(s)",
		Example: "httpr sub --domain hackerone.com --all",
		Run:     SubExt,
	}
)

func SubExt(_ *cobra.Command, _ []string) {

	utils.Start(root_Silent)

	o.MinDelay = i.root_MinDelay
	o.MaxDelay = i.root_MaxDelay

	o.Browser = request.Browser(root_Proxy, root_Timeout, root_Chromium)
	defer o.Browser.MustClose()

	o.Dork.Wildcard = true

	switch {

	case i.sub_Domain != "":
		mainDomain = i.sub_Domain
		o.Dork.Domain = mainDomain
		SubEnum()

	case i.sub_Domains != "":

		sub_domains, err := file.ReadByLine(i.sub_Domains)
		errors.Check(err)

		for _, domain := range *sub_domains {

			mainDomain = domain
			o.Dork.Domain = mainDomain
			SubEnum()
		}
	}
}

func SubEnum() {

	var blockKeys [][]string

	SubScrape()

	for {

		dividedKeys := DivideKeys(results)

		if len(blockKeys) != len(dividedKeys) {

			for len(dividedKeys) != len(blockKeys) {

				blockKeys = append(blockKeys, []string{})
			}

			for i, v := range dividedKeys {

				if len(blockKeys[i]) != len(v) {

					o.Dork.Block = v
					SubScrape()
				}
			}

			blockKeys = dividedKeys
		} else {

			break
		}
	}

	if i.sub_All {

		for _, subdomain := range results {

			if !slices.Contains(checkedResult, subdomain) {

				o.Dork.Domain = subdomain + "." + mainDomain
				checkedResult = append(checkedResult, subdomain)
				SubEnum()
			}
		}
	}
}

func SubScrape() {

	o.Dork.Page = 0
	var data []request.Data

	for o.Dork.Page < i.sub_Depth {

		data = *request.Scrape(&o)

		for _, d := range data {

			if SubValidate(&d) {

				switch true {

				case i.sub_ShowSub:
					fmt.Println(d.Subdomain)

				case i.sub_ShowURL:
					fmt.Println(d.URL)

				default:
					fmt.Println(d.Host)
				}
			}
		}

		o.Dork.Page++
	}
}

func SubValidate(d *request.Data) bool {

	if d.Subdomain != "" {

		if !slices.Contains(results, d.Subdomain) {

			if i.root_Verify {

				if request.Verify(d.URL) {

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

func DivideKeys(list []string) [][]string {

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

func init() {

	rootCmd.AddCommand(subCmd)

	subCmd.Flags().StringVarP(&i.sub_Domain, "domain", "d", "", "target domain to search subdomains")
	subCmd.Flags().StringVar(&i.sub_Domains, "domains", "", "target domains file path")
	subCmd.Flags().IntVar(&i.sub_Depth, "depth", 5, "number of pages to scrape per result")
	subCmd.Flags().BoolVarP(&i.sub_All, "all", "a", false, "redo the process for the result")
	subCmd.Flags().BoolVar(&i.sub_ShowSub, "show-sub", false, "show subdomains as result")
	subCmd.Flags().BoolVar(&i.sub_ShowURL, "show-url", false, "show URLs as result")

	subCmd.MarkFlagsOneRequired("domain", "domains")
	subCmd.MarkFlagsMutuallyExclusive("domain", "domains")
	subCmd.MarkFlagsMutuallyExclusive("show-sub", "show-url")
	subCmd.MarkFlagDirname("domains")
}
