package main

import (
    "os"
    "github.com/adrianriobo/crawlio"
)

func main(){

  crawlio.Crawl(os.Args[1])
}
