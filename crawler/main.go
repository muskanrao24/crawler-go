package main

import (
	"distributed-web-crawler/crawler/dating/parser"
	"distributed-web-crawler/crawler/engine"
)

const cityUrl = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		Url:        cityUrl,
		ParserFunc: parser.ParseCityList,
	})
}

// Parser 解析器
// input: utf-8编码文本
// output: Request{URL, 对应Parser}列表, Item列表
