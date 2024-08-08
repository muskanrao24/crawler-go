package model

import "distributed-web-crawler/crawler/engine"

type SearchResult struct {
	Hits     int
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []engine.Item
}
