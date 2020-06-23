package main

import (
	"github.com/gopherjs/vecty"
	"github.com/seanrmurphy/go-vecty-experiments/dashboard/frontend/src/app/model"
	"github.com/seanrmurphy/go-vecty-experiments/dashboard/frontend/src/components"
)

var (
	State model.ApplicationData
)

func main() {

	vecty.SetTitle("Simple Web App")
	vecty.AddStylesheet("/css/app.css")

	p := &components.PageView{}
	p.Initialize()

	vecty.RenderBody(p)
}
