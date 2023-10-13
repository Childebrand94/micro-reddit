package models

type Comment struct {
	ID     int64  `db:"id" json:"id"`
	PostID int64  `db:"post_id" json:"postId"`
	UserID int64  `db:"user_id" json:"userId"`
	Text   string `db:"text" json:"text"`
}
