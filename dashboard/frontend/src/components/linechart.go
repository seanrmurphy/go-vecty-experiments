package components

import (
	"log"
	"syscall/js"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/seanrmurphy/go-echarts/charts"
)

type LineChartView struct {
	vecty.Core
	Label     string `vecty:"prop"`
	ID        string `vecty:"prop"`
	LineChart *charts.Line
}

// createLineChart creates a simple line chart which contains random points
func createLineChart() *charts.Line {
	line := charts.NewLine()
	line.AddXAxis(nameItems).AddYAxis("Series A", randInt())
	return line
}

// Initialize generates an ID for the div as well as the line chart itself
func (l *LineChartView) Initialize() {
	l.ID = genChartID()
	l.LineChart = createLineChart()
	l.Label = "(a) Line Chart (ECharts)"
}

// Render is the function which gets called by vecty ro rendeer this component;
// it's a vew divs around the div which will contain the echart
func (l *LineChartView) Render() vecty.ComponentOrHTML {
	returnDiv := elem.Div(
		vecty.Markup(
			vecty.Class("content-panel"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("container"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("chart-panel"),
					prop.ID(l.ID),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("chart-label"),
			),
			elem.Label(
				vecty.Text(l.Label),
			),
		),
	)
	return returnDiv
}

// Mount is the part of the component interface which gets called after the
// div has been put on the page; only then can the echart be inserted
func (l *LineChartView) Mount() {
	if l.ID == "" {
		log.Printf("Component not initialized...ignoring...\n")
		return
	}
	o := l.LineChart.GenerateOptions()

	echarts := js.Global().Get("echarts")
	div := js.Global().Get("document").Call("querySelector", "#"+l.ID)
	chart := echarts.Call("init", div, "white")
	chart.Call("setOption", js.ValueOf(o))
}
