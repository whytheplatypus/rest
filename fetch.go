package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return b, err
	}

	if resp.StatusCode != http.StatusOK {
		return b, fmt.Errorf("Failed to fetch from server, %s", resp.Status)
	}

	return b, nil
}
