package message

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(message *Message) (*Message, error) {
	if err := r.db.Create(message).Error; err != nil {
		return nil, err
	}

	return message, nil
}
