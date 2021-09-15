package model

type Book struct {
	Id          uint   `gorm:"primary_key" json:"id" example:"1"`
	Title       string `gorm:"varchar(100), not null" json:"title" example:"Games of thrones"`
	ReleaseYear int    `gorm:"int, not null" json:"release_year" example:"1997"`
	AuthorId    uint   `json:"author_id"`
}
