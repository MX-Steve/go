package myhttp

// http funcs
import (
	"net/http"
	"io/ioutil"
)

func getHTTPRespWithRetry(URL string) ([]byte, error) {
	times := 0
	var lastError error
	for times < 3 {
		times = times + 1
		resp, err := http.Get(URL)
		if err != nil {
			lastError = err
			continue
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			lastError = err
			continue
		}
		return body, nil
	}

	return nil, lastError
}