package crawlio

import (
  "github.com/gocolly/colly"    
)

func Crawl(url string) string {	
   //Create context
   crawlioctx := NewCrawlioContext(url)
   //Initialize handler
   crawlioctxhandler := DefaultCrawlioContextHandler{}
   collector := colly.NewCollector()
   crawlioctxhandler.Init(crawlioctx, collector)
   //Crawl
   crawlioctxhandler.Crawl()
   return "done"
}
	
