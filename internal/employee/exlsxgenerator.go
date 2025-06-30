package employee

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g Generator) CreateExl(empls []Entity) ([]byte, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("error close file %v", err)
		}
	}()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		return nil, err
	}
	headers := []string{"Семинар", "Банк (краткое)", "Банк (полное)",
		"Фамилия", "Имя", "Отчество", "Должность", "Почта", "Дата контакта", "Телефон", "Добавочный", "Мобильный"}
	for i, header := range headers {
		err := f.SetCellValue("Sheet1", cell(i+1, 1), header)
		if err != nil {
			return nil, err
		}
	}
	for i, empl := range empls {
		clean := empl.ToCallReport()
		row := i + 2
		values := []any{
			clean.ProductName,
			clean.ShortBankName,
			clean.FullBankName,
			clean.LastName,
			clean.Name,
			clean.Patronymic,
			clean.JobTitle,
			clean.Email,
			clean.ContactDate,
			clean.Phone,
			clean.ExtensionPhone,
			clean.Mobile,
		}
		for idx, value := range values {
			if err := f.SetCellValue("Sheet1", cell(idx+1, row), value); err != nil {
				return nil, err
			}
		}
	}
	f.SetActiveSheet(index)
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, fmt.Errorf("failed to write xlsx to buffer: %w", err)
	}
	return buf.Bytes(), nil
}

func cell(col, row int) string {
	name, err := excelize.CoordinatesToCellName(col, row)
	if err != nil {
		return ""
	}
	return name
}
