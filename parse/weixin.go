package parse

import "strings"

type Weixin struct {
	BaseSite
}

func (wx Weixin) GetTitle(HTML string, xpath string) string {
	htmls := strings.Split(HTML, "\n")
	for _, v := range htmls {
		if strings.Contains(v, "og:title") {
			v = strings.Split(v, "content=\"")[1]
			v = strings.Split(v, "\" />")[0]
			v = strings.Trim(v, " ")
			v = strings.Trim(v, "\n")
			return v
		}
	}
	return ""
}

func (wx Weixin) ProcessDiv(div string) string {
	divs := strings.Split(div, "<pre")
	var newDIV string
	for _, v := range divs {
		v = "<pre" + v
		v = strings.Replace(v, "</span>", "\n</span>", -1)
		newDIV += v
	}
	newDIV = strings.Replace(newDIV, "data-src", "src", -1)
	return newDIV
}
