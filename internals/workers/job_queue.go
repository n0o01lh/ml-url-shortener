package workers

import "sync"

type JobQueue struct {
	JobChannel    chan Job
	ResultChannel chan JobResult
	Wg            *sync.WaitGroup
}

func NewJobQueue(workers int) JobQueue {
	jobChan := make(chan Job)
	resultChan := make(chan JobResult)
	wg := &sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		worker := Worker{
			ID:            i + 1,
			JobQueue:      jobChan,
			ResultChannel: resultChan,
			Wg:            wg,
		}
		go worker.Start()
	}

	return JobQueue{JobChannel: jobChan, ResultChannel: resultChan, Wg: wg}
}
