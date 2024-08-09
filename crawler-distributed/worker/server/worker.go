package main

import (
	"distributed-web-crawler/crawler-distributed/rpcSupport"
	"distributed-web-crawler/crawler-distributed/worker"
	"flag"
	"fmt"
)

var port = flag.Int("port", 0, "the port to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	err := rpcSupport.ServeRpc(fmt.Sprintf(":%d", *port),
		worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}
