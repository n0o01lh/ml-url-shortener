package services

import (
	"time"

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
	now := time.Now().UTC()
	isoString := now.Format(time.RFC3339)

	shortedUrl := &domain.ShortedUrl{
		Id:          "http://ml.short/" + id,
		OriginalUrl: shortRequest.Url,
		Available:   shortRequest.Available,
		CreatedAt:   isoString,
		UpdatedAt:   isoString,
	}

	result, error := s.shortRepository.Create(shortedUrl)

	if error != nil {
		log.Error(error)
		return nil, error
	}

	return result, nil
}

func (s *ShortService) Update(id string, shortRequest *domain.ShortRequest) (*domain.ShortedUrl, error) {
	return nil, nil
}
