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

	for i := range valueSlice {
		dataValue, err := strconv.ParseFloat(data[i+1], 64)

		if err != nil {
			fmt.Println("Failed To Parse Data Element into Float64", err)

			return err
		}

		valueSlice[i] = dataValue
	}

	plot.Title.Text = title
	plot.X.Label.Text = xAxisName
	plot.Y.Label.Text = yDataName

	histogram, histoErr := plotter.NewHist(valueSlice, 10)

	if histoErr != nil {
		fmt.Println("Creating New Histogram Error", histoErr)

		return histoErr
	}

	plot.Add(histogram)

	plotErr := plot.Save(30*vg.Centimeter, 30*vg.Centimeter, fileName)

	if plotErr != nil {
		fmt.Println("Saving New Histogram Failed", plotErr)

		return plotErr
	}

	fmt.Println("Saving New Histogram Completed")

	return nil
}
