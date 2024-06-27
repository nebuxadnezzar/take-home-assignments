package model

type (
	Meme struct {
		ID       string `json:"id"`
		URL      string `json:"url"`
		Name     string `json:"name"`
		Page     string `json:"page"`
		MD5      string `json:"md5"`
		FileSize string `json:"fileSize"`
		Text     string `json:"text"`
	}
	MemeRequest struct {
		Lon   float32 `query:"lon" json:"lon"`
		Lat   float32 `query:"lat" json:"lat"`
		Query string  `query:"query" json:"query"`
		Fuzzy bool    `query:"fuzzy" json:"fuzzy"`
	}
)
