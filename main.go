package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/country", func(w http.ResponseWriter, req *http.Request) {

	})

	http.ListenAndServe(":8090", nil)
}
