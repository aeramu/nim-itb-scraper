package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user"

//paste your cookie here, format(19-09-2020): ci_session=xxxxxx)
var cookie = "ci_session=4uimcpvsdfkad4eealjgfpn255fj9dh3"

func main() {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", cookie)
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
