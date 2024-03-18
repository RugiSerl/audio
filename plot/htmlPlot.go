package plot

import (
	"fmt"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateItems(data [][]float64) []opts.HeatMapData {
	items := make([]opts.HeatMapData, len(data)*len(data[0]))
	for t := 0; t < len(data); t++ {
		for i := 0; i < len(data[t]); i++ {
			items[t*len(data[0])+i] = opts.HeatMapData{
				Name:  fmt.Sprintf("%d", i),
				Value: []interface{}{t, i, data[t][i]},
			}
		}
	}
	return items
}

func heatMapBase(data [][]float64) *charts.HeatMap {
	hm := charts.NewHeatMap()
	hm.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "basic heatmap example",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Type:      "category",
			SplitArea: &opts.SplitArea{Show: true},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type:      "category",
			SplitArea: &opts.SplitArea{Show: true},
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Min:        0,
			Max:        10000,
			InRange: &opts.VisualMapInRange{
				Color: []string{"#d94e5d", "#eac736", "#50a3ba"},
			},
		}),
	)

	hm.AddSeries("heatmap", generateItems(data))
	return hm
}

func CreatePlot(data [][]float64, name string) {
	page := components.NewPage()
	page.AddCharts(
		heatMapBase(data),
	)
	f, err := os.Create("assets/" + name + ".html")

	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
