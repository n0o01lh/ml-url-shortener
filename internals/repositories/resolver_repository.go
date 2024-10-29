package repositories

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
	"github.com/n0o01lh/ml-url-shortener/internals/data"
)

type ResolverRepository struct {
	database *data.DynamoDb
}

func NewResolverRepository(database *data.DynamoDb) *ResolverRepository {
	return &ResolverRepository{
		database: database,
	}
}

var _ ports.ResolverRepository = (*ResolverRepository)(nil)

func (r *ResolverRepository) Get(id string) (*domain.ShortedUrl, error) {

	result, err := r.database.GetUrl(id)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return result, nil
}
