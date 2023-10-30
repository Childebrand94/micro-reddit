package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Comment struct {
	ID        int64       `db:"id"         json:"id"`
	Post_ID   int64       `db:"post_id"    json:"postId"`
	Author_ID int64       `db:"author_id"  json:"authorId"`
	Parent_ID pgtype.Int8 `db:"parent_id"  json:"parentId"`
	// Parent_ID  sql.NullInt64 `db:"parent_id"  json:"parentID"`
	Message    string      `db:"message"    json:"message"`
	Vote       pgtype.Int8 `                json:"upVotes"`
	Created_at time.Time   `db:"created_at" json:"createdAt"`
}

type CommentResp struct {
	Comment
	Author `json:"author"`
}
