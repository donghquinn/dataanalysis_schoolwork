package graph

import (
	"fmt"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func DrawScatter(title string, yData []string, xData []string, xAxisName string, yDataName string, fileName string) error {
	fmt.Printf("title: %s, xName: %s, yName: %s, fileName: %s ", title, xAxisName, yDataName, fileName)

	plot := plot.New()

	valueSlice := make(plotter.XYs, len(xData))

	for i := 1; i < len(xData); i += 1 {
		yDataValue, yDataErr := strconv.ParseFloat(yData[i], 64)
		xDataValue, xDataErr := strconv.ParseFloat(xData[i], 64)

		if yDataErr != nil {
			fmt.Println("Failed To Parse Y Data Element into Float64", yDataErr)

			return yDataErr
		}

		if xDataErr != nil {
			fmt.Println("Failed To Parse X Data Elememnt into Float 64", xDataErr)

			return xDataErr
		}

		valueSlice[i].X = xDataValue
		valueSlice[i].Y = yDataValue
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

	scatter, scatterErr := plotter.NewScatter(valueSlice)

	if scatterErr != nil {
		fmt.Println("Scatter Error", scatterErr)

		return scatterErr
	}

	plot.Add(scatter)

	plotErr := plot.Save(40*vg.Centimeter, 40*vg.Centimeter, fileName+".png")

	if plotErr != nil {
		fmt.Println("Saving Scatter Error", plotErr)

		return plotErr
	}

	return nil
}
