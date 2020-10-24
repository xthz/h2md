package hmutil

import (
	"bytes"
	"fmt"
	. "github.com/mattn/godown"
	"log"
	"strings"
)

// HtmlToMarkdown HTML to Markdown, return markdown, string
func HtmlToMarkdown(title string, link string, body string, codeLanguage string) string {
	var buf bytes.Buffer
	err := Convert(&buf, strings.NewReader(body), &Option{
		GuessLang: func(s string) (string, error) { return codeLanguage, nil },
	})
	if err != nil {
		log.Fatal(err)
	}
	markdown := buf.String()
	markdown = ProcessMarkDown(title, link, markdown, codeLanguage)
	return markdown
}

// ProcessMarkDown secondary process markdown
func ProcessMarkDown(title string, link string, markdown string, codeLanguage string) string {
	ss := strings.Split(markdown, "\n")
	var result string
	for _, v := range ss {
		// 如果包含"__```", 删除"__"
		if strings.Contains(v, "__```") {
			v = strings.Replace(v, "__```", "\n```", 1)
		}
		// 博客园可能存在 **tips: **
		if strings.Contains(v, "**") {
			v = strings.Replace(v, "** ", "**", 1)
			v = strings.Replace(v, " **", "**", 1)
			v = strings.Replace(v, "** ", "**", 1)
			v = strings.Replace(v, " **", "**", 1)
		}
		// 将类似```go前加\n
		if strings.Contains(v, fmt.Sprintf("```%s", codeLanguage)) {
			v = strings.Replace(v, fmt.Sprintf("```%s", codeLanguage), fmt.Sprintf("\n```%s", codeLanguage), 1)
		}
		// 将 ![]( 前加\n
		if strings.Contains(v, "![](") {
			v = strings.Replace(v, "![](", "\n![](", 1)
		}
		// 应对简书的图片路径问题
		if strings.Contains(v, "upload-images") {
			v = strings.Replace(v, "//upload-images", "https://upload-images", 1)
		}
		// 如果一行全部都是"image.png", 删除这行
		if strings.TrimSpace(v) == "image.png" {
			continue
		}
		// 博客园的+-号删除 ExpandedBlockStart.gif
		if strings.Contains(v, "ExpandedBlockStart.gif") {
			if strings.Contains(v, fmt.Sprintf("```%s", codeLanguage)) {
				v = fmt.Sprintf("```%s", codeLanguage)
			}
		}
		result += "\n" + v
	}
	result = strings.Trim(result, " ")
	result = strings.Trim(result, "\n")
	return fmt.Sprintf("## %s\n\n> 原文地址: %s\n\n%s", title, link, result)
}
