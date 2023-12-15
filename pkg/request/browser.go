package request

import (
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Browser(proxy string, timeout int, chromium bool) *rod.Browser {

	var (
		launch  = launcher.New()
		browser = rod.New()

		chrome_path = []string{
			`C:\Program Files\Google\Chrome\Application\chrome.exe`,        // Windows
			`/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`, // MacOS
			`/usr/bin/google-chrome`,                                       // Linux
			`/opt/google/chrome/chrome`,                                    // Linux
		}
	)

	if !chromium {

		for _, path := range chrome_path {

			if exists(path) {

				launch.Bin(path)
				break
			}
		}
	}

	if proxy != "" {

		launch.Set("proxy-server", proxy)
	}

	browser.Timeout(time.Duration(timeout) * time.Minute)
	browser.ControlURL(launch.MustLaunch())

	return browser.MustConnect()
}

func exists(file_path string) bool {

	if _, err := os.Stat(file_path); err == nil {

		return true
	}

	return false
}
