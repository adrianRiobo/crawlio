package crawlio

import (
    
)

func Crawl(url string) string {	
    //Create context
    crawlioctx := NewCrawlioContext(url)
    //Initialize handler
    crawlioctxhandler := DefaultCrawlioContextHandler{}
    crawlioctxhandler.Init(crawlioctx)
    //Crawl
    crawlioctxhandler.Crawl()
    return "done"
}
	
