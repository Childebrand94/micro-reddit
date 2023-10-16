package models

import "time"

type Comment struct {
	ID         int64     `db:"id" json:"ID"`
	Post_ID    int64     `db:"post_id" json:"postID"`
	Author_ID  int64     `db:"author_id" json:"authorID"`
	Parent_ID  *int64    `db:"parent_id" json:"parentID"`
	Message    string    `db:"message" json:"message"`
	Created_at time.Time `db:"created_at" json:"created_at"`
}
