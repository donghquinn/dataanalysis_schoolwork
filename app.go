package main

import (
	"github.com/donghquinn/xlsx/libraries"
	westroxbury "github.com/donghquinn/xlsx/westRoxbury"
)

func main() {
	fileName := "/Users/kimdonghyun/Desktop/schoolWork/23-1/비즈니스 애널리틱스/DMBA-3rd-ed-datasets-xlsx_2/West Roxbury.xlsx"

	// 배열 첫번째는 칼럼 이름
	totalValue, tax, lotSqft, yrBuilt, grossArea, livingArea, floors, rooms, bedRooms, fullBath, halfBath, kitchen, firePlace, remodel := westroxbury.OpenFile(fileName)

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
}
