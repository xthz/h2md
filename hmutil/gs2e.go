package hmutil

import (
	"github.com/wanghuijz/gsession"
	"log"
)

// Request web page
func Request(link string) string {
	session := gsession.Session()
	header := make(map[string]string)
	header["Connection"] = "close"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0"
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	header["Accept-Encoding"] = "gzip, deflate, br"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
	response, err := session.GET(link, header, true)
	if response == nil || err != nil {
		log.Fatal("response is nil")
	}
	html := response.Text()
	return html
}
