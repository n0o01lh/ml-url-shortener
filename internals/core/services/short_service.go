package services

import (
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type ShortService struct {
	shortRepository ports.ShortRepository
}

func NewShortService(shortRepository ports.ShortRepository) *ShortService {
	return &ShortService{shortRepository: shortRepository}
}

var _ ports.ShortService = (*ShortService)(nil)

func (s *ShortService) Create(shortRequest *domain.ShortRequest) (*domain.ShortedUrl, error) {
	return nil, nil
}

func (s *ShortService) Update(id string, shortRequest *domain.ShortRequest) (*domain.ShortedUrl, error) {
	return nil, nil
}
