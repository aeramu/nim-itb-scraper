package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var url = "https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user"

//paste your cookie here, (19-09-2020) ci_session and ITBnic)
var session = "hcpn1ljlqvs20sktd3b560uejj1v61d4"
var nic = "ff2f1058a1f91f384f38f9af83b2bef2"
var nim = "18118041"
var cookie = "ITBnic=" + nic + "; ci_session=" + session

func main() {
	req, _ := http.NewRequest("POST", url, payload(nim))
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	ioutil.WriteFile("out.html", body, 0644)

	user := extract(string(body))
	fmt.Println(user)

}

var html = `placeholder="(.*?)"`

func payload(nim string) io.Reader {
	p := "NICitb=" + nic + "&uid=" + nim
	return strings.NewReader(p)
}

func extract(html string) model.user {
	reg := regexp.MustCompile(`placeholder="(.*?)"`)
	match := reg.FindAllStringSubmatch(html, -1)
	return model.user{
		username: match[1][1],
		nimTPB:   match[2][1][:8],
		nim:      match[2][1][10:],
		nama:     match[3][1],
		status:   match[4][1],
		fakultas: match[5][1],
		jurusan:  match[5][1],
		emailITB: match[6][1],
		email:    match[7][1],
	}
}
