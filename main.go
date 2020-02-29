package main

import (
	"ip_country/cmd/src"
	"net/http"
)

func main() {
	http.HandleFunc("/country", src.CheckIP)
	http.ListenAndServe(":8090", nil)
}
