package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var counter = int(0)
var lock = new(sync.Mutex)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	increaseCounter()
}

func increaseCounter() {
	lock.Lock()
	counter++
	lock.Unlock()
}

func count(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	fmt.Fprintf(w, "Count = %d\n", counter)
	lock.Unlock()
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", count)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
