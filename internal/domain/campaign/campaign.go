package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedOn time.Time `json:"created_at"`
	Content   string    `json:"content"`
	Contacts  []Contact `json:"contacts"`
}
