package main

import (
	"fmt"
)

var url = "https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user"

//paste your cookie here, (19-09-2020) ci_session and ITBnic)
var session = "hcpn1ljlqvs20sktd3b560uejj1v61d4"
var nic = "ff2f1058a1f91f384f38f9af83b2bef2"
var cookie = "ITBnic=" + nic + "; ci_session=" + session

func main() {
	nimChan := make(chan string)
	resChan := make(chan []byte)
	userChan := make(chan *user)
	successChan := make(chan bool)

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
				go fmt.Println(*user)
				successChan <- true
			}
		}
	}
	//TODO: write to file or database
	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadString('\n')
}
