package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var url = "https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user"

//paste your cookie here, (19-09-2020) ci_session and ITBnic)
var session = "hcpn1ljlqvs20sktd3b560uejj1v61d4"
var nic = "ff2f1058a1f91f384f38f9af83b2bef2"
var cookie = "ITBnic=" + nic + "; ci_session=" + session

func main() {
	for nim := 18119001; nim <= 18119037; nim++ {
		go exec(nim)
	}
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func exec(nim int) {
	body := sendRequest(strconv.Itoa(nim))
	//ioutil.WriteFile("out.html", body, 0644)

	extract(string(body))
	fmt.Println(nim)
}

func sendRequest(nim string) []byte {
	req, _ := http.NewRequest("POST", url, payload(nim))
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func payload(nim string) io.Reader {
	p := "NICitb=" + nic + "&uid=" + nim
	return strings.NewReader(p)
}

type user struct {
	Username   string `json:"username"`
	NimTPB     string `json:"nim_tpb"`
	NimJurusan string `json:"nim_jurusan"`
	Nama       string `json:"nama"`
	Status     string `json:"status"`
	Fakultas   string `json:"fakultas"`
	Jurusan    string `json:"jurusan"`
	EmailITB   string `json:"email itb"`
	Email      string `json:"email"`
}

func extract(html string) user {
	reg := regexp.MustCompile(`placeholder="(.*?)"`)
	match := reg.FindAllStringSubmatch(html, -1)

	nimTPB, nimJurusan := cleanNIM(match[2][1])
	fakultas, jurusan := cleanJurusan(match[5][1])
	emailITB := cleanEmail(match[6][1])
	email := cleanEmail(match[7][1])

	return user{
		Username:   match[1][1],
		NimTPB:     nimTPB,
		NimJurusan: nimJurusan,
		Nama:       match[3][1],
		Status:     match[4][1],
		Fakultas:   fakultas,
		Jurusan:    jurusan,
		EmailITB:   emailITB,
		Email:      email,
	}
}

func cleanNIM(str string) (string, string) {
	list := strings.Split(str, ",")
	nim1 := strings.TrimSpace(list[0])
	nim2 := ""
	if len(list) > 1 {
		nim2 = strings.TrimSpace(list[1])
	}

	return nim1, nim2
}

func cleanJurusan(str string) (string, string) {
	list := strings.Split(str, "-")
	fakultas := strings.TrimSpace(list[0])
	jurusan := ""
	if len(list) > 1 {
		jurusan = strings.TrimSpace(list[1])
	}

	return fakultas, jurusan
}

func cleanEmail(email string) string {
	email = strings.ReplaceAll(email, "(dot)", ".")
	email = strings.ReplaceAll(email, "(at)", "@")

	return email
}
