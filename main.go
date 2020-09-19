package main

import (
	"encoding/json"
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
	str, _ := json.Marshal(user)
	fmt.Println(string(str))

}

func payload(nim string) io.Reader {
	p := "NICitb=" + nic + "&uid=" + nim
	return strings.NewReader(p)
}

type user struct {
	Username string `json:"username"`
	NimTPB   string
	Nim      string
	Nama     string
	Status   string
	Fakultas string
	Jurusan  string
	EmailITB string
	Email    string
}

func extract(html string) user {
	reg := regexp.MustCompile(`placeholder="(.*?)"`)
	match := reg.FindAllStringSubmatch(html, -1)

	nimTPB, nim := cleanNIM(match[2][1])
	fakultas, jurusan := cleanJurusan(match[5][1])
	emailITB := cleanEmail(match[6][1])
	email := cleanEmail(match[7][1])

	return user{
		Username: match[1][1],
		NimTPB:   nimTPB,
		Nim:      nim,
		Nama:     match[3][1],
		Status:   match[4][1],
		Fakultas: fakultas,
		Jurusan:  jurusan,
		EmailITB: emailITB,
		Email:    email,
	}
}

func cleanNIM(str string) (string, string) {
	list := strings.Split(str, ",")
	nimTPB := strings.TrimSpace(list[0])
	nim := strings.TrimSpace(list[1])

	return nimTPB, nim
}

func cleanJurusan(str string) (string, string) {
	list := strings.Split(str, "-")
	fakultas := strings.TrimSpace(list[0])
	jurusan := strings.TrimSpace(list[1])

	return fakultas, jurusan
}

func cleanEmail(email string) string {
	email = strings.ReplaceAll(email, "(dot)", ".")
	email = strings.ReplaceAll(email, "(at)", "@")

	return email
}
