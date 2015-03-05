package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", r.URL.Path)
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "client/views/error.html")
	} else {
		http.ServeFile(w, r, "client/views/index.html")
	}

}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}
