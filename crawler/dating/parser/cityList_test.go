package parser

import (
	"distributed-web-crawler/crawler/fetcher"
	"fmt"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", contents)

	result := ParseCityList(contents)

	fmt.Print(len(result.Requests))
}
