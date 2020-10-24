package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"h2md/hmutil"
	"h2md/parse"
	"log"
)

func main() {
	var inTE, outTE *walk.TextEdit
	var window *walk.MainWindow

	err := MainWindow{
		Title:    "H2MD",
		AssignTo: &window,
		Size: Size{
			Width:  210,
			Height: 200,
		},
		Layout: VBox{},
		Children: []Widget{
			Label{
				Text: "Language:",
			},
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "jianshu.com",
				OnClicked: func() {
					myStruct := parse.Jianshu{}
					_ = outTE.SetText("")
					lang := inTE.Text()
					link, _ := hmutil.UtilPaste()
					html := hmutil.Request(link)
					title := myStruct.GetTitle(html, "")
					body := myStruct.GetBodyHTML(html, "//*[@class=\"_2rhmJa\"]")
					markdown := hmutil.HtmlToMarkdown(title, link, body, lang)
					_ = hmutil.UtilCopy(markdown)
					_ = outTE.SetText("done")
				},
			},
			PushButton{
				Text: "csdn.net",
				OnClicked: func() {
					myStruct := parse.CSDN{}
					_ = outTE.SetText("")
					lang := inTE.Text()
					link, _ := hmutil.UtilPaste()
					html := hmutil.Request(link)
					title := myStruct.GetTitle(html, "")
					body := myStruct.GetBodyHTML(html, "//*[@id=\"content_views\"]")
					markdown := hmutil.HtmlToMarkdown(title, link, body, lang)
					_ = hmutil.UtilCopy(markdown)
					_ = outTE.SetText("done")
				},
			},
			PushButton{
				Text: "cnblogs.com",
				OnClicked: func() {
					myStruct := parse.CNBlog{}
					_ = outTE.SetText("")
					lang := inTE.Text()
					link, _ := hmutil.UtilPaste()
					html := hmutil.Request(link)
					title := myStruct.GetTitle(html, "")
					body := myStruct.GetBodyHTML(html, "//*[@id=\"cnblogs_post_body\"]")
					markdown := hmutil.HtmlToMarkdown(title, link, body, lang)
					_ = hmutil.UtilCopy(markdown)
					_ = outTE.SetText("done")
				},
			},
			PushButton{
				Text: "weixin.qq.com",
				OnClicked: func() {
					myStruct := parse.Weixin{}
					_ = outTE.SetText("")
					lang := inTE.Text()
					link, _ := hmutil.UtilPaste()
					html := hmutil.Request(link)
					title := myStruct.GetTitle(html, "")
					body := myStruct.GetBodyHTML(html, "//*[@id=\"js_content\"]")
					markdown := hmutil.HtmlToMarkdown(title, link, body, lang)
					_ = hmutil.UtilCopy(markdown)
					_ = outTE.SetText("done")
				},
			},
		},
	}.Create()
	if err != nil {
		log.Println(err)
	}
	win.SetWindowLong(
		window.Handle(), win.GWL_STYLE,
		win.GetWindowLong(window.Handle(), win.GWL_STYLE) & ^win.WS_MAXIMIZEBOX & ^win.WS_THICKFRAME,
	)
	window.Run()
}
