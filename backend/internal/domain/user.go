package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string
	Birthday string
	NickName string
	Ctime    time.Time
}
