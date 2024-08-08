package client

import (
	"distributed-web-crawler/crawler-distributed/config"
	"distributed-web-crawler/crawler-distributed/rpcSupport"
	"distributed-web-crawler/crawler/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {

	client, err := rpcSupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d: %v",
				itemCount, item)
			itemCount++

			// Call rpc to save item
			result := ""
			err := client.Call(config.ItemSaverRpcMethod,
				item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
