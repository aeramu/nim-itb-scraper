package pkg

import (
	"regexp"
	"strings"
)

func extract(html string, c chan *user) {
	reg := regexp.MustCompile(`placeholder="(.*?)"`)
	match := reg.FindAllStringSubmatch(html, -1)

	if len(match) < 2 {
		c <- nil
		return
	}

	nimTPB, nimJurusan := cleanNIM(match[2][1])
	fakultas, jurusan := cleanJurusan(match[5][1])
	emailITB := cleanEmail(match[6][1])
	email := cleanEmail(match[7][1])

	c <- &user{
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
