package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	c := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, c)
	}
	for _, _ = range os.Args[1:] {
		resp := <-c
		fmt.Println(resp)
	}
}

func fetch(url string, c chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprint(err)
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c <- fmt.Sprint(err)
		return
	}
	res.Body.Close()
	end := time.Now()
	elapsed := end.Sub(start).Seconds()
	c <- fmt.Sprintf("Fetching %s took %0.2f seconds, length: %d\n", url, elapsed, len(b))
}
