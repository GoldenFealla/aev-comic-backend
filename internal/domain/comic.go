package domain

type Comic struct {
	Code string `json:"code" db:"code"`
	Name string `json:"name" db:"name"`
}

type ComicImage struct {
	Number int    `json:"number" db:"number"`
	Code   string `json:"code" db:"code"`
	Url    string `json:"url" db:"url"`
}
