package pkg

import (
	"encoding/csv"
	"os"
)

var writer *csv.Writer

// NewWriter initialize csv writer
func NewWriter(filename string) {
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	writer = csv.NewWriter(file)
	writer.Write([]string{"NIM TPB", "NIM Jurusan", "Username", "Nama", "Status", "Fakultas", "Jurusan", "Email ITB", "Email"})
}

func save(u *user) {
	writer.Write(u.toCSV())
	writer.Flush()
}

func (u user) toCSV() []string {
	var s []string
	s = append(
		s,
		u.NimTPB,
		u.NimJurusan,
		u.Username,
		u.Nama,
		u.Status,
		u.Fakultas,
		u.Jurusan,
		u.EmailITB,
		u.Email,
	)
	return s
}
