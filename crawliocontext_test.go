package crawlio

import (
  "sync"
  "testing"
  "strconv"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"   
)

type CrawlioContextTestSuite struct {
    suite.Suite
    context *CrawlioContext
}

//Setup content before each test
func (suite *CrawlioContextTestSuite) SetupTest() {
    suite.context = NewCrawlioContext("https://test.test")
}

// Tests
func (suite *CrawlioContextTestSuite) TestCrawlioContextTest() {
    assert.Equal(suite.T(), "https://test.test", suite.context.initialdomain)
}

// Tests
func (suite *CrawlioContextTestSuite) TestConcurrentAddScrapedUrl() {
    workers:= &sync.WaitGroup{}    
    for i := 0; i < 10; i++ {
      workers.Add(1)
      go func(context *CrawlioContext, workers *sync.WaitGroup, url string) {
          defer workers.Done()          
          context.AddScrapedUrl(url)
      }(suite.context, workers, "url" + strconv.Itoa(i))
    }
    workers.Wait()    
    //crawledurls should be 10 from loop + initialdomain
    assert.Equal(suite.T(), 11, len(suite.context.crawledurls))
}


// Test suit runner
func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(CrawlioContextTestSuite))
}

