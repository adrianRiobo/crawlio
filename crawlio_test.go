package crawlio

import (
    "testing"
    //"os"
)

func TestCrawl(t *testing.T) {
    want := "done"
    //os.Args = []string{"crawlio","https://test.test"}
    if got := Crawl("https://test.test"); got != want {
        t.Errorf("Crawl() = %q, want %q", got, want)
    }
}
