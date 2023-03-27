package graph

import (
	"fmt"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func DrawScatter(title string, data []string, xAxisName string, yDataName string, fileName string) error {
	fmt.Printf("title: %s, xName: %s, yName: %s, fileName: %s ", title, xAxisName, yDataName, fileName)

	plot := plot.New()

	valueSlice := make(plotter.XYs, len(data))

	for i := 1; i < len(data); i += 1 {
		dataValue, err := strconv.ParseFloat(data[i], 64)

		if err != nil {
			fmt.Println("Failed To Parse Data Element into Float64", err)

			return err
		}

		valueSlice[i].X = dataValue
	}

	plot.Title.Text = title
	plot.X.Label.Text = xAxisName
	plot.Y.Label.Text = yDataName

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
