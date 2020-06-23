package components

import (
	"log"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/seanrmurphy/go-echarts/charts"
)

// BarChartView is a vecty component which creates the bar chart once the
// div has been rendered, using the Mount() interface
type BarChartView struct {
	vecty.Core
	Label    string `vecty:"prop"`
	ID       string `vecty:"prop"`
	BarChart *charts.Bar
}

var nameItems = []string{"Red", "Orange", "Blue", "Yellow", "White", "Black"}
var maxNum = 50
var seed = rand.NewSource(time.Now().UnixNano())

// randInt generates a small set of integrers between 1 and maxNum which will
// be used for the bar chart data
func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, int(seed.Int63())%maxNum)
	}
	return r
}

// createBarChart creates the bar chart using the go-echarts library
func createBarChart() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("Series A", randInt()).
		AddYAxis("Series B", randInt())
	return bar
}

// Initialize creates an ID for the bar chart and creates a bar chart with
// containing random data
func (g *BarChartView) Initialize() {
	g.ID = genChartID()
	g.BarChart = createBarChart()
	g.Label = "(c) Bar Chart (ECharts)"
}

// Render renders this component when called by vecty rendering - the component
// simply contains a couple of divs around the echart; styling of these divs is
// in the CSS
func (g *BarChartView) Render() vecty.ComponentOrHTML {
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
					prop.ID(g.ID),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("chart-label"),
			),
			elem.Label(
				vecty.Text(g.Label),
			),
		),
	)
	return returnDiv
}

// Mount is called when the div is rendered; we need to wait until this time
// because we need to pass the div object to echarts - if it does not exist, obviously
// we cannot pass to echarts
func (g *BarChartView) Mount() {
	if g.ID == "" {
		log.Printf("Component not initialized...ignoring...\n")
		return
	}
	o := g.BarChart.GenerateOptions()

	echarts := js.Global().Get("echarts")
	div := js.Global().Get("document").Call("querySelector", "#"+g.ID)
	chart := echarts.Call("init", div, "white")
	chart.Call("setOption", js.ValueOf(o))
}
