package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var url = "https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user"

//paste your cookie here, (19-09-2020) ci_session and ITBnic)
var session = "hcpn1ljlqvs20sktd3b560uejj1v61d4"
var nic = "ff2f1058a1f91f384f38f9af83b2bef2"
var cookie = "ITBnic=" + nic + "; ci_session=" + session

//depends on your internet speed. Less delay, more fast data scraping,
//but if your internet not good enough, it will crash
var delay = 5

var code = []string{
	"160", // FMIPA
	"161", // SITH
	"162", // SF
	"163", // FITB
	"164", // FTTM
	"165", // STEI
	"166", // FTSL
	"167", // FTI
	"168", // FSRD
	"169", // FTMD
	"190", // SBM
	"199", // SAPPK
}

var w sync.WaitGroup

func main() {
	//TODO: add iteration to fakultas dan angkatan for range code { for i:= 14->20 {}}
	resChan := make(chan []byte)
	go func() {
		for nim := 1651800; nim <= 1651870; nim++ {
			fmt.Println(nim)
			for i := 0; i < 10; i++ {
				go sendRequest(resChan, strconv.Itoa(nim)+strconv.Itoa(i))
			}
			time.Sleep(time.Second * time.Duration(delay))
		}
	}()

	//ioutil.WriteFile("out.html", <-res, 0644)
	userChan := make(chan user)
	var count int
	for {
		select {
		case res := <-resChan:
			go extract(userChan, string(res))
		case user := <-userChan:
			count++
			println(count)
			go fmt.Println(user)
		}
	}

	// go func() {
	// 	for res := range resChan {
	// 		count++
	// 		println(count)
	// 		extract(userChan, string(res))
	// 	}
	// }()

	// fmt.Println("printing")
	// for user := range userChan {
	// 	fmt.Println(user)
	// }
	//TODO: write to file or database
	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadString('\n')
}

func sendRequest(c chan<- []byte, nim string) {
	req, _ := http.NewRequest("POST", url, payload(nim))
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	c <- body
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

func extract(c chan user, html string) {
	reg := regexp.MustCompile(`placeholder="(.*?)"`)
	match := reg.FindAllStringSubmatch(html, -1)

	if len(match) < 2 {
		return
	}

	nimTPB, nimJurusan := cleanNIM(match[2][1])
	fakultas, jurusan := cleanJurusan(match[5][1])
	emailITB := cleanEmail(match[6][1])
	email := cleanEmail(match[7][1])

	c <- user{
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
