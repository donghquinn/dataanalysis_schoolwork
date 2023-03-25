package westroxbury

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

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
)

func OpenFile(fileName string) ([]string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string) {
	f, err := excelize.OpenFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	rows := getRowData(f, "Data")

	generateDataFrame(rows)

	return totalValue, tax, lotSqft, yrBuilt, grossArea, livingArea, floors, rooms, bedRooms, fullBath, halfBath, kitchen, firePlace, remodel
}

func getRowData(f *excelize.File, sheetName string) *excelize.Rows {
	rows, rowErr := f.Rows(sheetName)

	if rowErr != nil {
		fmt.Println(rowErr)
	}

	return rows
}

func generateDataFrame(rows *excelize.Rows) {
	for rows.Next() {
		columns, columnErr := rows.Columns()

		if columnErr != nil {
			fmt.Println(columnErr)
		}

		totalValue = append(totalValue, columns[0])
		tax = append(tax, columns[1])
		lotSqft = append(lotSqft, columns[2])
		yrBuilt = append(yrBuilt, columns[3])
		grossArea = append(grossArea, columns[4])
		livingArea = append(livingArea, columns[5])
		floors = append(floors, columns[6])
		rooms = append(rooms, columns[7])
		bedRooms = append(bedRooms, columns[8])
		fullBath = append(fullBath, columns[9])
		halfBath = append(halfBath, columns[10])
		kitchen = append(kitchen, columns[11])
		firePlace = append(firePlace, columns[12])
		remodel = append(remodel, columns[13])

	}

	defer rows.Close()

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("West RoxBury Data Loaded")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@")
}
