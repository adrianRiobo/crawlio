// +build integration

// from https://peter.bourgon.org/go-in-production/#testing-and-validation
// Use tags + flags
package crawlio

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIntegrationCrawl(t *testing.T) {
    //Improve getting url and number of scrapped urls from tags
    result := Crawl("https://httpbin.org/links/4/0")
    //Improve getting context scrapped urls and check len is 4
    assert.Equal(t, result, "done", "Should be done")
}
