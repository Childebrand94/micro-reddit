package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID        int64 `db:"id"         json:"id"`
	Post_ID   int64 `db:"post_id"    json:"postId"`
	Author_ID int64 `db:"author_id"  json:"authorId"`
	// Parent_ID pgtype.Int8 `db:"parent_id"  json:"parentId"`
	Parent_ID  sql.NullInt64 `db:"parent_id"  json:"parentID"`
	Message    string        `db:"message"    json:"message"`
	Vote       int           `                json:"upVotes"`
	Created_at time.Time     `db:"created_at" json:"createdAt"`
	Path       string        `db:"path" json:"path"`
}

type CommentResp struct {
	Comment
	UsersVoteStatus VoteStatus `json:"usersVoteStatus"`
	Author          `json:"author"`
}
