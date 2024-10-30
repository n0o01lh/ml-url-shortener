package services

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type StatsService struct {
	statsRepository ports.StatsRepository
}

func NewStatsService(statsRepository ports.StatsRepository) *StatsService {
	return &StatsService{
		statsRepository: statsRepository,
	}
}

var _ ports.StatsService = (*StatsService)(nil)

func (s *StatsService) Create(id string) error {

	stats := domain.NewStats(id, 0)

	err := s.statsRepository.Create(stats)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *StatsService) Update(id string) error {

	err := s.statsRepository.Update(id)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *StatsService) Get(id string) (*domain.Stats, error) {

	result, err := s.statsRepository.Get(id)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return result, nil
}
