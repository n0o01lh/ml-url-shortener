package workers

import "github.com/n0o01lh/ml-url-shortener/internals/core/domain"

type JobResult struct {
	ShortedUrl *domain.ShortedUrl
	Error      error
}
