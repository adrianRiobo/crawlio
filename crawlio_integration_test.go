// +build integration

package crawlio

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestIntegrationCrawl(t *testing.T) {
    result := Crawl("https://httpbin.org/html")
    assert.Equal(t, result, "done", "Should be done")
}
