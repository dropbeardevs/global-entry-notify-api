package users

import "time"

type User struct {
	UserId      int64
	MessageId   string
	LocationIds []int
	BeforeDate  time.Time
}
