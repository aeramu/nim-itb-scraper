package pkg

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var fetcher struct {
	client  *http.Client
	url     string
	session string
	nic     string
	cookie  string
}

//NewFetcher initialize fetch client
func NewFetcher(url, session, nic string) {
	fetcher.client = &http.Client{}
	fetcher.url = url
	fetcher.session = session
	fetcher.nic = nic
	fetcher.cookie = "ITBnic=" + nic + "; ci_session=" + session
}

func fetch(nim string, c chan *user) {
	req, _ := http.NewRequest("POST", fetcher.url, payload(nim))
	req.Header.Add("Cookie", fetcher.cookie)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := fetcher.client.Do(req)
	for err != nil {
		fmt.Println("trying again... " + nim)
		fmt.Println(err)
		res, err = fetcher.client.Do(req)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	extract(string(body), c)
}

func payload(nim string) io.Reader {
	p := "NICitb=" + fetcher.nic + "&uid=" + nim
	return strings.NewReader(p)
}
