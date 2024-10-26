package model

type Post interface {
	GetThread() int
	GetBoard() string
	GetName() string
	GetEmail() string
	GetSubject() string
	GetSpoiler() bool
	GetBody() string
	GetEmbed() string
	GetPassword() string
	GetSage() bool
	GetFiles() []string
}
