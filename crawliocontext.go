package crawlio

import (
    "fmt"
    "sync"
)

type CrawlioContext struct{
        initialdomain string
        urlschannel chan string
        crawlers *sync.WaitGroup
        scheduler *sync.WaitGroup
        crawledurls []string
        lock *sync.RWMutex
}

// NewContext initializes a new Context instance
func NewCrawlioContext(initialUrl string) *CrawlioContext {
	return &CrawlioContext{
                initialdomain: initialUrl,
                urlschannel: make(chan string),
                crawlers: &sync.WaitGroup{},
                scheduler: &sync.WaitGroup{},
                crawledurls: make([]string, 1),
                lock: &sync.RWMutex{},
	}
}

// Method to add scrapped urls (synced)
func (context *CrawlioContext) AddScrapedUrl(url string) {
	context.lock.Lock()
        context.crawledurls = append(context.crawledurls, url)
        fmt.Println(url)
	context.lock.Unlock()
}

// Method for print stats on urls (synced)
func (context *CrawlioContext) PrintScrappedUrlsStats() {
        context.lock.RLock()
        //Free the lock after finish
	defer context.lock.RUnlock()
        fmt.Printf("len=%d cap=%d \n", len(context.crawledurls), cap(context.crawledurls))
}


