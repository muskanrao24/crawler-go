package main

import (
	"distributed-web-crawler/crawler-distributed/config"
	"distributed-web-crawler/crawler-distributed/rpcSupport"
	"distributed-web-crawler/crawler-distributed/worker"
	"fmt"
)

func main() {
	err := rpcSupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}
