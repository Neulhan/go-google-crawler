package models

type Data struct {
	Name     string   `json:"name"`
	Img      string   `json:"img"`
	Rating   string   `json:"rating"`
	RateNum  string   `json:"rateNum"`
	HashTags []string `json:"hashTags"`
}
