package workers

import (
	"context"
	"sync"

	"github.com/lancepokaiwang/Golang_Web_Crawling/crawling"
	productPB "github.com/lancepokaiwang/Golang_Web_Crawling/proto/product"
)

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan crawling.CrawlerInterface, results chan<- productPB.ProductResponse) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			results <- job.Crawl("test", 1)
		case <-ctx.Done():
			// fmt.Printf("cancelled worker. Error detail: %v\n", ctx.Err())
			// results <- productPB.ProductResponse{
			// 	Err: ctx.Err(),
			// }
			return
		}
	}
}

type WorkerPool struct {
	workersCount int
	jobs         chan crawling.CrawlerInterface
	results      chan productPB.ProductResponse
	Done         chan struct{}
}

func New(wcount int) WorkerPool {
	return WorkerPool{
		workersCount: wcount,
		jobs:         make(chan crawling.CrawlerInterface, wcount),
		results:      make(chan productPB.ProductResponse, wcount),
		Done:         make(chan struct{}),
	}
}

func (wp WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		go worker(ctx, &wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.Done)
	close(wp.results)
}

func (wp WorkerPool) Results() <-chan productPB.ProductResponse {
	return wp.results
}

func (wp WorkerPool) NewJob(jobsBulk []crawling.CrawlerInterface) {
	for _, job := range jobsBulk {
		wp.jobs <- job
	}
	close(wp.jobs)
}
