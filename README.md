# Web Crawler written in Go!

An experimental crawler written in Golang, with Elasticsearch as the storage to learn Golang, coroutines, distributed system, and Elasticsearch.

## Tech stack

1. Golang 1.11
2. Elasticsearch 6.5.4
3. Docker

## Folder structure

`/crawler` : Single node version

`/crawler-distributed` : Distributed version

`/front-end` : Simple front end page

## Run

### Single instance

```shell
docker run -d -p 9200:9200 elasticsearch:x.x.x # your es version
cd crawler
go run main.go
```

## Multiple instances

```shell
docker run -d -p 9200:9200 elasticsearch:x.x.x # your es version
cd persist
go run itemSaver.go
cd worker/server
go run worker.go # start as many server as you want, as long as you add port configuration and set them in config.go
cd crawler-distributed
go run main.go
```

### Basic front end

```shell
cd front-end
go run start.go
```

## Todos

[ ] Crawl using css selector or xpath (instead of regular expression)
[ ] Follow robots agreement
[ ] A more useful front-end page
[ ] Use Docker + Kubernetes to package and deploy
