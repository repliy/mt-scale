package utils

import (
	"fmt"
	"mt-scale/models/vo"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// WriteToExcelFile Generate excel file
func WriteToExcelFile(datas []vo.BoxTallyVo) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	for i, data := range datas {
		colName := numToCol(i + 1)
		cellRow1 := colName + "1"
		f.SetCellValue("Sheet2", cellRow1, data.Type)
		cellRow2 := colName + "2"
		f.SetCellValue("Sheet2", cellRow2, data.Num)
		for j, weight := range data.Weights {
			cellRow := colName + strconv.Itoa(j+3)
			f.SetCellValue("Sheet2", cellRow, weight.Weight)
		}
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("test.xlsx"); err != nil {
		fmt.Println(err)
	}
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
