package main

import (
	"coding-180/crawler_distributed/rpcsupport"
	"distributed-web-crawler/crawler-distributed/config"
	"distributed-web-crawler/crawler-distributed/persist"
	"flag"
	"fmt"
	"github.com/olivere/elastic"
)

var port = flag.Int("port", 0, "the port to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	// log.Fatal(serveRpc(host, index))
	err := serveRpc(fmt.Sprintf(":%d", *port),
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
