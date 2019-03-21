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
	rows := xlsx.GetRows("data")
	// fmt.Println(xlsx.GetCellValue("data", "A1"))
	fmt.Println("Results as as follow")
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
		// if len(firstCol) > i {
		// 	fmt.Print(firstCol[i])
		// }
		// if len(secondCol) > i {
		// 	fmt.Printf(" %d", secondCol[i])
		// }
		// fmt.Print("\n")
	}

	printSeries(firstCol)
	printSeries(secondCol)

	// seriesOne := [8][12]int{}
	// for _, val := range firstCol {
	// 	var year int
	// 	year = (val / 100)
	// 	month := val - year*100 - 1
	// 	year -= 2008
	// 	println(year, month)
	// 	if year < 8 && month < 12 {
	// 		seriesOne[year][month] = seriesOne[year][month] + 1
	// 	}
	// }

	// fmt.Println(seriesOne)

}

func printSeries(col []int) {

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

	fmt.Println(series)

}
