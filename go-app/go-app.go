package main

import "github.com/maxence-charriere/go-app/v6/pkg/app"

// go-app是一个使用 Go + WebAssembly 技术编写渐进式 Web 应用的库。
// WebAssembly 是一种可以运行在现代浏览器中的新式代码。
// 可以使用 C/C++/Rust/Go 等高级语言编写 WebAssembly 代码

type Greeting struct {
	app.Compo
	name string
}

func (g *Greeting) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			app.H1().Body(
				app.Text("Hello, "),
				app.If(g.name != "",
					app.Text(g.name),
				).Else(
					app.Text("World"),
				),
			),
		),
		app.Input().Value(g.name).Placeholder("What is your name?").AutoFocus(true).OnChange(g.OnInputChange))
}

func (g *Greeting) OnInputChange(src app.Value, e app.Event) {
	g.name = src.Get("value").String()
	g.Update()
}

func main() {
	app.Route("/", &Greeting{})
	app.Run()
}
