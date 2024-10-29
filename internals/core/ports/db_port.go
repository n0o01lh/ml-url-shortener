package ports

import "github.com/n0o01lh/ml-url-shortener/internals/core/domain"

type DB interface {
	List()
	PutUrl(shortedUrl *domain.ShortedUrl) error
	UpdateUrl(key string, shortedUrl *domain.ShortedUrl) error
	GetUrl(key string) (*domain.ShortedUrl, error)
	PutStats(stats *domain.Stats) error
	UpdateStats(key string) error
	GetStats(key string) (*domain.Stats, error)
}
