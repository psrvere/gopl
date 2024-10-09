package chapter1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fetch() {
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "https://")
		if !hasPrefix {
			url = "https://" + url
		}

		// make http call
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "status: %v, fetch: %v\n", resp.Status, err)
			os.Exit(1)
		}
		// Approach 1
		// io.ReadAll reads till EOF i.e. reads everything in one shot
		// It required large enough buffer to hold the entire stream data
		// b, err := io.ReadAll(resp.Body)

		// Approach 2
		// io.Copy copies data from stream (source) to destination in chunks
		// Better than io.ReadAll
		n, err := io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "status: %v, fetch: reading %s: %v\n", resp.Status, url, err)
			os.Exit(1)
		}
		fmt.Printf("n: %v\n", n)
	}
}
