package components

import (
	"log"
	"math/rand"
	"syscall/js"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/seanrmurphy/go-echarts/charts"
)

type PieChartView struct {
	vecty.Core
	Label    string `vecty:"prop"`
	ID       string `vecty:"prop"`
	PieChart *charts.Pie
}

// createPieChart creates a pie chart using the echarts libary
func createPieChart() *charts.Pie {
	pie := charts.NewPie()
	pie.Add("pie", genKvData())
	return pie
}

// genKvData generates some random key value data which is used as input to
// the pie chart
func genKvData() map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(nameItems); i++ {
		m[nameItems[i]] = rand.Intn(maxNum)
	}
	return m
}

// Initialize generates a random ID for the div containing the echart as well
// as creating a pie chart using echarts
func (p *PieChartView) Initialize() {
	p.ID = genChartID()
	p.PieChart = createPieChart()
	p.Label = "(b) Pie Chart (Echarts)"
}

// Render is called by vecty when this component needs to be rendered
func (p *PieChartView) Render() vecty.ComponentOrHTML {
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
					prop.ID(p.ID),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("chart-label"),
			),
			elem.Label(
				vecty.Text(p.Label),
			),
		),
	)
	return returnDiv
}

// Mount is called by vecty when this component has been displayed on the
// page - then we need to populate it with the echart by calling the echart
// library
func (p *PieChartView) Mount() {
	if p.ID == "" {
		log.Printf("Component not initialized...ignoring...\n")
		return
	}
	o := p.PieChart.GenerateOptions()

	echarts := js.Global().Get("echarts")
	div := js.Global().Get("document").Call("querySelector", "#"+p.ID)
	chart := echarts.Call("init", div, "white")
	chart.Call("setOption", js.ValueOf(o))
}
