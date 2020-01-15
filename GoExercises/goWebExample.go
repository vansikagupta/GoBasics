package main

import (
	"fmt"
	"net/http"
)

func index_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello! This is the index page")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":9000", nil)
}
