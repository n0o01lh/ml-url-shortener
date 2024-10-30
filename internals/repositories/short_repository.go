package repositories

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
	"github.com/n0o01lh/ml-url-shortener/internals/data"
)

type ShortRepository struct {
	database *data.DynamoDb
}

func NewShortRepository(database *data.DynamoDb) *ShortRepository {
	return &ShortRepository{
		database: database,
	}
}

var _ ports.ShortRepository = (*ShortRepository)(nil)

func (r *ShortRepository) Create(shortedUrl *domain.ShortedUrl) error {
	err := r.database.PutUrl(shortedUrl)

	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *ShortRepository) Update(id string, shortedUrl *domain.ShortedUrl) error {
	err := r.database.UpdateUrl(id, shortedUrl)

	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
