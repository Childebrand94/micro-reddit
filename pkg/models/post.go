package models

import "time"

type Post struct {
	ID        int64  `db:"id" json:"id"`
	Author_ID int64  `db:"author_id" json:"authorID"`
	URL       string `db:"url" json:"url"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
type PostWithComments struct {
	PostData Post
	Comments []Comment
}	