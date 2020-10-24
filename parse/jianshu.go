package parse

import (
	"strings"
)

type Jianshu struct {
	BaseSite
}

func (j Jianshu) GetTitle(HTML string, xpath string) string {
	htmls := strings.Split(HTML, "\n")
	for _, v := range htmls {
		s := "<meta property=\"og:title\" content=\""
		if strings.Contains(v, s) {
			v = strings.Split(v, s)[1]
			v = strings.Split(v, "\"/>")[0]
			v = strings.Trim(v, " ")
			v = strings.Trim(v, "\n")
			return v
		}
	}
	return ""
}
