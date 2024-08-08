package engine

import (
	"distributed-web-crawler/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	// log.Printf("Fetching url %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}
