package models

import (
	"time"
)

type User struct {
	ID         int64     `db:"id"            json:"id"`
	First_name string    `db:"first_name"    json:"firstName"`
	Last_name  string    `db:"last_name"     json:"lastName"`
	Username   string    `db:"username"      json:"username"`
	Email      string    `db:"email"         json:"email"`
	Password   string    `db:"password"      json:"password"`
	DateJoined time.Time `db:"registered_at" json:"dateJoined"`
}

type UserResp struct {
	ID         int64     `db:"id"            json:"id"`
	First_name string    `db:"first_name"    json:"firstName"`
	Last_name  string    `db:"last_name"     json:"lastName"`
	Username   string    `db:"username"      json:"username"`
	Email      string    `db:"email"         json:"email"`
	DateJoined time.Time `db:"registered_at" json:"dateJoined"`
}
type Session struct {
	Session_id string
	User_id    int64
}
type UserPoints struct {
	PostCount        int `db:"post_count" json:"postCount"`
	PostUpVotes      int `db:"post_up_votes" json:"postUpVotes"`
	PostDownVotes    int `db:"post_down_votes" json:"postDownVotes"`
	CommentUpVotes   int `db:"comment_up_votes" json:"commentUpVotes"`
	CommentDownVotes int `db:"comment_down_votes" json:"commentDownVotes"`
	Karma            int `db:"karma" json:"karma"`
}

// type UserResp struct {
// 	User
// 	Posts    []PostWithAuthor `json:"posts"`
// 	Comments []Comment        `json:"comments"`
// }
