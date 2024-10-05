package client

// Reply represents a response to a thread on a specific board.
//
// It includes details such as the thread ID, board name, and message of the reply.
//
// Thread and Board are the only mandatory fields checked by the client but keep in mind that
// since each board has its own settings, more fields may be required (such as Body).
type Reply struct {
	Thread   int      // ID of the thread to reply to (required)
	Board    string   // Board where the reply is to be posted (e.g., 'b') (required)
	Name     string   // Name of the poster
	Email    string   // Email of the poster
	Spoiler  bool     // Marks attached files as spoiler
	Body     string   // The message of the reply
	Embed    string   // URL for an embedded media link (YouTube, Spotify...)
	Password string   // Password used to delete or edit the reply
	Sage     bool     // Replaces email with 'rabbia' and prevents bumping the thread
	Files    []string // Paths of the files to attach to the reply
}

// GetThread returns the ID of the thread to reply to.
func (r Reply) GetThread() int { return r.Thread }

// GetBoard returns the name of the board where the reply is to be posted.
func (r Reply) GetBoard() string { return r.Board }

// GetName returns the name of the poster.
func (r Reply) GetName() string { return r.Name }

// GetEmail returns the email of the poster.
func (r Reply) GetEmail() string { return r.Email }

// GetSubject returns the subject of the reply (empty for replies).
func (r Reply) GetSubject() string { return "" }

// GetSpoiler returns whether the attached files are to be marked as spoiler.
func (r Reply) GetSpoiler() bool { return r.Spoiler }

// GetBody returns the message of the reply.
func (r Reply) GetBody() string { return r.Body }

// GetEmbed returns the embedded media URL associated with the reply.
func (r Reply) GetEmbed() string { return r.Embed }

// GetPassword returns the password for editing or deleting the reply.
func (r Reply) GetPassword() string { return r.Password }

// GetSage returns whether the reply will bump the thread.
func (r Reply) GetSage() bool { return r.Sage }

// GetFiles returns a list of paths of the files to be attached to the reply.
func (r Reply) GetFiles() []string { return r.Files }
