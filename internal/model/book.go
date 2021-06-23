package model

type Book struct {
	Id          int    `json:"id" example:"1"`
	Title       string `json:"title" example:"Games of thrones"`
	ReleaseYear int    `json:"release_year" example:"1997"`
	Author      Author `json:"author"`
}
