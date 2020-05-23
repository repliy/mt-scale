package utils

import (
	"fmt"
	"mt-scale/models/vo"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// WriteToExcelFile Generate excel file
func WriteToExcelFile(datas []vo.BoxTallyVo) string {
	f := excelize.NewFile()
	// Create a new sheet.
	const sheet string = "Sheet1"
	index := f.NewSheet(sheet)
	dataStyle, _ := f.NewStyle(
		`{
			"border":[
				{
					"type":"left",
					"color":"000000",
					"style":2
				},
				{
					"type":"top",
					"color":"000000",
					"style":2
				},
				{
					"type":"bottom",
					"color":"000000",
					"style":2
				},
				{
					"type":"right",
					"color":"000000",
					"style":2
				}
			],
			"alignment":{
				"horizontal":"center",
				"vertical":"center",
				"wrap_text":true
			},
			"font":{
				"bold":true,
				"family":"Berlin Sans FB Demi",
				"size":12
			}
		}`)
	titleStyle, _ := f.NewStyle(
		`{
			"border": [
				{
					"type":"bottom",
					"color":"000000",
					"style":2
				}
			],
			"alignment":{
				"horizontal":"center",
				"vertical":"center",
				"wrap_text":true
			},
			"font":{
				"bold":true,
				"family":"Berlin Sans FB Demi",
				"size":16
			}
		}`)
	subtitleStyle, _ := f.NewStyle(
		`{
			"border": [
				{
					"type":"bottom",
					"color":"000000",
					"style":2
				}
			],
			"alignment":{
				"horizontal":"start",
				"vertical":"center",
				"wrap_text":true
			},
			"font":{
				"bold":true,
				"family":"Berlin Sans FB Demi",
				"size":14
			}
		}`)
	// title
	f.MergeCell(sheet, "A1", "C2")
	f.SetCellValue(sheet, "A1", "Vessel Plant Tally")
	f.SetCellStyle(sheet, "A1", "A1", titleStyle)
	// Date
	f.MergeCell(sheet, "A3", "C3")
	f.SetCellValue(sheet, "A3", time.Now().Format("2006/1/2"))
	f.SetCellStyle(sheet, "A3", "C3", subtitleStyle)
	f.SetCellValue(sheet, "A4", "Date")
	f.SetCellStyle(sheet, "A4", "A4", subtitleStyle)

	// Name

	// Plant

	const dataStartRow int = 5
	maxRows := dataStartRow
	for i, data := range datas {
		// from A
		colName := numToCol(i + 1)
		// A5
		cellRow1 := colName + strconv.Itoa(dataStartRow)
		f.SetCellValue(sheet, cellRow1, data.Type)
		// A6
		cellRow2 := colName + strconv.Itoa(dataStartRow+1)
		f.SetCellValue(sheet, cellRow2, data.Num)
		for j, weight := range data.Weights {
			cellRow := colName + strconv.Itoa(j+dataStartRow+2)
			f.SetCellValue(sheet, cellRow, weight.Weight)
			if (j + dataStartRow) > maxRows {
				maxRows = j + dataStartRow
			}
		}
	}
	// border
	for i := range datas {
		// from A
		colName := numToCol(i + 1)
		for j := 0; j < maxRows; j++ {
			cellRow := colName + strconv.Itoa(j+dataStartRow)
			f.SetCellStyle(sheet, cellRow, cellRow, dataStyle)
		}
	}
	// line height
	for i := 0; i < (maxRows + dataStartRow); i++ {
		f.SetRowHeight(sheet, (i + 1), 20)
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	var file string = "test.xlsx"
	if err := f.SaveAs(file); err != nil {
		fmt.Println(err)
	}
	return file
}

func numToCol(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int
	)
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 {
		for {
			k = Num % 26
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26
			if Num <= 26 {
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}

	for _, value := range temp {
		Str = Slice[value] + Str
	}
	return Str
}
