package main

import (
	"os"
	"io"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func visualizePoints(points []Point){
	scatter3d := charts.NewScatter3D()
	scatter3d.Chart3D.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Closest Pair Problem : 3D Visualization",
			Subtitle: "Author : Addin Munawwar Yusuf and Fakih Anugerah Pratama", 
			Top: "5%",
			Bottom: "5%",
		}),
	)

	scatter3d.AddSeries("scatter3d", genScatter3dData(points))
	renderToPage(scatter3d)
}

func renderToPage(chart *charts.Scatter3D) {
	page := components.NewPage()
	page.AddCharts(
		chart,
	)

	f, err := os.Create("visualization.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}


func genScatter3dData(points []Point) []opts.Chart3DData {
	var data []opts.Chart3DData
	for i, point := range points {
		data = append(data, opts.Chart3DData{
			Name: "Point " + string(i),
			Value: []interface{}{point.x, point.y, point.z},
			ItemStyle : &opts.ItemStyle{	
				Color: "#1ecbe1",
				Opacity: 1,
			},
		})
	}
	return data
}

// func main() {
// 	points := []Point{ *NewPoint(1, 2, 3), *NewPoint(4, 5, 6), *NewPoint(7, 8, 9), *NewPoint(10, 11, 12)}
// 	visualizePoints(points)
// }