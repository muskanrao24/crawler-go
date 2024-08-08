# distributed-web-crawler
course project for Introduction yo Golang on imooc.

## 单机版 / Single node version
/crawler

## 分布式版 / Distributed version
/crawler-distributed

## 前端页面 / Simple front end page
/frond-end

## 启动 / Run
### 单机版 / Single node version :
`docker run -d -p 9200:9200 elasticsearch:x.x.x (your es version)`

`(under project root directory) cd crawler`

`go run main.go`
### 页面 / Simple front end page :
`(under project root directory) cd front-end`

`go run start.go`

### 分布式版 / Distributed version :
`docker run -d -p 9200:9200 elasticsearch:x.x.x (your es version)`

`(under project root directory) cd crawler-distributed`

`go run main.go`

`(under crawler-distributed) cd persist`

`go run main.go`

`(under crawler-distributed) cd worker/server`

`go run main.go (start as many server as you want, as long as you add port configuration and set them in config.go)`