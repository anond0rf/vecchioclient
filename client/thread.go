package client

type Thread struct {
	Board    string
	Name     string
	Email    string
	Subject  string
	Spoiler  bool
	Body     string
	Embed    string
	Password string
	Sage     bool
	Files    []string
}

func (t Thread) GetThread() int      { return 0 }
func (t Thread) GetBoard() string    { return t.Board }
func (t Thread) GetName() string     { return t.Name }
func (t Thread) GetEmail() string    { return t.Email }
func (t Thread) GetSubject() string  { return t.Subject }
func (t Thread) GetSpoiler() bool    { return t.Spoiler }
func (t Thread) GetBody() string     { return t.Body }
func (t Thread) GetEmbed() string    { return t.Embed }
func (t Thread) GetPassword() string { return t.Password }
func (t Thread) GetSage() bool       { return t.Sage }
func (t Thread) GetFiles() []string  { return t.Files }
