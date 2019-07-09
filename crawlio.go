package main

import (
    "fmt"
    "os"
    "sync"
    "strings"
    "github.com/bobesa/go-domain-util/domainutil"
    "github.com/gocolly/colly"
    "github.com/thoas/go-funk"
)

type CrawlingContext struct{
	initialdomain string
	urlschannel chan string
        crawlers sync.WaitGroup
        scheduler sync.WaitGroup
        crawledurls []string
}

func main() {	
    
    //improve test 1 arg

    crawlingcontext := CrawlingContext{
      initialdomain: os.Args[1],
      urlschannel: make(chan string), 
    }
    
    //Add one routine wait for initial crawler 
    crawlingcontext.crawlers.Add(1)
    //Add one routine wait for scheduler
    crawlingcontext.scheduler.Add(1)
    
    go UrlCrawlingDecisor(&crawlingcontext)
    go Crawler(&crawlingcontext, crawlingcontext.initialdomain)
   
    //when there is nothing else for crawl
    //close channel 
    crawlingcontext.crawlers.Wait()
    close(crawlingcontext.urlschannel)

    //finally wait for scheduler
    crawlingcontext.scheduler.Wait()
}
	
//Improve interface another search (by regex or whatever)
func Crawler(context *CrawlingContext, crawledurl string) {

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
func UrlCrawlingDecisor(context *CrawlingContext) {

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
           context.crawledurls = append(context.crawledurls, url)
           printSlice(context.crawledurls)
           fmt.Println(url)
           context.crawlers.Add(1)
           go Crawler(context, url)
        } 
      } else {
          fmt.Println("done")
          keepRunning = false
      }
    }
}

func printSlice(s []string) {
  //fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
  fmt.Printf("len=%d cap=%d \n", len(s), cap(s)) 
}

