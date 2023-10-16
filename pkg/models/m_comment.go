package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Comment struct {
	ID         pgtype.Int8        `db:"id" json:"ID"`
	Post_ID    pgtype.Int8        `db:"post_id" json:"postID"`
	Author_ID  pgtype.Int8        `db:"author_id" json:"authorID"`
	Parent_ID  pgtype.Int8        `db:"parent_id" json:"parentID"`
	Message    pgtype.Text        `db:"message" json:"message"`
	Created_at pgtype.Timestamptz `db:"created_at" json:"created_at"`
}
