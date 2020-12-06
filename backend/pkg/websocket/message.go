package websocket

type Message struct {
	UserID  string "json:\"userid\""
	ChatID  string "json:\"chatid\""
	Content string "json:\"content\""
}
