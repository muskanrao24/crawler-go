package main

import (
	"distributed-web-crawler/crawler-distributed/config"
	itemSaver "distributed-web-crawler/crawler-distributed/persist/client"
	worker "distributed-web-crawler/crawler-distributed/worker/client"
	"distributed-web-crawler/crawler/dating/parser"
	"distributed-web-crawler/crawler/engine"
	"distributed-web-crawler/crawler/scheduler"
	"fmt"
)

const cityUrl = "http://www.zhenai.com/zhenghun"

func main() {
	itemChan, err := itemSaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

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

// Parser 解析器
// input: utf-8编码文本
// output: Request{URL, 对应Parser}列表, Item列表
