package parse

import "strings"

type CSDN struct {
	BaseSite
}

func (c CSDN) GetTitle(HTML string, xpath string) string {
	htmls := strings.Split(HTML, "\n")
	for _, v := range htmls {
		if strings.Contains(v, "var articleTitle = \"") {
			v = strings.Replace(v, "var articleTitle = \"", "", -1)
			v = v[:len(v)-2]
			v = strings.Trim(v, " ")
			v = strings.Trim(v, "\n")
			return v
		}
	}
	return ""
}
