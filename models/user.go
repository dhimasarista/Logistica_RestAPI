package models

import "database/sql"

type User struct {
	ID          sql.NullInt64  `json:"id"`
	Username    sql.NullString `json:"username"`
	Password    sql.NullString `json:"password"`
	IsUser      sql.NullInt64  `json:"is_user"`
	IsSuperuser sql.NullString `json:"is_superuser"`
}
