package models

type Post struct {
	ID        int64  `db:"id" json:"id"`
	Author_ID int64  `db:"author_id" json:"authorId"`
	Title     string `db:"title" json:"title"`
	URL       string `db:"url" json:"url"`
}
