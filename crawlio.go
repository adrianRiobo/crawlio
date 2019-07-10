package crawlio

import (
    //"os"
    //"sync"
    "strings"
    "github.com/bobesa/go-domain-util/domainutil"
    "github.com/gocolly/colly"
    "github.com/thoas/go-funk"
)

func Crawl(url string) string {	
    
    //improve test 1 arg
    //from main os.Args[1]
    crawlioctx := NewCrawlioContext(url)

    //Add one routine wait for initial crawler 
    crawlioctx.crawlers.Add(1)
    //Add one routine wait for scheduler
    crawlioctx.scheduler.Add(1)
    
    go UrlCrawlingDecisor(crawlioctx)
    go Crawler(crawlioctx, crawlioctx.initialdomain)
   
    //when there is nothing else for crawl
    //close channel 
    crawlioctx.crawlers.Wait()
    close(crawlioctx.urlschannel)

    //finally wait for scheduler
    crawlioctx.scheduler.Wait()

    return "done"
}
	
//Improve interface another search (by regex or whatever)
func Crawler(context *CrawlioContext, crawledurl string) {

  //Inform finish
  defer context.crawlers.Done()
  
  c := colly.NewCollector()

  c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    href := e.Attr("href")
    if domainutil.Domain(href) == "" && href != "/" {
      context.urlschannel <- (crawledurl + href)
    } else {
      context.urlschannel <- href
    }
  })

  c.Visit(crawledurl)

  //Create sitemap

}

//sync-async pattern governance of crawling
func UrlCrawlingDecisor(context *CrawlioContext) {

    //Done when finish
    defer context.scheduler.Done()
    keepRunning := true

    for keepRunning {
      url, ok := <-context.urlschannel
      if ok {
        //fmt.Println(url)
        if domainutil.Domain(url) == domainutil.Domain(context.initialdomain) &&
           ! funk.Contains(context.crawledurls, url) && 
           ! domainutil.HasSubdomain(url) &&
           ! strings.ContainsRune(url, 35) &&
           ! strings.Contains(url, "..") {
           context.AddScrapedUrl(url)
           context.PrintScrappedUrlsStats()
           context.crawlers.Add(1)
           go Crawler(context, url)
        } 
      } else {
          keepRunning = false
      }
    }
}


