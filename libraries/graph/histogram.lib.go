package graph

import (
	"fmt"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func DrawHistogram(title string, data []string, xAxisName string, yDataName string, fileName string) error {
	fmt.Printf("title: %s, xName: %s, yName: %s, fileName: %s ", title, xAxisName, yDataName, fileName)
	plot := plot.New()

	valueSlice := make(plotter.Values, len(data))

	for i := 1; i < len(data); i += 1 {
		dataValue, err := strconv.ParseFloat(data[i], 64)

		if err != nil {
			fmt.Println("Failed To Parse Data Element into Float64", err)

			return err
		}

		valueSlice[i] = dataValue
	}

	plot.Title.Text = title
	plot.Title.Padding = 5
	// plot.Title.TextStyle.Font.Style = "Regular"
	plot.X.Label.Text = xAxisName
	plot.Y.Label.Text = yDataName
	plot.X.Label.Padding = 5
	plot.Y.Label.Padding = 5

	plot.Title.TextStyle.Font.Size = 30
	plot.X.Label.TextStyle.Font.Size = 20
	plot.Y.Label.TextStyle.Font.Size = 20

	histogram, histoErr := plotter.NewHist(valueSlice, 20)

	if histoErr != nil {
		fmt.Println("Creating New Histogram Error", histoErr)

		return histoErr
	}

	plot.Add(histogram)

	plotErr := plot.Save(40*vg.Centimeter, 40*vg.Centimeter, fileName+"png")

	if plotErr != nil {
		fmt.Println("Saving New Histogram Failed", plotErr)

		return plotErr
	}

	fmt.Println("Saving New Histogram Completed")

	return nil
}
