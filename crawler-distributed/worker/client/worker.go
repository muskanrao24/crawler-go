package client

import (
	"distributed-web-crawler/crawler-distributed/config"
	"distributed-web-crawler/crawler-distributed/worker"
	"distributed-web-crawler/crawler/engine"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(req engine.Request) (engine.ParseResult, error) {
		serializedRequest := worker.SerializeRequest(req)

		var serializedResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpcMethod,
			serializedRequest, &serializedResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(serializedResult), nil
	}

}
