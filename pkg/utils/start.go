package utils

import (
	"fmt"
	"os"
)

func Start(silent bool) {

	if !silent {

		fmt.Fprintln(os.Stderr, BANNER)
	}
}
