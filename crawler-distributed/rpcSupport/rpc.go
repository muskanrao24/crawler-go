package rpcSupport

import (
	"distributed-web-crawler/crawler-distributed/config"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {

	err := rpc.Register(service)
	if err != nil {
		return err
	}

	listener, err := net.Listen(config.RpcProtocol, host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}

	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial(config.RpcProtocol, host)

	if err != nil {
		return nil, err
	}

	client := jsonrpc.NewClient(conn)
	return client, nil
}
