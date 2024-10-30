package services

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type ResolverService struct {
	resolverRepository ports.ResolverRepository
}

func NewResolverService(resolverRepository ports.ResolverRepository) *ResolverService {
	return &ResolverService{
		resolverRepository: resolverRepository,
	}
}

var _ ports.ResolverService = (*ResolverService)(nil)

func (s *ResolverService) Get(id string) (*domain.ShortedUrl, error) {

	shortedUrl, err := s.resolverRepository.Get(id)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return shortedUrl, nil
}
