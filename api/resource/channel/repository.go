package channel

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

func (r *Repository) Create(channel *Channel) (*Channel, error) {
	if err := r.db.Create(channel).Error; err != nil {
		return nil, err
	}

	return channel, nil
}

func (r *Repository) List() (Channels, error) {
	channels := make([]*Channel, 0)
	if err := r.db.Find(&channels).Error; err != nil {
		return nil, err
	}
	return channels, nil
}
