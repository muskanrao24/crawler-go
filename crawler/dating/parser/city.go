package parser

import (
	"distributed-web-crawler/crawler-distributed/config"
	"distributed-web-crawler/crawler/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

// parse city in city list
func ParseCity(contents []byte, _ string) engine.ParseResult {
	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(
		//	result.Items, "User "+string(m[2]))
		url := string(m[1])
		name := string(m[2])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    url,
				Parser: NewProfileParser(name),
			})
	}

	cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			})
	}

	return result
}
