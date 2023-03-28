package models

type PartaiData struct {
	Warna       string `json:"warna"`
	IsAceh      bool   `json:"is_aceh"`
	IdPilihan   int    `json:"id_pilihan"`
	NomorUrut   int    `json:"nomor_urut"`
	Nama        string `json:"nama"`
	NamaLengkap string `json:"nama_lengkap"`
}