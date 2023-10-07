package models

type Post struct {
    ID       int64  `db:"id" json:"id"`
    UserID   int64  `db:"user_id" json:"userId"`
    Title    string `db:"title" json:"title"`
    Content  string `db:"content" json:"content"`
	URL string `db:"url" json:"url"`
}