package model

type Chat struct {
	ID          int64
	CreatedTime string
	UpdateTime  string
	Messages    []Message
	Users       []User
}
