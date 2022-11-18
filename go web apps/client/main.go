package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	response, err := http.Get("http://localhost:8080")

	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// URL key-value form
	response, err = http.PostForm("http://localhost/url", url.Values{"name": {"Christian"}})
	if err != nil {
		return
	}

	// with body in json: {"name": {"Christian"}}
	type Ninja struct {
		name string
	}
	chris := Ninja{"Christian"}
	chrisJson, _ := json.Marshal(chris)

	response, err = http.Post("http://localhost:8080/body", "application/json", bytes.NewBuffer(chrisJson))
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
