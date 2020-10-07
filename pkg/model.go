package pkg

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
