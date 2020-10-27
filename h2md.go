package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"h2md/hmutil"
	"h2md/parse"
	"log"
)

func notify() {
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}

	icon, err := walk.Resources.Icon("img/stop.ico")
	if err != nil {
		log.Fatal(err)
	}

	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = ni.Dispose()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}

	if err := ni.SetToolTip("click menu to exit."); err != nil {
		log.Fatal(err)
	}

	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}

		if err := ni.ShowCustom(
			"Walk NotifyIcon Example",
			"There are multiple ShowX methods sporting different icons.",
			icon); err != nil {

			log.Fatal(err)
		}
	})

	// We put an exit action into the context menu.
	exitAction := walk.NewAction()
	if err := exitAction.SetText("E&xit"); err != nil {
		log.Fatal(err)
	}
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	// The notify icon is hidden initially, so we have to make it visible.
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	// Now that the icon is visible, we can bring up an info balloon.
	if err := ni.ShowInfo("Walk NotifyIcon Example", "Click the icon to show again."); err != nil {
		log.Fatal(err)
	}

	// Run the message loop.
	mw.Run()
}

func main() {
	var inTE, outTE *walk.TextEdit
	var window *walk.MainWindow

	err := MainWindow{
		Title:    "H2MD",
		AssignTo: &window,
		Size: Size{
			Width:  230,
			Height: 200,
		},
		//Icon: appIcon,
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

	icon, _ := walk.NewIconFromFile("favicon.ico")
	_ = window.SetIcon(icon)

	win.SetWindowLong(
		window.Handle(), win.GWL_STYLE,
		win.GetWindowLong(window.Handle(), win.GWL_STYLE) & ^win.WS_MAXIMIZEBOX & ^win.WS_THICKFRAME,
	)
	window.Run()
}
