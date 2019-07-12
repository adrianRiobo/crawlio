package crawlio

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCrawl(t *testing.T) {
    result := Crawl("http://test.test")
    assert.Equal(t, result, "done", "Should be done")
}
    
/*
func TestCrawl(t *testing.T) {
    want := "done"
    //os.Args = []string{"crawlio","https://test.test"}
    if got := Crawl("https://test.test"); got != want {
        t.Errorf("Crawl() = %q, want %q", got, want)
    }
}
*/
