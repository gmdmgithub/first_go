package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("./08_xls/data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell := xlsx.GetCellValue("data", "B2")
	fmt.Println(cell)

	// Get all the rows in the tab - name data.
	// rows is two-dimensional string array - [][]string - rows/cells
	rows := xlsx.GetRows("data")

	//sample form snipped
	// for _, row := range xlsx.GetRows("Sheet1") {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(xlsx.GetCellValue("data", "A1"))
	var firstCol []int
	var secondCol []int
	for _, row := range rows {
		for j, colCell := range row {
			// fmt.Print(colCell.GetCellValue())
			val, err := strconv.ParseInt(colCell, 0, 0)
			if err != nil {
				continue
			}
			if j == 0 {
				firstCol = append(firstCol, int(val))
			} else if j == 1 {
				secondCol = append(secondCol, int(val))
			}
		}
	}

	printSeries(firstCol, xlsx, "town1")
	printSeries(secondCol, xlsx, "town2")

}

func printSeries(col []int, xlsx *excelize.File, sheetName string) {

	series := [8][12]int{}
	for _, val := range col {
		var year int
		year = (val / 100)
		month := val - year*100 - 1
		year -= 2008
		// println(year, month)
		if year < 8 && month < 12 {
			series[year][month] = series[year][month] + 1
		}
	}
	xlsx.NewSheet(sheetName)

	for i := 0; i < 8; i++ {
		for j := 0; j < 12; j++ {
			// start from B column (unicode 66) and second row (in excel starts from 1 index)
			xlsx.SetCellValue(sheetName, fmt.Sprintf("%c%d", 66+i, j+2), series[i][j])
		}
	}

	// fmt.Println(series, sheetName)

	err := xlsx.Save()
	if err != nil {
		fmt.Println(err)
	}

}
