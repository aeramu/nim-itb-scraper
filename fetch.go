package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var client = http.Client{}

func fetch(resChan chan<- []byte, nim string) {
	req, _ := http.NewRequest("POST", url, payload(nim))
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	resChan <- body
}

func payload(nim string) io.Reader {
	p := "NICitb=" + nic + "&uid=" + nim
	return strings.NewReader(p)
}
