package client

// Thread represents a thread on a specific board.
//
// It includes details such as the board name, and message of the thread.
//
// Board is the only mandatory field checked by the client but keep in mind that
// since each board has its own settings, more fields may be required (such as Body).
type Thread struct {
	Board    string   // Board where the thread is to be posted (e.g., 'b') (required)
	Name     string   // Name of the poster
	Email    string   // Email of the poster
	Subject  string   // Subject of the thread
	Spoiler  bool     // Marks attached files as spoiler
	Body     string   // The message of the thread
	Embed    string   // URL for an embedded media link (YouTube, Spotify...)
	Password string   // Password used to delete or edit the thread
	Sage     bool     // Replaces email with 'rabbia' and prevents bumping the thread
	Files    []string // Paths of the files to attach to the thread
}

// GetThread returns the ID of the thread to reply to (0 for threads).
func (t Thread) GetThread() int { return 0 }

// GetBoard returns the name of the board where the thread is to be posted.
func (t Thread) GetBoard() string { return t.Board }

// GetName returns the name of the poster.
func (t Thread) GetName() string { return t.Name }

// GetEmail returns the email of the poster.
func (t Thread) GetEmail() string { return t.Email }

// GetSubject returns the subject of the thread.
func (t Thread) GetSubject() string { return t.Subject }

// GetSpoiler returns whether the attached files are to be marked as spoiler.
func (t Thread) GetSpoiler() bool { return t.Spoiler }

// GetBody returns the message of the thread.
func (t Thread) GetBody() string { return t.Body }

// GetEmbed returns the embedded media URL associated with the thread.
func (t Thread) GetEmbed() string { return t.Embed }

// GetPassword returns the password for editing or deleting the thread.
func (t Thread) GetPassword() string { return t.Password }

// GetSage returns whether the thread will bump in the board.
func (t Thread) GetSage() bool { return t.Sage }

// GetFiles returns a list of paths of the files to be attached to the thread.
func (t Thread) GetFiles() []string { return t.Files }
