package main

import (
	"distributed-web-crawler/crawler-distributed/config"
	"distributed-web-crawler/crawler/dating/parser"
	"distributed-web-crawler/crawler/engine"
	"distributed-web-crawler/crawler/persist"
	"distributed-web-crawler/crawler/scheduler"
)

const cityUrl = "http://www.zhenai.com/zhenghun"

const index = "dating_profile"

func main() {
	itemChan, err := persist.ItemSaver(index)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
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
