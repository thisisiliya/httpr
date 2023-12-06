package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/thisisiliya/httpr/pkg/request/validate"
)

func Start(proxy string, silent bool) (context.Context, context.CancelFunc, context.CancelFunc) {

	var ctx context.Context
	var cancel1 context.CancelFunc
	var cancel2 context.CancelFunc

	if !silent {

		fmt.Fprintln(os.Stderr, BANNER)
	}

	if proxy != "" && validate.Proxy(proxy) {

		opts := append(chromedp.DefaultExecAllocatorOptions[:],

			chromedp.ProxyServer(proxy),
		)
		ctx, cancel1 = chromedp.NewExecAllocator(context.Background(), opts...)
	} else {

		ctx, cancel1 = chromedp.NewContext(context.Background())
	}

	ctx, cancel2 = context.WithTimeout(ctx, 10*time.Minute)
	return ctx, cancel1, cancel2
}
