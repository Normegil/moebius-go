package utils

import (
	"io/ioutil"
	"net/http"
)

// Fetcher will fetch data from chosen source and return a JSON as a byte array
type Fetcher interface {
	Fetch(url string) ([]byte, error)
}

// HTTPFetcher implement Fetcher interface and get it's data throught http requests
type HTTPFetcher struct{}

// Fetch make a http request and load data from the network, using "net/http"
func (fetcher HTTPFetcher) Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}
	return content, nil
}
