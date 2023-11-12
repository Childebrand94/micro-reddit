package models

import (
	"time"
)

type Post struct {
	ID        int64     `db:"id"         json:"id"`
	Author_ID int64     `db:"author_id"  json:"authorId"`
	Title     string    `db:"title"      json:"title"`
	URL       string    `db:"url"        json:"url"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
	Vote      int       `                json:"upVotes"`
}

type PostResponse struct {
	Post
	UsersVoteStatus VoteStatus    `json:"usersVoteStatus"`
	Author          Author        `json:"author"`
	Comments        []CommentResp `json:"comments"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
}

type PostWithAuthor struct {
	Post
	UsersVoteStatus VoteStatus `json:"usersVoteStatus"`
	Author          Author     `json:"author"`
}

type VoteStatus string

const (
	Upvote   VoteStatus = "upvote"
	Downvote VoteStatus = "downvote"
	NoVote   VoteStatus = "noVote"
)
