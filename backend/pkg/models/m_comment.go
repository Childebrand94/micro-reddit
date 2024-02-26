package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type CommentRequest struct {
	Post_ID   int64  ` json:"postId"`
	Parent_ID int64  ` json:"parentID"`
	Message   string ` json:"message"`
	Path      string `json:"path"`
}

type Comment struct {
	ID         int64         `db:"id"         json:"id"`
	Post_ID    int64         `db:"post_id"    json:"postId"`
	Author_ID  int64         `db:"author_id"  json:"authorId"`
	Parent_ID  sql.NullInt64 `db:"parent_id"  json:"-"`
	Message    string        `db:"message"    json:"message"`
	Vote       int           `                json:"upVotes"`
	Created_at time.Time     `db:"created_at" json:"createdAt"`
	Path       string        `db:"path" json:"path"`
	Children   []CommentResp `json:"childComments"`
}

type CommentResp struct {
	Comment
	UsersVoteStatus VoteStatus `json:"usersVoteStatus"`
	Author          `json:"author"`
}

type CommentRespAlias CommentResp

func (cr *CommentResp) MarshalJSON() ([]byte, error) {
	var parentID interface{}

	if cr.Parent_ID.Valid {
		parentID = cr.Parent_ID.Int64
	} else {
		parentID = nil
	}

	commentRespAlias := CommentRespAlias(*cr)

	return json.Marshal(&struct {
		*CommentRespAlias
		Parent_ID interface{} `json:"parentID"`
	}{
		CommentRespAlias: &commentRespAlias,
		Parent_ID:        parentID,
	})
}
