package news

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (News, error)
	GetBySpec(ctx context.Context, spec FilterSpec) ([]News, error)
	Upsert(ctx context.Context, ar News) (News, error)
	DeleteByID(ctx context.Context, id int) error
}

type repository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRepository(db *gorm.DB, redis *redis.Client) Repository {
	return &repository{db, redis}
}

func (m *repository) GetByID(ctx context.Context, id int) (res News, err error) {
	result := m.db.First(&res, id)

	return res, result.Error
}

func (m *repository) GetBySpec(ctx context.Context, spec FilterSpec) (res []News, err error) {
	query := m.db

	if spec.Topic != "" {
		query = query.Where("topic ILIKE ?", "%"+spec.Topic+"%")
	}

	if spec.Status != "" {
		query = query.Where("status = ?", spec.Status)
	}

	result := query.Find(&res)

	return res, result.Error
}

func (m *repository) Upsert(ctx context.Context, news News) (res News, err error) {
	result := m.db.Save(&news)

	return news, result.Error
}

func (m *repository) DeleteByID(ctx context.Context, id int) (err error) {
	result := m.db.Model(&News{}).Where("id = ?", id).Update("status", "deleted")

	return result.Error
}
