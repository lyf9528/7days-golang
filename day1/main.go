package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path=%s\n", r.URL)
	})
	r.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}
