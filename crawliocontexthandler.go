package crawlio

import (
  "errors"
  "sync"
  "strings"
  "github.com/bobesa/go-domain-util/domainutil"
  "github.com/gocolly/colly"
  "github.com/thoas/go-funk"
)

type CrawlioContextHandler interface {
  // Init the handler with a given conext
  Init(context *CrawlioContext) error
  // Crawl the context
  Crawl()
}

//Default impl
type DefaultCrawlioContextHandler struct {
  context *CrawlioContext
  urlschannel chan string
  crawlers *sync.WaitGroup
  scheduler *sync.WaitGroup
}

func (cch *DefaultCrawlioContextHandler) Init(context *CrawlioContext) error {
  if context == nil {
   return errors.New("No context available")
  } 
  cch.context = context
  cch.urlschannel= make(chan string)
  cch.crawlers= &sync.WaitGroup{}
  cch.scheduler= &sync.WaitGroup{}
  return nil
}

func (cch *DefaultCrawlioContextHandler) Crawl() {

    //Add one routine wait for initial crawler
    cch.crawlers.Add(1)
    //Add one routine wait for scheduler
    cch.scheduler.Add(1)

    go UrlCrawlingDecisor(cch.context, cch.urlschannel, cch.crawlers, cch.scheduler)
    go Crawler(cch.context, cch.context.initialdomain, cch.urlschannel, cch.crawlers)

    //when there is nothing else for crawl
    //close channel
    cch.crawlers.Wait()
    close(cch.urlschannel)

    //finally wait for scheduler
    cch.scheduler.Wait()
}

//Improve interface another search (by regex or whatever)
func Crawler(context *CrawlioContext, crawledurl string, urlschannel chan string, crawlers *sync.WaitGroup) {

  //Inform finish
  defer crawlers.Done()

  c := colly.NewCollector()

  c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    href := e.Attr("href")
    if domainutil.Domain(href) == "" && href != "/" {
      urlschannel <- (crawledurl + href)
    } else {
      urlschannel <- href
    }
  })

  c.Visit(crawledurl)

  //Create sitemap

}

//sync-async pattern governance of crawling
func UrlCrawlingDecisor(context *CrawlioContext, urlschannel chan string, crawlers *sync.WaitGroup, scheduler *sync.WaitGroup) {

    //Done when finish
    defer scheduler.Done()
    keepRunning := true

    for keepRunning {
      url, ok := <-urlschannel
      if ok {
        if IsCrawlable(context, url) {
           context.AddScrapedUrl(url)
           context.PrintScrappedUrlsStats()
           crawlers.Add(1)
           go Crawler(context, url, urlschannel, crawlers)
        }
      } else {
          keepRunning = false
      }
    }
}

//condition for crawlable url
func IsCrawlable(context *CrawlioContext, url string) bool {
  return domainutil.Domain(url) == domainutil.Domain(context.initialdomain) &&
           ! funk.Contains(context.crawledurls, url) &&
           ! domainutil.HasSubdomain(url) &&
           ! strings.ContainsRune(url, 35) &&
           ! strings.Contains(url, "..")
}
