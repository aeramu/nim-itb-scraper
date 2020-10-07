package main

import (
	"encoding/csv"
	"os"

	"github.com/aeramu/nim-itb-scraper/pkg"
)

//paste your cookie here, (19-09-2020) ci_session and ITBnic
var session = "REPLACE_THIS_WITH_YOUR_COPIED_VALUE"
var nic = "ff2f1058a1f91f384f38f9af83b2bef2"

var url = "https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user"

var startYear = 17
var endYear = 20

var filename = "itb.csv"

func init() {
	pkg.NewFetcher(url, session, nic)
}

func main() {
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	writer := csv.NewWriter(file)
	writer.Write([]string{"NIM TPB", "NIM Jurusan", "Username", "Nama", "Status", "Fakultas", "Jurusan", "Email ITB", "Email"})
	writer.Flush()

	dataCh := make(chan []string)
	go pkg.Scrape(faculty, startYear, endYear, dataCh)

	for data := range dataCh {
		writer.Write(data)
		writer.Flush()
	}

}

var faculty = []string{
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
	"197", // SBM
	"198", // SITH-R
	"199", // SAPPK
}
