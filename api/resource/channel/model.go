package channel

import (
	"time"

	"fuchsrot/notification-hub/api/resource/message"

	"github.com/google/uuid"
)

type Form struct {
	Title       string `json:"title" form:"required,max=255"`
	Description string `json:"description"`
}

type DTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Channel struct {
	ID          uuid.UUID `gorm:"primarykey"`
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Messages    []message.Message `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Channels []*Channel

func (f *Form) ToModel() *Channel {
	return &Channel{
		Title:       f.Title,
		Description: f.Description,
	}
}

func (c *Channel) ToDto() *DTO {
	return &DTO{
		ID:          c.ID.String(),
		Title:       c.Title,
		Description: c.Description,
	}
}

func (cs Channels) ToDto() []*DTO {
	dtos := make([]*DTO, len(cs))
	for i, v := range cs {
		dtos[i] = v.ToDto()
	}

	return dtos
}
