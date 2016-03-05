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
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %s\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form [%q] = %q\n", k, v)
	}
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
