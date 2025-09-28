package client

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("create request failed: %w", err)
	}

	req.Header = Headers
	req.Header.Set("Cookie", GetCookiesString())
	resp, err := httpClient.Do(req)

	if resp.StatusCode == 401 {
		err := fetch_cookies(httpClient)
		if err != nil {
			return nil, fmt.Errorf("request failed for session refresh: %w", err)
		}
		return nil, fmt.Errorf("refreshing session")
	}

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response failed: %w", err)
	}

	return body, nil
}
