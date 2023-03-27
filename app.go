package main

import (
	"fmt"

	"github.com/donghquinn/xlsx/libraries"
	drawGraph "github.com/donghquinn/xlsx/libraries/graph"
	westroxbury "github.com/donghquinn/xlsx/westRoxbury"
)

type graphStyles struct {
	histogram string
	barChart  string
	scatter   string
}

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
	xData           string
	yData           string
	graphTitle      string
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
			fmt.Scanln(&graphTitle)

			fmt.Println("사용할 x축 그래프 데이터를 알려주세요.")
			fmt.Scanln(&xData)

			fmt.Println("사용할 y축 그래프 데이터를 알려주세요")
			fmt.Scanln(&yData)

			drawGraph.DrawHistogram("Data", totalValue, "Total Value", "Y")
		}
	}

	fmt.Println("Please Input two DataSets you want to draw Histogram")
}
