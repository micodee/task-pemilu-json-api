package models

type Response struct {
	Wilayah   string            `json:"wilayah"`
	Perolehan PerolehanResponse `json:"perolehan"`
}

type PerolehanResponse struct {
	Partai     string `json:"partai"`
	TotalSuara int    `json:"total_suara"`
}
