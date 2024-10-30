package repositories

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
	"github.com/n0o01lh/ml-url-shortener/internals/data"
)

type StatsRepository struct {
	database *data.DynamoDb
}

func NewStatsRepository(database *data.DynamoDb) *StatsRepository {
	return &StatsRepository{
		database: database,
	}
}

var _ ports.StatsRepository = (*StatsRepository)(nil)

func (r *StatsRepository) Create(stats *domain.Stats) error {

	if err := r.database.PutStats(stats); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *StatsRepository) Update(id string) error {

	if err := r.database.UpdateStats(id); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *StatsRepository) Get(id string) (*domain.Stats, error) {

	stats, err := r.database.GetStats(id)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return stats, nil
}
