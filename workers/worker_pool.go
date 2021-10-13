package workers

import (
	"context"
	"sync"

	"github.com/lancepokaiwang/Golang_Web_Crawling/crawling"
	s "github.com/lancepokaiwang/Golang_Web_Crawling/errors"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
)

func worker(ctx context.Context, jobCh <-chan *job) {
	for {
		select {
		case job := <-jobCh:
			job.resultCh <- job.cc.PerformCrawling()
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
	jobCh        chan *job
}

func New(workersCount int, jobsCount int) *WorkerPool {
	return &WorkerPool{
		workersCount: workersCount,
		jobCh:        make(chan *job, jobsCount),
	}
}

func (wp *WorkerPool) Close() {
	close(wp.jobCh)
}

func (wp *WorkerPool) Run(ctx context.Context) {
	for i := 0; i < wp.workersCount; i++ {
		go worker(ctx, wp.jobCh)
	}
}

type job struct {
	cc       *crawling.CrawlClient
	wg       *sync.WaitGroup
	resultCh chan<- []productPB.ProductResponse
}

func newJob(cc *crawling.CrawlClient, wg *sync.WaitGroup, resultCh chan<- []productPB.ProductResponse) *job {
	return &job{
		cc:       cc,
		wg:       wg,
		resultCh: resultCh,
	}
}

func (wp *WorkerPool) QueueJob(ccs []*crawling.CrawlClient, wg *sync.WaitGroup, resultCh chan<- []productPB.ProductResponse) {
	for _, cc := range ccs {
		job := newJob(cc, wg, resultCh)
		// If jobCh is full, it will block here.
		wp.jobCh <- job
	}
}
