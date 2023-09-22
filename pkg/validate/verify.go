package validate

import "net/http"

func Verify(URL string) bool {

	if resp, err := http.Head(URL); err == nil {

		if resp.StatusCode == 200 {

			return true
		}
	}

	return false
}
