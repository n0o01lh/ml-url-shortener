package repositories

import (
	"fmt"

	"github.com/n0o01lh/ml-url-shortener/internals/core/domain"
	"github.com/n0o01lh/ml-url-shortener/internals/core/ports"
)

type ShortRepository struct {
}

func NewShortRepository() *ShortRepository {
	return &ShortRepository{}
}

var _ ports.ShortRepository = (*ShortRepository)(nil)

func (r *ShortRepository) Create(shortedUrl *domain.ShortedUrl) (*domain.ShortedUrl, error) {
	fmt.Println(shortedUrl)
	return nil, nil
}

func (r *ShortRepository) Update(id string, shortedUrl *domain.ShortedUrl) (*domain.ShortedUrl, error) {
	return nil, nil

}
