package cmd

import (
	"fmt"
	"os"

	"github.com/thisisiliya/httpr/pkg/start"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display httpr's version and exit",
	Run: func(_ *cobra.Command, _ []string) {

		fmt.Fprintln(os.Stderr, start.VERSION)
	},
}

func init() {

	rootCmd.AddCommand(versionCmd)
}
