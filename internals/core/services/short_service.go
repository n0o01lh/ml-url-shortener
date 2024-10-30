package services

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
	"github.com/n0o01lh/ml-url-shortener/internals/utils"
)

type ShortService struct {
	shortRepository ports.ShortRepository
}

func NewShortService(shortRepository ports.ShortRepository) *ShortService {
	return &ShortService{shortRepository: shortRepository}
}

var _ ports.ShortService = (*ShortService)(nil)

func (s *ShortService) Create(shortRequest *domain.ShortRequest) (*domain.ShortedUrl, error) {

	id := utils.GetRandomString()
	currentDate := utils.GetCurrentDate()

	shortedUrl := &domain.ShortedUrl{
		Id:          id,
		OriginalUrl: shortRequest.Url,
		Available:   shortRequest.Available,
		CreatedAt:   currentDate,
		UpdatedAt:   currentDate,
	}

	err := s.shortRepository.Create(shortedUrl)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return shortedUrl, nil
}

func (s *ShortService) Update(id string, shortRequest *domain.ShortRequest) (*domain.ShortedUrl, error) {

	shortedUrl := &domain.ShortedUrl{
		Id:          id,
		OriginalUrl: shortRequest.Url,
		Available:   shortRequest.Available,
		UpdatedAt:   utils.GetCurrentDate(),
	}

	err := s.shortRepository.Update(id, shortedUrl)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return shortedUrl, nil
}
