package workers

import (
	"context"
	"sync"

	"github.com/lancepokaiwang/Golang_Web_Crawling/crawling"
	s "github.com/lancepokaiwang/Golang_Web_Crawling/errors"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
)

func worker(ctx context.Context, jobs <-chan *job) {
	for {
		select {
		case job := <-jobs:
			job.resultChan <- job.cc.PerformCrawling(job.cc.Keyword, job.cc.Web)
			job.wg.Done()

		// Cancel worker.
		// It will be called when ctx cancel() is called.
		case <-ctx.Done():
			s.Println("Worker is canceld")
			return
		}
	}
}

type WorkerPool struct {
	workersCount int
	jobs         chan *job
}

func New(workersCount int) *WorkerPool {
	return &WorkerPool{
		workersCount: workersCount,
		// TODO: change hard coded "10" here. Just for test.
		jobs: make(chan *job, 10),
	}
}

func (wp *WorkerPool) Run(ctx context.Context) {
	for i := 0; i < wp.workersCount; i++ {
		go worker(ctx, wp.jobs)
	}
}

type job struct {
	cc         *crawling.CrawlClient
	wg         *sync.WaitGroup
	resultChan chan<- []productPB.ProductResponse
}

func newJob(cc *crawling.CrawlClient, wg *sync.WaitGroup, resultChan chan<- []productPB.ProductResponse) *job {
	return &job{
		cc:         cc,
		wg:         wg,
		resultChan: resultChan,
	}
}

func (wp *WorkerPool) NewJob(ccs []*crawling.CrawlClient, wg *sync.WaitGroup, resultChan chan<- []productPB.ProductResponse) {
	// If jobs is full, it will block here.
	for _, cc := range ccs {
		job := newJob(cc, wg, resultChan)
		wp.jobs <- job
	}
}
