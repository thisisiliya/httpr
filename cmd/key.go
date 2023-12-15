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

var keyCmd = &cobra.Command{
	Use:     "key",
	Short:   "keywords enumeration for domains",
	Long:    "keyword(s) enumeration for domain(s)",
	Example: "httpr key --domain hackerone.com --keyword report --depth 3",
	Run:     KeyExt,
}

func KeyExt(_ *cobra.Command, _ []string) {

	utils.Start(root_Silent)

	o.MinDelay = i.root_MinDelay
	o.MaxDelay = i.root_MaxDelay

	o.Browser = request.Browser(root_Proxy, root_Timeout, root_Chromium)
	defer o.Browser.MustClose()

	switch {

	case i.key_Domain != "":
		o.Dork.Domain = i.key_Domain

		switch {

		case i.key_Keyword != "":
			o.Dork.Key = i.key_Keyword
			KeyEnum()

		case i.key_Keywords != "":
			keywords, err := file.ReadByLine(i.key_Keywords)
			errors.Check(err)

			for _, key := range *keywords {

				o.Dork.Key = key
				KeyEnum()
			}
		}

	case i.key_Domains != "":
		switch {

		case i.key_Keyword != "":
			o.Dork.Key = i.key_Keyword

			key_domains, err := file.ReadByLine(i.key_Domains)
			errors.Check(err)

			for _, domain := range *key_domains {

				o.Dork.Domain = domain
				KeyEnum()
			}

		case i.key_Keywords != "":
			key_domains, err := file.ReadByLine(i.key_Domains)
			errors.Check(err)

			for _, domain := range *key_domains {

				o.Dork.Domain = domain

				key_keywords, err := file.ReadByLine(i.key_Domains)
				errors.Check(err)

				for _, key := range *key_keywords {

					o.Dork.Key = key
					KeyEnum()
				}
			}
		}
	}
}

func KeyEnum() {

	o.Dork.Page = 0
	var data []request.Data

	for o.Dork.Page < i.key_Depth {

		data = append(data, *request.Scrape(&o)...)
		o.Dork.Page++
	}

	for _, d := range data {

		if KeyValidate(&d) {

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

func KeyValidate(data *request.Data) bool {

	if data.Host == o.Dork.Domain || data.Host == "www."+o.Dork.Domain {

		if !slices.Contains(results, data.URL) {

			if i.root_Verify {

				if request.Verify(data.URL) {

					results = append(results, data.URL)
					return true
				}
			} else {

				results = append(results, data.URL)
				return true
			}
		}
	}

	return false
}

func init() {

	rootCmd.AddCommand(keyCmd)

	keyCmd.Flags().StringVarP(&i.key_Domain, "domain", "d", "", "target domain to search keyword(s)")
	keyCmd.Flags().StringVar(&i.key_Domains, "domains", "", "target domains file path")
	keyCmd.Flags().StringVarP(&i.key_Keyword, "keyword", "k", "", "target keyword to search")
	keyCmd.Flags().StringVar(&i.key_Keywords, "keywords", "", "target keywords path")
	keyCmd.Flags().IntVar(&i.key_Depth, "depth", 3, "number of pages to scrape per key")
	keyCmd.Flags().BoolVar(&i.key_ShowHost, "show-host", false, "show hosts as result")
	keyCmd.Flags().BoolVar(&i.key_ShowSub, "show-sub", false, "show subdomains as result")
	keyCmd.Flags().BoolVar(&i.key_ShowPath, "show-path", false, "show paths as result")

	keyCmd.MarkFlagsOneRequired("domain", "domains")
	keyCmd.MarkFlagsMutuallyExclusive("domain", "domains")
	keyCmd.MarkFlagsOneRequired("keyword", "keywords")
	keyCmd.MarkFlagsMutuallyExclusive("keyword", "keywords")
	keyCmd.MarkFlagsMutuallyExclusive("show-host", "show-sub", "show-path")
	keyCmd.MarkFlagDirname("domains")
	keyCmd.MarkFlagDirname("keywords")
}
