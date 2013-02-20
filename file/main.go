package main

import (
	"io"
	"os"
	"net/http"
)

func main() {
	tr := &http.Transport{}
	tr.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
	c := &http.Client{Transport: tr}
	r, _ := c.Get("file:///main.go")
	io.Copy(os.Stdout, r.Body)
}
