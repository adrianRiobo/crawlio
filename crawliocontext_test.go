package crawlio

import (
    "testing"
)

//test initialization
func TestNewCrawlioContext(t *testing.T) {
	crawlioctx := NewCrawlioContext("https://test.test")
        if crawlioctx.initialdomain != "https://test.test" {
           t.Fatal("value not equal")
        }
}
