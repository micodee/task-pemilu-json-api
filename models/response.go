package models

type Response struct {
	Wilayah   string              `json:"wilayah"`
	Perolehan []PerolehanResponse `json:"perolehan"`
}

type PerolehanResponse struct {
	Partai     string  `json:"partai"`
	TotalSuara float64 `json:"total_suara"`
}
