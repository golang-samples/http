package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := &Person{"Jonh", 27}
	buf, _ := xml.Marshal(person)
	body := bytes.NewBuffer(buf)
	r, _ := http.Post("http://localhost:8080", "text/xml", body)
	response, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(response))
}
