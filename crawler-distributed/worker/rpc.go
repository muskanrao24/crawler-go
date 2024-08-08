package worker

import "distributed-web-crawler/crawler/engine"

type CrawlService struct{}

func (CrawlService) Process(
	req Request, result *ParseResult) error {
	engineRequest, err := DeserializeRequest(req)
	if err != nil {
		return err

	}

	engineResult, err := engine.Worker(engineRequest)
	if err != nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
