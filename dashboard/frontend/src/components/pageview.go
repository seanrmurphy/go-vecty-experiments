package components

import (
	"log"
	"math"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/nathanhack/svg"
	"github.com/nathanhack/svg/attr"
	"github.com/nathanhack/svg/attr/path"
	"github.com/nathanhack/svg/svgelem"

	"github.com/seanrmurphy/go-vecty-experiments/dashboard/frontend/src/app/model"
)

type GraphData struct {
	minX   int
	maxX   int
	minY   int
	maxY   int
	points [][]int
}

type PieChartData struct {
	points []int
}

type BarChartData struct {
	points []int
}

// The main components of the page view are the visible tab and the set of
// individual e-charts based components
type PageView struct {
	vecty.Core

	app          *model.ApplicationData
	visibleTab   int
	GraphData    GraphData
	PieChartData PieChartData
	BarChartData BarChartData
}

// InitializeData generates some initial data for the line, bar and graph charts;
// this data is only used by the SVG graphics - the e-charts based charts have
// seperate initialization
func (p *PageView) InitializeData() {
	points := [][]int{
		{0, 0},
		{10, 10},
		{50, 20},
		{75, 25},
		{100, 27},
	}
	p.GraphData = GraphData{
		minX:   0,
		maxX:   100,
		minY:   0,
		maxY:   100,
		points: points,
	}
	p.PieChartData = PieChartData{
		points: []int{66, 66, 66},
	}
	p.BarChartData = BarChartData{
		points: []int{20, 30, 60, 95},
	}
}

// Initialize just initializes the data
func (p *PageView) Initialize() {
	p.InitializeData()
}

// renderSVGClicked is a call back which is triggered when the render svg menu
// item is cliekd
func (p *PageView) renderSVGClicked(e *vecty.Event) {
	p.visibleTab = 0
	vecty.Rerender(p)
}

// renderEChartsClicked is a callback which is triggered when the render echarts
// menu item on the nav bar is clicked
func (p *PageView) renderEChartsClicked(e *vecty.Event) {
	p.visibleTab = 1
	vecty.Rerender(p)
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		p.renderHeader(),
		elem.Section(
			vecty.Markup(
				vecty.Class("main-area"),
			),
			p.renderMainSection(),
		),
	)
}

// renderHeader renders the page header
func (p *PageView) renderHeader() *vecty.HTML {
	return elem.Header(
		vecty.Markup(
			vecty.Class("header"),
		),

		elem.Heading1(
			vecty.Text("Vecty Experimentation - Dashboard Demo "),
		),
	)
}

// renderMainSection renders the main section of the page which consists of
// a navbar on the side and a main panel
func (p *PageView) renderMainSection() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("main-section"),
		),
		p.renderNavbar(),
		p.renderPanel(),
	)
}

// renderNavbar renders a simple navbar on the side, using styling in the CSS. It
// also sets up a couple of callbacks which are triggered when items in the navbar
// are clicked.
func (p *PageView) renderNavbar() *vecty.HTML {
	return elem.Aside(
		vecty.Markup(
			vecty.Class("menu"),
		),
		elem.UnorderedList(
			vecty.Markup(
				vecty.Class("menu-list"),
			),
			elem.ListItem(
				vecty.Markup(
					event.Click(p.renderSVGClicked),
				),
				elem.Label(
					vecty.Text("Render SVG"),
				),
			),
			elem.ListItem(
				vecty.Markup(
					event.Click(p.renderEChartsClicked),
				),
				elem.Label(
					vecty.Text("Render EChart"),
				),
			),
		),
	)
}

// renderPanel renders the panel which is visible.
func (p *PageView) renderPanel() *vecty.HTML {
	return elem.Div(
		p.renderVisiblePanel(),
	)
}

// renderVisiblePanel determines which tab should be displayed and returns that
func (p *PageView) renderVisiblePanel() *vecty.HTML {
	switch p.visibleTab {
	case 0:
		return p.renderSVG()
	case 1:
		return p.renderECharts()
	}

	return p.renderSVG()
}

// renderSVG shows the svg based 'charts'
func (p *PageView) renderSVG() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("main-panel"),
		),
		elem.Heading1(
			elem.Label(
				vecty.Text("SVG Charts!"),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("chart-panels"),
			),
			p.renderLineChart(),
			p.renderPieChart(),
			p.renderBarChart(),
		),
	)
}

// renderECharts renders the echarts using the linechart, piechart and barchart
// components
func (p *PageView) renderECharts() *vecty.HTML {

	linechart := LineChartView{}
	linechart.Initialize()
	piechart := PieChartView{}
	piechart.Initialize()
	barchart := BarChartView{}
	barchart.Initialize()

	return elem.Div(
		vecty.Markup(
			vecty.Class("main-panel"),
		),
		elem.Heading1(
			elem.Label(
				vecty.Text("ECharts Graphs"),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("chart-panels"),
			),
			&linechart,
			&piechart,
			&barchart,
		),
	)
}

// renderLineChart creates a very simple SVG representation of a line chart using
// the svg Path Element, puts it into the panel with a caption
func (p *PageView) renderLineChart() *vecty.HTML {
	pathString := p.calculatePathForLineChart(p.GraphData)

	image := []svg.Component{
		attr.Class("svg-image"),
		attr.Width(300),
		attr.Height(300),
		attr.ViewBox(0, 0, 100, 100),
		attr.Stroke("red"),
		svgelem.Path(attr.Stroke("red"), attr.Fill("none"), attr.D(pathString)),
	}

	return p.chartPanel(image, "(a) Line Chart (SVG)")
}

// calculatePathForLineChart calculates an SVG Path for the linechart, premised
// on the assumption that the max width and height of the chart are 100 and 100.
func (p *PageView) calculatePathForLineChart(g GraphData) (pa path.Cmd) {
	pa = path.M(0, 100)
	width := 100
	height := 100
	for _, px := range g.points {

		gx := float64(px[0])
		gy := float64(px[1])
		plotx := ((gx - float64(g.minX)) / (float64(g.maxX - g.minX))) * float64(width)
		// this calculation is done because the vertical axis goes from top to bottom
		ploty := float64(height) - ((gy-float64(g.minY))/(float64(g.maxY-g.minY)))*float64(height)
		pa = pa.L(plotx, ploty)
	}
	return
}

// renderPieChart creates the simple SVG-based representation of the pie chart; at
// present, this can only handle three different colours. The sectors are generated
// as paths which are closed and then filled.
func (p *PageView) renderPieChart() *vecty.HTML {
	paths := p.determinePathsForPieChart()

	image := []svg.Component{
		attr.Class("svg-image"),
		attr.Width(300),
		attr.Height(300),
		attr.ViewBox(-50, -50, 100, 100),
		attr.Stroke("red"),
	}

	// this logic only supports 3 segments right now
	if len(paths) > 3 {
		log.Printf("Warning - only 3 colours available for rendering pie chart\n")
		return nil
	}

	colors := []string{"red", "blue", "green"}
	for i, path := range paths {
		image = append(image, svgelem.Path(attr.Stroke(colors[i]), attr.Fill(colors[i]), attr.D(path)))
	}

	return p.chartPanel(image, "(b) Pie Chart (SVG)")
}

// determinePathsForPieChart generates paths for each of the sectors corresponding
// to a piece of the pie. These are centred on 0,0, have 2 straight lines to the
// ends of the circular arc by which they are joined. The path is closed.
func (p *PageView) determinePathsForPieChart() (paths []path.Cmd) {
	// calcTotal
	sum := 0
	for _, px := range p.PieChartData.points {
		sum += px
	}
	var percentages []float64
	for _, px := range p.PieChartData.points {
		percentages = append(percentages, (float64(px) / float64(sum)))
	}
	// having calculated percentages, now we need to calculate paths
	x, y, rx, ry := 0.0, 0.0, 50.0, 50.0
	log.Printf("percentage = %v\n", percentages)
	cumulativePercentage := 0.0
	for _, percent := range percentages {
		p := path.MoveTo(0, 0)

		angle := cumulativePercentage * 2 * math.Pi
		x = float64(rx) * math.Cos(angle)
		y = -float64(ry) * math.Sin(angle)
		p = p.LineTo(x, y)

		cumulativePercentage += percent
		angle = cumulativePercentage * 2 * math.Pi
		x = float64(rx) * math.Cos(angle)
		y = -float64(ry) * math.Sin(angle)
		p = p.Arc(rx, ry, 0.0, 0, 0, x, y)

		p = p.ClosePath()
		paths = append(paths, p)
	}

	return
}

// renderBarChart creates a simple SVG based bar chart, embeds it in a panel
// with a caption and returns it
func (p *PageView) renderBarChart() *vecty.HTML {

	g := []svg.Component{
		attr.Class("svg-image"),
		attr.Width(300),
		attr.Height(300),
		attr.ViewBox(0, 0, 100, 100),
		attr.Stroke("red"),
	}

	w := 100 / len(p.BarChartData.points)
	color := "blue"
	for i, point := range p.BarChartData.points {
		x := w * i
		y := 100 - point
		h := point
		r := svgelem.Rect(attr.Stroke(color), attr.Fill(color), attr.X(x), attr.Y(y), attr.Width(w), attr.Height(h))
		g = append(g, r)
	}

	return p.chartPanel(g, "(c) Bar Chart (SVG)")
}

// chartPanel takes an SVG compponent as an input, wraps it in a couple of divs
// and adds the caption specified in the input string
func (p *PageView) chartPanel(g []svg.Component, s string) *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("content-panel"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("chart-panel"),
			),
			svg.Render(
				svg.SVG(g...),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("chart-label"),
			),
			elem.Label(
				vecty.Text(s),
			),
		),
	)
}
