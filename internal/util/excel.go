package util

import (
	"fmt"
	"go-code-export/internal/model"

	"github.com/xuri/excelize/v2"
)

func GenerateExcel(path string, found []model.Company, notFound []string) error {
	file := excelize.NewFile()

	// Aba 1 - Encontrados
	foundSheet := "Encontrados"
	file.NewSheet(foundSheet)
	file.SetCellValue(foundSheet, "A1", "ID")
	file.SetCellValue(foundSheet, "B1", "Code")
	file.SetCellValue(foundSheet, "C1", "Nome")
	file.SetCellValue(foundSheet, "D1", "CNPJ")

	for i, c := range found {
		row := i + 2
		file.SetCellValue(foundSheet, fmt.Sprintf("A%d", row), c.ID)
		file.SetCellValue(foundSheet, fmt.Sprintf("B%d", row), c.Code)
		file.SetCellValue(foundSheet, fmt.Sprintf("C%d", row), c.Name)
		file.SetCellValue(foundSheet, fmt.Sprintf("D%d", row), c.CNPJ)
	}

	// Aba 2 - Não encontrados
	notFoundSheet := "NaoEncontrados"
	file.NewSheet(notFoundSheet)
	file.SetCellValue(notFoundSheet, "A1", "Code")

	for i, code := range notFound {
		file.SetCellValue(notFoundSheet, fmt.Sprintf("A%d", i+2), code)
	}

	file.DeleteSheet("Sheet1")
	return file.SaveAs(path)
}
