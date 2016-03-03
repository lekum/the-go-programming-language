package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v fetching %s\n", err, url)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
