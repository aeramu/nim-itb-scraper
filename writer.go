package main

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
