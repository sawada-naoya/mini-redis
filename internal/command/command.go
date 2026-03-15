package command

// Command はクライアントから受け取ったコマンドを表す構造体
type Command struct {
	Name string
	Args []string
}