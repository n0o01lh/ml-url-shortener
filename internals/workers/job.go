package workers

import "github.com/n0o01lh/ml-url-shortener/internals/core/ports"

type Job struct {
	key            string
	resolveService ports.ResolverService
	statsService   ports.StatsService
}

func NewJob(key string, resolveService ports.ResolverService, statsService ports.StatsService) Job {
	return Job{
		key:            key,
		resolveService: resolveService,
		statsService:   statsService,
	}
}
