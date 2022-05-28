package tag

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (Tag, error)
	GetAll(ctx context.Context) ([]Tag, error)
	Upsert(ctx context.Context, ar Tag) (Tag, error)
	DeleteByID(ctx context.Context, id int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) Repository {
	return &repository{db, redis}
}

func (m *repository) GetByID(ctx context.Context, id int) (res Tag, err error) {
	result := m.db.First(&res, id)

	return res, result.Error
}

func (m *repository) GetAll(ctx context.Context) (res []Tag, err error) {
	query := m.db

	result := query.Find(&res)

	return res, result.Error
}

func (m *repository) Upsert(ctx context.Context, tag Tag) (res Tag, err error) {
	result := m.db.Save(&tag)

	return tag, result.Error
}

func (m *repository) DeleteByID(ctx context.Context, id int) (err error) {
	result := m.db.Delete(&Tag{ID: id})

	return result.Error
}
