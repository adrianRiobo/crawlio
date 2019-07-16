package main

import (
    "os"
    "github.com/adrianRiobo/crawlio"
)

func main(){

  crawlio.Crawl(os.Args[1])
}
