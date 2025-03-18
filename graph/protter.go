package graph

import (
	"fmt"
	"image/color"

	"dividend-visualizer/models"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	imageWidth  = 1200 // 画像の横幅
	imageHeight = 800  // 画像の縦幅
)

// PlotGraph generates a PNG image from the given data and saves it to a file.
func PlotGraph(data models.DataSet, outputPath string) error {
	p := plot.New()
	p.Title.Text = "2024年"
	p.X.Label.Text = "日付"
	p.Y.Label.Text = "金額"
	// dataのDividendをValues型に変換
	nums := plotter.Values{}
	xLabels := []string{}
	for _, d := range data {
		nums = append(nums, d.Dividend)
		xLabels = append(xLabels, d.DepositDate)
	}

	breadth := vg.Points(15)

	bar, err := plotter.NewBarChart(nums, breadth)
	if err != nil {
		return fmt.Errorf("failed to create bar chart: %w", err)
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = color.RGBA{R: 255, G: 99, B: 71, A: 255}
	p.Add(bar)

	p.NominalX(xLabels...)

	if err := p.Save(imageWidth, imageHeight, outputPath); err != nil {
		return fmt.Errorf("failed to save image")
	}

	return nil
}
