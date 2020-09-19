package main

func (u user) toCSV() []string {
	var s []string
	s = append(
		s,
		u.Username,
		u.NimTPB,
		u.NimJurusan,
		u.Nama,
		u.Status,
		u.Fakultas,
		u.Jurusan,
		u.EmailITB,
		u.Email,
	)
	return s
}
