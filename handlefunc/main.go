package main

import (
	"fmt"
	"net/http"
)

func main() {

	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, http!!")
	}

	http.HandleFunc("/", f)
	http.ListenAndServe(":8080", nil)
}
