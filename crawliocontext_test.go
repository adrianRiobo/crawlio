package crawlio

import (
  "testing"
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

// Test suit runner
func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(CrawlioContextTestSuite))
}

