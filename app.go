package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/donghquinn/xlsx/libraries"
	drawGraph "github.com/donghquinn/xlsx/libraries/graph"
	westroxbury "github.com/donghquinn/xlsx/westRoxbury"
)

type graphStyles struct {
	histogram string
	barChart  string
	scatter   string
}

// type dataSets struct {
// }

var (
	totalValue []string
	tax        []string
	lotSqft    []string
	yrBuilt    []string
	grossArea  []string
	livingArea []string
	floors     []string
	rooms      []string
	bedRooms   []string
	fullBath   []string
	halfBath   []string
	kitchen    []string
	firePlace  []string
	remodel    []string

	wantToDrawGraph string
	graph           string
	whichDataSet    string
	usingDataSet    []string
	xData           string
	yData           string
	graphTitle      string
	fileName        string
)

func main() {
	fileName := "/Users/kimdonghyun/Desktop/schoolWork/23-1/비즈니스 애널리틱스/DMBA-3rd-ed-datasets-xlsx_2/West Roxbury.xlsx"

	// 배열 첫번째는 칼럼 이름
	totalValue, tax, lotSqft, yrBuilt, grossArea, livingArea, floors, rooms, bedRooms, fullBath, halfBath, kitchen, firePlace, remodel = westroxbury.OpenFile(fileName)

	libraries.GetFloatMean(totalValue, "TOTAL VALUE")

	libraries.GetIntMean(tax, "TAX")
	libraries.GetIntMean(lotSqft, "LOT SQFT")
	libraries.GetIntMean(yrBuilt, "YR BUILT")
	libraries.GetIntMean(grossArea, "GROSS AREA")
	libraries.GetIntMean(livingArea, "LIVING AREA")
	libraries.GetFloatMean(floors, "Floors")
	libraries.GetIntMean(rooms, "ROOMS")
	libraries.GetIntMean(bedRooms, "BEDROOMS")
	libraries.GetIntMean(fullBath, "FULL BATH")
	libraries.GetIntMean(halfBath, "HALF BATh")
	libraries.GetIntMean(kitchen, "KITCHEN")
	libraries.GetIntMean(firePlace, "FIREPLACE")

	libraries.RemodeledCount(remodel)

	callGraph()
}

func callGraph() {
	fmt.Println("그래프를 그리길 원하십니까? y / n")
	fmt.Scanln(&wantToDrawGraph)

	graphTypes := graphStyles{histogram: "히스토그램", barChart: "바 차트", scatter: "산포도"}

	if wantToDrawGraph == "y" {
		fmt.Println("어떤 종류의 그래프를 원하십니까: ", graphTypes)
		fmt.Scanln(&graph)

		if graph == graphTypes.histogram {
			fmt.Println("생성할 그래프의 제목을 알려주세요")
			_, titleErr := fmt.Scanln(&graphTitle)

			if titleErr != nil {
				fmt.Println("Invalid Input", titleErr)
				panic(titleErr.Error())
			}

			fmt.Println("제목: ", graphTitle)

			fmt.Println("사용할 데이터 셋을 알려주세요: totalValue, tax, lotSqft, yrBuilt, grossArea, livingArea, floors, rooms, bedRooms, fullBath, halfBath, kitchen, firePlace, remodel")
			_, selectDataSetErr := fmt.Scanln(&whichDataSet)

			if selectDataSetErr != nil {
				fmt.Println("Select DataSet Input Error", selectDataSetErr)
				panic(selectDataSetErr.Error())
			}

			switch whichDataSet {
			case "totalValue":
				usingDataSet = totalValue
			case "tax":
				usingDataSet = tax
			case "lotSqrt":
				usingDataSet = lotSqft
			case "yrBuilt":
				usingDataSet = yrBuilt
			case "grossArea":
				usingDataSet = grossArea
			case "livingArea":
				usingDataSet = livingArea
			case "floors":
				usingDataSet = floors
			case "fullBath":
				usingDataSet = fullBath
			case "halfBath":
				usingDataSet = halfBath
			case "kitchen":
				usingDataSet = kitchen
			case "firePlace":
				usingDataSet = firePlace
			case "remodel":
				usingDataSet = remodel
			}

			fmt.Println("사용할 x축 그래프 데이터를 알려주세요.")
			_, xAxisErr := fmt.Scanln(&xData)

			if xAxisErr != nil {
				fmt.Println("x Axis Data Failed", xAxisErr)
				bufio.NewReader(os.Stdin)
				panic(xAxisErr.Error())
			}

			fmt.Println("사용할 y축 그래프 데이터를 알려주세요")
			_, yAxisErr := fmt.Scanln(&yData)

			if yAxisErr != nil {
				fmt.Println("y Axis Data Failed", yAxisErr)
				bufio.NewReader(os.Stdin)
				panic(yAxisErr.Error())
			}

			fmt.Println("생성할 파일의 이름을 알려주세요")
			fmt.Scanln(&fileName)

			histogramErr := drawGraph.DrawHistogram(graphTitle, usingDataSet, xData, yData, fileName)

			if histogramErr != nil {
				fmt.Println("HistoGram Error", histogramErr)
				bufio.NewReader(os.Stdin)
				panic(histogramErr.Error())
			}

		} else if wantToDrawGraph == "n" {
			fmt.Println("프로세스를 종료합니다.")
		}
	}

	// fmt.Println("Please Input two DataSets you want to draw Histogram")
}
