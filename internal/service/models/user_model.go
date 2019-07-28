package models

import "time"

type User struct {
	ID        int64      `json:"id"`        // 列名为 `id`
	UserName  string     `json:"user_name"` // 列名为 `user_name`
	Password  string     `json:"password"`  // 列名为 `password`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (u *User) TableName() string {
	return "user"
}
