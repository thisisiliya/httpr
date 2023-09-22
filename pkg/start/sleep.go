package start

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func Sleep(cmd *cobra.Command) {

	var (
		minDelay    = 1
		maxDelay, _ = cmd.Flags().GetInt("delay")
		randomDelay = rand.Intn(maxDelay-minDelay) + minDelay
	)

	time.Sleep(time.Duration(randomDelay) * time.Second)
}
