package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var url = "https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user"

//paste your cookie here, (19-09-2020) ci_session and ITBnic)
var session = "hcpn1ljlqvs20sktd3b560uejj1v61d4"
var nic = "ff2f1058a1f91f384f38f9af83b2bef2"
var cookie = "ITBnic=" + nic + "; ci_session=" + session

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

func main() {
	file, _ := os.OpenFile("mahasiswaitb.csv", os.O_CREATE|os.O_WRONLY, 0777)
	writer := csv.NewWriter(file)

	nimChan := make(chan string)
	resChan := make(chan []byte)
	userChan := make(chan *user)
	successChan := make(chan bool)

	// ioutil.WriteFile("out.html", body, 0644)

	go manager(successChan, nimChan)

	for {
		select {
		case nim := <-nimChan:
			go fetch(resChan, nim)
		case res := <-resChan:
			go extract(userChan, string(res))
		case user := <-userChan:
			if user == nil {
				go fmt.Println("fail")
				successChan <- false
			} else {
				successChan <- true
				fmt.Println(user)
				writer.Write(user.toCSV())
				writer.Flush()
			}
		}
	}
}
