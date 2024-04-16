package netops

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// FetchURL retrieves the content from the specified URL.
func FetchURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch URL %s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read content from %s: %s\n", url, err)
		return
	}
	fmt.Printf("Content from %s:\n%s\n", url, string(body))
}
