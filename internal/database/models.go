// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID
	PostID    uuid.NullUUID
	UserID    uuid.NullUUID
	Comment   string
	Createdat sql.NullTime
}

type Follow struct {
	FollowerID  uuid.UUID
	FollowingID uuid.UUID
	CreatedAt   sql.NullTime
}

type Like struct {
	PostID uuid.UUID
	UserID uuid.UUID
}

type Post struct {
	ID        uuid.UUID
	Author    uuid.NullUUID
	Title     string
	Body      sql.NullString
	Media     []string
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Isprivate sql.NullBool
	Likes     sql.NullInt64
	Comments  sql.NullInt64
	Shares    sql.NullInt64
	Isdeleted sql.NullBool
	Flagged   sql.NullBool
	Flag      sql.NullString
}

type User struct {
	ID        uuid.UUID
	Firstname string
	Lastname  sql.NullString
	Username  string
	Email     string
	Password  string
	Region    string
	About     sql.NullString
	Pfpurl    sql.NullString
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Gender    sql.NullString
	Isprivate sql.NullBool
	Followers sql.NullInt32
	Following sql.NullInt32
	Posts     sql.NullInt32
	Site      sql.NullString
	Isdeleted sql.NullBool
	Dob       sql.NullTime
}
