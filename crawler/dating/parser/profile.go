package parser

import (
	"distributed-web-crawler/crawler-distributed/config"
	"distributed-web-crawler/crawler/engine"
	"distributed-web-crawler/crawler/models"
	"github.com/valyala/fastjson"
	"log"
	"regexp"
	"strconv"
)

var jsonRe = regexp.MustCompile(`<script>window.__INITIAL_STATE__=({.+});`)

var expandRe = regexp.MustCompile(
	`<a target="_blank" href="(http://www.zhenai.com/zhenghun/[^"]+)">`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

// parse user profile
func parseProfile(content []byte, url string, name string) engine.ParseResult {

	matches := jsonRe.FindAllSubmatch(content, -1)

	// log.Println("matches: ", len(matches))

	// new profile object
	profile := models.Profile{}

	for _, m := range matches {

		objectString := string(m[1])
		var p fastjson.Parser
		v, err := p.Parse(objectString)
		if err != nil {
			log.Fatal(err)
		}

		profile.Name = name

		// 用户信息对象
		userInfoObject := v.GetObject("objectInfo")
		if userInfoObject != nil {
			// name
			//if userInfoObject.Get("nickname") != nil {
			//	s := userInfoObject.Get("nickname").String()
			//	profile.Name = s[1 : len(s) - 1]
			//}
			// nickname
			//if userInfoObject.Get("nickname") != nil {
			//	profile.Nickname = userInfoObject.Get("nickname").String()
			//}
			// gender
			if userInfoObject.Get("genderString") != nil {
				s := userInfoObject.Get("genderString").String()
				profile.Gender = s[1 : len(s)-1]
			}
			// age
			if userInfoObject.Get("age") != nil {
				profile.Age = userInfoObject.Get("age").GetInt()
			}
			// income
			if userInfoObject.Get("salaryString") != nil {
				s := userInfoObject.Get("salaryString").String()
				profile.Income = s[1 : len(s)-1]
			}
			// marriage
			if userInfoObject.Get("marriageString") != nil {
				s := userInfoObject.Get("marriageString").String()
				profile.Marriage = s[1 : len(s)-1]
			}
			// education
			if userInfoObject.Get("educationString") != nil {
				s := userInfoObject.Get("educationString").String()
				profile.Education = s[1 : len(s)-1]
			}
			// city
			//if userInfoObject.Get("workCityString") != nil {
			//	profile.City = userInfoObject.Get("workCityString").String()
			//}
			// height
			if userInfoObject.Get("heightString") != nil {
				heightString := userInfoObject.Get("heightString").String()
				heightWithoutUnit := heightString[1:4]
				height, _ := strconv.Atoi(heightWithoutUnit)
				profile.Height = height
			}

			// 基础信息列表
			// basicInfo, _ := userInfoObject.Get("basicInfo").Array()
			// 详细信息列表
			//detailInfo, _ := userInfoObject.Get("detailInfo").Array()
			//
			//if basicInfo != nil && len(basicInfo) >= 8 {
			//	// weight
			//	weightString := basicInfo[4].String()
			//	weightWithoutUnit := weightString[1:3]
			//	weight, _ := strconv.Atoi(weightWithoutUnit)
			//	profile.Weight = weight
			//	// occupation
			//	profile.Occupation = basicInfo[len(basicInfo)-2].String()
			//	// constellation
			//	profile.Constellation = basicInfo[2].String()
			//}

			//if detailInfo != nil && len(detailInfo) >= 8 {
			//	// nationality
			//	profile.Nationality = detailInfo[0].String()
			//	// smoke
			//	profile.Smoke = detailInfo[3].String()
			//	// drink
			//	profile.Drink = detailInfo[4].String()
			//	// house
			//	profile.House = detailInfo[5].String()
			//	// car
			//	profile.Car = detailInfo[6].String()
			//}
		}
	}

	id := extractString([]byte(url), idUrlRe)

	result := engine.ParseResult{}
	result.Items = append(result.Items,
		engine.Item{
			Url:     url,
			Type:    "zhenai",
			Id:      id,
			Payload: profile,
		})

	expandMatches := expandRe.FindAllSubmatch(content, -1)
	for _, m := range expandMatches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			})
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
