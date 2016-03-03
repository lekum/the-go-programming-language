package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		res, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v fetching %s\n", err, arg)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", arg, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
