package parse

import (
	"github.com/antchfx/htmlquery"
	"strings"
)

type BaseSite struct {
}

// GetTitle 获取HTML的文章标题
func (w BaseSite) GetTitle(HTML string, xpath string) string {
	root, _ := htmlquery.Parse(strings.NewReader(HTML))
	article := htmlquery.FindOne(root, xpath)
	title := htmlquery.InnerText(article)
	title = strings.Trim(title, " ")
	title = strings.Trim(title, "\n")
	return title
}

// GetBodyHTML 获取HTML里的文章正文
func (w BaseSite) GetBodyHTML(HTML string, xpath string) string {
	root, _ := htmlquery.Parse(strings.NewReader(HTML))
	article := htmlquery.FindOne(root, xpath)
	div := htmlquery.OutputHTML(article, true)
	div = w.ProcessDiv(div)
	return div
}

// ProcessDiv 二次处理文章正文的HTML内容, 生成易于转换md的规范化HTML
func (w BaseSite) ProcessDiv(div string) string {
	htmls := strings.Split(div, "\n")
	var html string
	for _, code := range htmls {
		if strings.Contains(code, "data-original-src") {
			code = strings.Replace(code, "data-original-src", "src", -1)
		}
		html += "\n" + code
	}
	return html
}
