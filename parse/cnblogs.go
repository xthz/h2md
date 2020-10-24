package parse

import (
	"regexp"
	"strings"
)

type CNBlog struct {
	BaseSite
}

func (c CNBlog) ProcessDiv(div string) string {
	htmls := strings.Split(div, "\n")
	var html string
	for _, code := range htmls {
		// 将代码块内的行号取消掉
		if strings.Contains(code, "color: rgba(0, 128, 128, 1)") {
			var rexTemp = regexp.MustCompile("<span\\s*style=\"color:\\s*rgba\\(0,\\s*128,\\s*128,\\s*1\\)\">\\s*\\d+</span>")
			var results []string = rexTemp.FindAllString(code, -1)
			for _, value := range results {
				code = strings.Replace(code, value, "", -1)
			}
		}
		html += "\n" + code
	}
	return html
}
