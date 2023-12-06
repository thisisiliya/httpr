package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thisisiliya/httpr/pkg/utils"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display httpr's version and exit",
	Run: func(_ *cobra.Command, _ []string) {

		fmt.Fprintln(os.Stderr, utils.VERSION)
	},
}

func init() {

	rootCmd.AddCommand(versionCmd)
}
