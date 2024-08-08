package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const byteLength = 1024
const userAgentKey = "User-Agent"
const userAgentValue = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/71.0.3578.87 Safari/537.21"

var rateLimiter = time.Tick(10 * time.Millisecond)

// fetch content from url
func Fetch(url string) ([]byte, error) {

	// 限流
	<-rateLimiter

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(userAgentKey, userAgentValue)
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}
	resp, err := client.Do(req)

	if resp == nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// determine the encoding of input stream
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// log.Printf("检测输入流的编码")
	bytes, err := r.Peek(byteLength)

	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	// log.Printf("编码为: %s", e)
	return e
}
