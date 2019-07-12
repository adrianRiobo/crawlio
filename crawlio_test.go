package crawlio

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCrawl(t *testing.T) {
    result := Crawl("http://test.test")
    assert.Equal(t, result, "done", "Should be done")
}
    
