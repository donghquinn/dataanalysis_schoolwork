package bostonhousing

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

/*
CRIM     per capita crime rate by town
ZN         proportion of residential land zoned for lots over 25,000 sq.ft.
INDUS    proportion of non-retail business acres per town.
CHAS     Charles River dummy variable (1 if tract bounds river; 0 otherwise)
NOX       nitric oxides concentration (parts per 10 million)
RM        average number of rooms per dwelling
AGE       proportion of owner-occupied units built prior to 1940
DIS       weighted distances to five Boston employment centres
RAD      index of accessibility to radial highways
TAX      full-value property-tax rate per $10,000
PTRATIO pupil-teacher ratio by town
LSTAT   % lower status of the population
MEDV    Median value of owner-occupied homes in $1000
*/
var (
	crime    []string
	zn       []string
	indus    []string
	chas     []string
	nox      []string
	rm       []string
	age      []string
	distance []string
	radial   []string
	tax      []string
	ptratio  []string
	lstat    []string
	medv     []string
)

func OpenFile(fileName string) ([]string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string, []string) {
	f, err := excelize.OpenFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	rows := getRowData(f, "Data")

	generateDataFrame(rows)

	return crime, zn, indus, chas, nox, rm, age, distance, radial, tax, ptratio, lstat, medv
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

		crime = append(crime, columns[0])
		zn = append(zn, columns[1])
		indus = append(indus, columns[2])
		chas = append(chas, columns[3])
		rm = append(rm, columns[4])
		nox = append(nox, columns[5])
		age = append(age, columns[6])
		distance = append(distance, columns[7])
		tax = append(tax, columns[8])
		radial = append(radial, columns[9])
		ptratio = append(ptratio, columns[10])
		lstat = append(lstat, columns[11])
		medv = append(medv, columns[12])
	}

	defer rows.Close()

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("Boston Housing Data Loaded")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@")
}
