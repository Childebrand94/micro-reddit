package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Post struct {
	ID        int64       `db:"id"         json:"id"`
	Author_ID int64       `db:"author_id"  json:"authorId"`
	Title     string      `db:"title"      json:"title"`
	URL       string      `db:"url"        json:"url"`
	CreatedAt time.Time   `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time   `db:"updated_at" json:"updatedAt"`
	Vote      pgtype.Int8 `                json:"upVotes"`
}

type PostResponse struct {
	Post
	Author   Author
	Comments []CommentResp
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
}
