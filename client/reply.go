package client

type Reply struct {
	Thread   int
	Board    string
	Name     string
	Email    string
	Spoiler  bool
	Body     string
	Embed    string
	Password string
	Sage     bool
	Files    []string
}

func (r Reply) GetThread() int      { return r.Thread }
func (r Reply) GetBoard() string    { return r.Board }
func (r Reply) GetName() string     { return r.Name }
func (r Reply) GetEmail() string    { return r.Email }
func (r Reply) GetSubject() string  { return "" }
func (r Reply) GetSpoiler() bool    { return r.Spoiler }
func (r Reply) GetBody() string     { return r.Body }
func (r Reply) GetEmbed() string    { return r.Embed }
func (r Reply) GetPassword() string { return r.Password }
func (r Reply) GetSage() bool       { return r.Sage }
func (r Reply) GetFiles() []string  { return r.Files }
