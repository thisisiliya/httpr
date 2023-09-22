package start

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thisisiliya/httpr/pkg/extract"
	"github.com/thisisiliya/httpr/pkg/validate"
)

func Start(cmd *cobra.Command) {

	silent, _ := cmd.Flags().GetBool("silent")
	proxy, _ := cmd.Flags().GetString("proxy")

	if !silent {

		fmt.Fprintln(os.Stderr, BANNER)
	}

	if proxy != "" {

		if validate.Proxy(proxy) {

			extract.PROXY = proxy
		}
	}
}
