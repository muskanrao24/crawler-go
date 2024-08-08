package client

import (
	"distributed-web-crawler/crawler-distributed/config"
	"distributed-web-crawler/crawler-distributed/rpcSupport"
	"distributed-web-crawler/crawler-distributed/worker"
	"distributed-web-crawler/crawler/engine"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcSupport.NewClient(
		fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err

	}
	return func(req engine.Request) (engine.ParseResult, error) {
		serializedRequest := worker.SerializeRequest(req)

		var serializedResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpcMethod,
			serializedRequest, &serializedResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(serializedResult), nil
	}, nil

}
