package main

import (
	"distributed-web-crawler/crawler-distributed/config"
	itemSaver "distributed-web-crawler/crawler-distributed/persist/client"
	"distributed-web-crawler/crawler-distributed/rpcSupport"
	worker "distributed-web-crawler/crawler-distributed/worker/client"
	"distributed-web-crawler/crawler/dating/parser"
	"distributed-web-crawler/crawler/engine"
	"distributed-web-crawler/crawler/scheduler"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

const cityUrl = "http://www.zhenai.com/zhenghun"

var (
	itemSaverHost = flag.String(
		"itemSaver_host", "", "itemSaver host")
	workerHosts = flag.String(
		"worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    cityUrl,
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcSupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Panicf("Connected to %s", h)
		} else {
			log.Panicf(
				"error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}

// Parser 解析器
// input: utf-8编码文本
// output: Request{URL, 对应Parser}列表, Item列表
