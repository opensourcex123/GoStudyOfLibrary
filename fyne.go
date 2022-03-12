package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)
// fyne 是 Go 语言编写的跨平台的 UI 库，它可以很方便地移植到手机设备上。
// fyne使用上非常简单，同时它还提供fyne命令打包静态资源和应用程序
func main() {
	myApp := app.New()

	myWin := myApp.NewWindow("Hello")
	myWin.SetContent(widget.NewLabel("Hello Fyne"))
	myWin.Resize(fyne.NewSize(200, 200))
	myWin.ShowAndRun()
}