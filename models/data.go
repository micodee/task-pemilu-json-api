package models

type DapilData struct {
	Nama  string  `json:"nama"`
	Dapil []int64 `json:"dapil"`
}

type PartaiData struct {
	Warna       string `json:"warna"`
	IsAceh      bool   `json:"is_aceh"`
	IdPilihan   int    `json:"id_pilihan"`
	NomorUrut   int    `json:"nomor_urut"`
	Nama        string `json:"nama"`
	NamaLengkap string `json:"nama_lengkap"`
}

type DprData struct {
	Chart int `json:"chart"`
	Table int `json:"table"`
}
