package main

import (
	"coding-180/crawler_distributed/rpcsupport"
	"distributed-web-crawler/crawler-distributed/config"
	"distributed-web-crawler/crawler-distributed/persist"
	"fmt"
	"github.com/olivere/elastic"
)

func main() {
	// log.Fatal(serveRpc(host, index))
	err := serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort),
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
