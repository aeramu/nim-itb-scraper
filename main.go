package main

import (
	"fmt"
	"strconv"
	"time"
)

var url = "https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user"

//paste your cookie here, (19-09-2020) ci_session and ITBnic)
var session = "hcpn1ljlqvs20sktd3b560uejj1v61d4"
var nic = "ff2f1058a1f91f384f38f9af83b2bef2"
var cookie = "ITBnic=" + nic + "; ci_session=" + session

//depends on your internet speed. Less delay, more fast data scraping,
//but if your internet not good enough, it will crash
var delay = 4

func main() {
	nimChan := make(chan string)
	resChan := make(chan []byte)
	userChan := make(chan user)

	go func() {
		for nim := 1651800; nim <= 1651870; nim++ {
			fmt.Println(nim)
			for i := 0; i < 10; i++ {
				nimChan <- strconv.Itoa(nim) + strconv.Itoa(i)
			}
			time.Sleep(time.Second * time.Duration(delay))
		}
	}()

	//ioutil.WriteFile("out.html", <-res, 0644)

	var count int
	for {
		select {
		case nim := <-nimChan:
			go fetch(resChan, nim)
		// case eval := <-evalChan:
		// 	go evaluation(eval, )
		case res := <-resChan:
			go extract(userChan, string(res))
		case user := <-userChan:
			count++
			println(count)
			go fmt.Println(user)
		}
	}
	//TODO: write to file or database
	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadString('\n')
}

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
