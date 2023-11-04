package models

import "database/sql"

type TokenEntity struct {
	Id              string
	Token           string
	User_id         string
	Expiration_time string
	Created_at      string
	Last_used_at    sql.NullString
	Canceled        bool
	Canceled_at     sql.NullString
}
