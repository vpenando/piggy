package main

// import (
// 	"fmt"
// 	"localization"
// 	"log"
// 	"time"

// 	xl "github.com/360EntSecGroup-Skylar/excelize"
// 	"github.com/vpenando/piggy/piggy"
// )

// // Report is an Excel report that contains each operation
// // of the year.
// // It contains a sheet for each month, even if there is no
// // operation.
// type Report struct {
// 	month      time.Month
// 	operations piggy.Operations
// 	categories []piggy.Category
// 	titles     []string
// }

// // NewReport returns an Excel report for a given year.
// func NewReport(year int, month time.Month) (Report, error) {
// 	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
// 	endDate := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.Local)
// 	operations, err := operationController.ReadAllBetween(startDate, endDate)
// 	if err != nil {
// 		log.Println("Error:", err)
// 		return Report{}, err
// 	}
// 	categories, _ := categoryController.ReadAll()
// 	report := Report{
// 		month:      month,
// 		categories: categories,
// 		operations: operations,
// 	}
// 	return report, err
// }

// func Export(filename string, report Report) error {
// 	file := xl.NewFile()
// 	additionalSheets := file.GetSheetMap()
// 	for month := time.January; month <= time.December; month++ {
// 		if err := exportMonth(file, report, month); err != nil {
// 			return err
// 		}
// 	}
// 	for _, name := range additionalSheets {
// 		// Note: There is supposed to be only 1 sheet named "Sheet1" (seen in Excelize source code :)).
// 		// However, I wanted to cover the case this convention changes, and provide generic
// 		// default sheet deletion, not depending on its name.
// 		file.DeleteSheet(name)
// 	}
// 	currentSheet := int(report.month) + len(additionalSheets)
// 	file.SetActiveSheet(currentSheet)
// 	return file.SaveAs(filename)
// }

// func amountToString(amount float32) string {
// 	if amount >= 0 {
// 		return fmt.Sprintf("+%.2f€", amount)
// 	}
// 	return fmt.Sprintf("%.2f€", amount)
// }

// func categoryToName(categories []piggy.Category, id int) (s string) {
// 	if id == 0 {
// 		return ""
// 	}
// 	return categories[id-1].Name
// }

// func writeTitles(file *xl.File, sheet string, report Report) {
// 	columns := localization.ColumnsByLanguage(currentLanguage)
// 	titles := map[string]string{
// 		"A": columns.Category,
// 		"B": columns.Date,
// 		"C": columns.Description,
// 		"D": columns.Amount,
// 		"E": columns.CreationDate,
// 	}
// 	for col, title := range titles {
// 		cellIndex := fmt.Sprintf("%s%d", col, 1)
// 		file.SetCellValue(sheet, cellIndex, title)
// 	}
// 	endCellIndex := fmt.Sprintf("E%d", len(report.operations)+1)
// 	err := file.AutoFilter(sheet, "A1", endCellIndex, "")
// 	if err != nil {
// 		log.Println("AutoFilter error:", err)
// 	}
// 	style, _ := file.NewStyle(titleStyleJSON)
// 	file.SetCellStyle(sheet, "A1", "E1", style)
// }

// func exportMonth(file *xl.File, report Report, month time.Month) error {
// 	sheetName := localization.MonthsByLanguage(currentLanguage)[month-1]
// 	file.NewSheet(sheetName)
// 	writeTitles(file, sheetName, report)
// 	// TODO - Move this line in an outer scope;
// 	//        The style doesn't have to be duplicated for each month.
// 	evenLineStyle, err := file.NewStyle(evenLineStyleJSON)
// 	if err != nil {
// 		log.Println("Error: ", err)
// 		return err
// 	}
// 	rowIndex := 2
// 	operations := report.operations.Where(func(op piggy.Operation) bool {
// 		return op.Date.Month() == month
// 	})
// 	dateFormat := localization.DateFormatsByLanguage(currentLanguage)
// 	for _, operation := range operations {
// 		row := map[string]interface{}{
// 			"A": categoryToName(report.categories, operation.CategoryID),
// 			"B": operation.Date.Format(dateFormat),
// 			"C": operation.Description,
// 			"D": amountToString(operation.Amount),
// 			"E": operation.CreationDate.Format(dateFormat),
// 		}
// 		for col, cell := range row {
// 			cellIndex := fmt.Sprintf("%s%d", col, rowIndex)
// 			file.SetCellValue(sheetName, cellIndex, cell)
// 			if rowIndex%2 == 0 {
// 				file.SetCellStyle(sheetName, cellIndex, cellIndex, evenLineStyle)
// 			}
// 		}
// 		rowIndex++
// 	}
// 	return nil
// }
