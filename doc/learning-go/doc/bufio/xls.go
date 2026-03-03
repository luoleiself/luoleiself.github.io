package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func xlsNote() {
	fmt.Println("----------------xlsNote()----------------")
	// go get github.com/xuri/excelize/v2
	f := excelize.NewFile()

	// NewStyle
	styleInd, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "0000FF", Style: 3},
			{Type: "top", Color: "00FF00", Style: 4},
			{Type: "bottom", Color: "FFFF00", Style: 5},
			{Type: "right", Color: "FF0000", Style: 6},
			{Type: "diagonalDown", Color: "A020F0", Style: 7},
			{Type: "diagonalUp", Color: "A020F0", Style: 8},
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     true,
			TextRotation:    45,
			Vertical:        "",
			WrapText:        true,
		},
	})
	if err != nil {
		fmt.Printf("f.NewStyle() error %s\n", err.Error())
	}
	fmt.Println(styleInd)
	err = f.SetCellStyle("sheet1", "H9", "H9", styleInd)
	if err != nil {
		fmt.Printf("f.SetCellStyle() error %s\n", err.Error())
	}

	// SetActiveSheet

	// SetCell*
	f.SetCellBool("sheet1", "A1", true)
	f.SetCellDefault("sheet1", "A2", "A2")
	f.SetCellFloat("sheet1", "B1", 123.123456, 4, 32)

	f.SetCellStr("sheet1", "C1", "SetCellStr value")

	f.SetCellInt("sheet1", "A3", 1)
	f.SetCellInt("sheet1", "B3", 2)
	f.SetCellFormula("sheet1", "C3", "=SUM(A3:B3)")
	f.SetCellUint("sheet1", "G1", 250)

	f.SetCellValue("sheet1", "H1", "Hello world")

	// HyperLink
	display, tooltip := "https://example.com", "Example WebSite"
	f.SetCellHyperLink("sheet1", "D1", "https://example.com", "external", excelize.HyperlinkOpts{
		Display: &display,
		Tooltip: &tooltip,
	})

	// RichText
	f.SetRowHeight("sheet1", 4, 35)
	f.SetColWidth("sheet1", "F", "F", 44)
	f.SetCellRichText("sheet1", "F4", []excelize.RichTextRun{
		{
			Text: "bold",
			Font: &excelize.Font{
				Bold:   true,
				Color:  "2354e8",
				Family: "Times New Roman",
			},
		},
		{
			Text: " and ",
			Font: &excelize.Font{
				Family: "Times New Roman",
			},
		},
		{
			Text: "italic ",
			Font: &excelize.Font{
				Bold:   true,
				Color:  "e83723",
				Italic: true,
				Family: "Times New Roman",
			},
		},
		{
			Text: "text with color and font-family,",
			Font: &excelize.Font{
				Bold:   true,
				Color:  "2354e8",
				Family: "Times New Roman",
			},
		},
		{
			Text: "\r\nlarge text with ",
			Font: &excelize.Font{
				Size:  14,
				Color: "ad23e8",
			},
		},
		{
			Text: "strike",
			Font: &excelize.Font{
				Color:  "e89923",
				Strike: true,
			},
		},
		{
			Text: " superscript",
			Font: &excelize.Font{
				Color:     "dbc21f",
				VertAlign: "superscript",
			},
		},
		{
			Text: " and ",
			Font: &excelize.Font{
				Size:      14,
				Color:     "ad23e8",
				VertAlign: "baseline",
			},
		},
		{
			Text: "underline",
			Font: &excelize.Font{
				Color:     "23e833",
				Underline: "single",
			},
		},
		{
			Text: " subscript.",
			Font: &excelize.Font{
				Color:     "017505",
				VertAlign: "subscript",
			},
		},
	})

	fmt.Println("iterator rows:----")
	// GetRows 返回所有行
	// rows, err := f.GetRows("sheet1")
	// if err != nil {
	// 	fmt.Printf("f.GetRows() error %s\n", err.Error())
	// }
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }
	// Rows 返回一个 rows 迭代器
	rows, err := f.Rows("sheet1")
	if err != nil {
		fmt.Printf("f.Rows() error %s\n", err.Error())
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			fmt.Printf("rows.Columns() error %s\n", err.Error())
		}
		fmt.Print("range row \t")
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	if err = rows.Close(); err != nil {
		fmt.Printf("rows.Close() error %s\n", err.Error())
	}
	fmt.Println("------------")

	fmt.Println("iterator columns:----")
	// GetCols 返回所有的列
	// cols, err := f.GetCols("sheet1")
	// if err != nil {
	// 	fmt.Printf("f.GetCols() error %s\n", err.Error())
	// }
	// for _, col := range cols {
	// 	for _, rowCell := range col {
	// 		fmt.Print(rowCell, "\t")
	// 	}
	// 	fmt.Println()
	// }
	// Cols 返回一个列迭代器
	cols, err := f.Cols("sheet1")
	if err != nil {
		fmt.Printf("f.Cols() error %s\n", err.Error())
	}
	for cols.Next() {
		col, err := cols.Rows()
		if err != nil {
			fmt.Printf("cols.Rows() error %s\n", err.Error())
		}
		fmt.Print("range col \t")
		for _, rowCell := range col {
			fmt.Print(rowCell, "\t")
		}
		fmt.Println()
	}
	fmt.Println("------------")

	fmt.Print("f.GetCellFormula():----")
	res, err := f.GetCellFormula("sheet1", "C3")
	if err != nil {
		fmt.Printf("f.GetCellFormula() error %s\n", err.Error())
	}
	fmt.Print(res, "\n")

	// Create a new sheet.
	// index, err := f.NewSheet("Sheet2")
	// f.SetActiveSheet(index)
	// if err != nil {
	// 	fmt.Printf("f.NewSheet() error %s\n", err.Error())
	// }

	// Save spreadsheet by the given path.
	err = f.SaveAs("./testdata/Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("-----------------------------------------")
}
