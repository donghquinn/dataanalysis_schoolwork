package graph

import (
	"fmt"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func DrawHistogram(data []string, xColumnName string, yColumnName string) {
	plot := plot.New()

	valueSlice := make(plotter.Values, len(data))

	for i := range valueSlice {
		dataValue, err := strconv.ParseFloat(data[i], 64)

		if err != nil {
			fmt.Println("Failed To Parse Data Element into Float64", err)
		}

		valueSlice[i] = dataValue
	}

	plot.Title.Text = columnName
}
