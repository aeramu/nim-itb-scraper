package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user"

//paste your cookie here, format(19-09-2020): ci_session=xxxxxx)
var session = "hcpn1ljlqvs20sktd3b560uejj1v61d4"
var nic = "ff2f1058a1f91f384f38f9af83b2bef2"
var cookie = "ITBnic=" + nic + "; ci_session=" + session
var payload = "NICitb=" + nic + "&uid=18119004"

func main() {
	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(payload))
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
	ioutil.WriteFile("out.txt", body, 0644)
}
