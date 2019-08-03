package model

import "time"

type User struct {
	Username  string    `josn:"username" db:"username"`
	Password  string    `josn:"password" db:"password"`
	Secretkey string    `josn:"secret_key" db:"secret_key"`
	Oid       string    `json:"oid" db:"oid"`
	CreateAt  time.Time `json:"created_at" db:"created_at"`
	UpdateAt  time.Time `json:"update_at" db:"update_at"`
	Email     string    `json:"email" db:"email"`
	Role      string    `json:"role" db:"role"`
}
