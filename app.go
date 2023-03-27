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
	whichDataSet1   string
	whichDataSet2   string
	usingDataSet1   []string
	usingDataSet2   []string
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
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("그래프를 그리길 원하십니까? y / n")
	fmt.Scanln(&wantToDrawGraph)

	graphTypes := graphStyles{histogram: "히스토그램", barChart: "막대그래프", scatter: "산포도"}

	if wantToDrawGraph == "y" {
		fmt.Println("어떤 종류의 그래프를 원하십니까: ", graphTypes)
		fmt.Scanln(&graph)

		switch graph {
		case graphTypes.histogram:
			callHistogram()
		case graphTypes.scatter:
			callScatter()
		default:
			fmt.Println("No Graph Is matching.")
		}

	} else if wantToDrawGraph == "n" {
		fmt.Println("프로세스를 종료합니다.")
	}

	// fmt.Println("Please Input two DataSets you want to draw Histogram")
}

func callHistogram() {
	fmt.Println("생성할 그래프의 제목을 알려주세요")
	_, titleErr := fmt.Scan(&graphTitle)

	if titleErr != nil {
		fmt.Println("Invalid Input", titleErr)
		bufio.NewReader(os.Stdin)
		panic(titleErr.Error())
	}

	fmt.Println("제목: ", graphTitle)

	seletFirstData()

	fmt.Println("사용할 x축의 이름을 알려주세요.")
	_, xAxisErr := fmt.Scan(&xData)

	if xAxisErr != nil {
		fmt.Println("x Axis Data Failed", xAxisErr)
		bufio.NewReader(os.Stdin)
		panic(xAxisErr.Error())
	}

	fmt.Println("사용할 y축의 이름을 알려주세요")
	_, yAxisErr := fmt.Scan(&yData)

	if yAxisErr != nil {
		fmt.Println("y Data Failed", yAxisErr)
		bufio.NewReader(os.Stdin)
		panic(yAxisErr.Error())
	}

	fmt.Println("생성할 파일의 이름을 알려주세요")
	_, fileNameErr := fmt.Scan(&fileName)

	if fileNameErr != nil {
		fmt.Println("Set fileName Failed", fileNameErr)
		bufio.NewReader(os.Stdin)
		panic(fileNameErr.Error())
	}

	histogramErr := drawGraph.DrawHistogram(graphTitle, usingDataSet1, xData, yData, fileName)

	if histogramErr != nil {
		fmt.Println("HistoGram Error", histogramErr)
		bufio.NewReader(os.Stdin)
		panic(histogramErr.Error())
	}
}

func callScatter() {
	fmt.Println("생성할 그래프의 제목을 알려주세요")
	_, titleErr := fmt.Scan(&graphTitle)

	if titleErr != nil {
		fmt.Println("Invalid Input", titleErr)
		bufio.NewReader(os.Stdin)
		panic(titleErr.Error())
	}

	fmt.Println("제목: ", graphTitle)

	fmt.Println("X 축 데이터를 선택해 주세요")
	seletFirstData()

	fmt.Println("Y 축 데이터를 선택해 주세요")
	seletSecondData()

	fmt.Println("사용할 x축의 이름을 알려주세요.")
	_, xAxisErr := fmt.Scan(&xData)

	if xAxisErr != nil {
		fmt.Println("x Axis Data Failed", xAxisErr)
		bufio.NewReader(os.Stdin)
		panic(xAxisErr.Error())
	}

	fmt.Println("사용할 y축의 이름을 알려주세요")
	_, yAxisErr := fmt.Scan(&yData)

	if yAxisErr != nil {
		fmt.Println("y Data Failed", yAxisErr)
		bufio.NewReader(os.Stdin)
		panic(yAxisErr.Error())
	}

	fmt.Println("생성할 파일의 이름을 알려주세요")
	_, fileNameErr := fmt.Scan(&fileName)

	if fileNameErr != nil {
		fmt.Println("Set fileName Failed", fileNameErr)
		bufio.NewReader(os.Stdin)
		panic(fileNameErr.Error())
	}

	histogramErr := drawGraph.DrawScatter(graphTitle, usingDataSet1, usingDataSet2, xData, yData, fileName)

	if histogramErr != nil {
		fmt.Println("Scatter Error", histogramErr)
		bufio.NewReader(os.Stdin)
		panic(histogramErr.Error())
	}
}

func seletFirstData() {
	fmt.Println("사용할 데이터 셋을 알려주세요: totalValue, tax, lotSqft, yrBuilt, grossArea, livingArea, floors, rooms, bedRooms, fullBath, halfBath, kitchen, firePlace, remodel")
	_, selectDataSetErr := fmt.Scanln(&whichDataSet1)

	if selectDataSetErr != nil {
		fmt.Println("Select DataSet Input Error", selectDataSetErr)
		bufio.NewReader(os.Stdin)
		panic(selectDataSetErr.Error())
	}

	switch whichDataSet1 {
	case "totalValue":
		usingDataSet1 = totalValue
	case "tax":
		usingDataSet1 = tax
	case "lotSqrt":
		usingDataSet1 = lotSqft
	case "yrBuilt":
		usingDataSet1 = yrBuilt
	case "grossArea":
		usingDataSet1 = grossArea
	case "livingArea":
		usingDataSet1 = livingArea
	case "floors":
		usingDataSet1 = floors
	case "fullBath":
		usingDataSet1 = fullBath
	case "halfBath":
		usingDataSet1 = halfBath
	case "kitchen":
		usingDataSet1 = kitchen
	case "firePlace":
		usingDataSet1 = firePlace
	case "remodel":
		usingDataSet1 = remodel
	}
	fmt.Println("Selected DataSet: ", whichDataSet1)
}

func seletSecondData() {
	fmt.Println("사용할 데이터 셋을 알려주세요: totalValue, tax, lotSqft, yrBuilt, grossArea, livingArea, floors, rooms, bedRooms, fullBath, halfBath, kitchen, firePlace, remodel")
	_, selectDataSetErr := fmt.Scanln(&whichDataSet2)

	if selectDataSetErr != nil {
		fmt.Println("Select DataSet Input Error", selectDataSetErr)
		bufio.NewReader(os.Stdin)
		panic(selectDataSetErr.Error())
	}

	switch whichDataSet2 {
	case "totalValue":
		usingDataSet2 = totalValue
	case "tax":
		usingDataSet2 = tax
	case "lotSqrt":
		usingDataSet2 = lotSqft
	case "yrBuilt":
		usingDataSet2 = yrBuilt
	case "grossArea":
		usingDataSet2 = grossArea
	case "livingArea":
		usingDataSet2 = livingArea
	case "floors":
		usingDataSet2 = floors
	case "fullBath":
		usingDataSet2 = fullBath
	case "halfBath":
		usingDataSet2 = halfBath
	case "kitchen":
		usingDataSet2 = kitchen
	case "firePlace":
		usingDataSet2 = firePlace
	case "remodel":
		usingDataSet2 = remodel
	}
	fmt.Println("Selected DataSet: ", whichDataSet2)
}
