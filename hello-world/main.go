package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Raspberry!\nThis was deployed asynchronously.\nWorks!")
}
