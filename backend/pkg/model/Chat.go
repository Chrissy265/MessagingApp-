package model

type Chat struct {
	CreatedTime string
	UpdateTime  string
	Messages    []Message
	Users       []User
}
