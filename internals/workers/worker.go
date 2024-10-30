package workers

import (
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

type Worker struct {
	ID            int
	JobQueue      <-chan Job
	ResultChannel chan<- JobResult
	Wg            *sync.WaitGroup
}

func (w *Worker) Start() {
	for job := range w.JobQueue {
		result := w.processJob(job)
		w.ResultChannel <- result
	}
}

func (w *Worker) processJob(job Job) JobResult {
	defer w.Wg.Done()
	shortedUrl, err := job.resolveService.Get(job.key)

	if err == nil {
		err := job.statsService.Update(job.key)
		if err != nil {
			log.Error("error")
		}
	}

	return JobResult{
		ShortedUrl: shortedUrl,
		Error:      err,
	}
}
