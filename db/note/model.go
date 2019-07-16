package note

import "time"

type Note struct {
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Created time.Time `json:"created,omitempty"`
}
