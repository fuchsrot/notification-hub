package message

import (
	"time"

	"github.com/google/uuid"
)

type Form struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Message struct {
	ID        uuid.UUID `gorm:"primarykey"`
	Title     string
	Content   string
	Type      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	ChannelID string
}

func (f *Form) ToModel() *Message {
	return &Message{
		Title:   f.Title,
		Content: f.Message,
		Type:    f.Type,
	}
}
