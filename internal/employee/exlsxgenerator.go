package employee

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var projectRoot string
var once sync.Once

func GetProjectRoot() string {
	once.Do(func() {
		projectRoot = FindProjectRoot()
	})
	return projectRoot
}

func CreateExl(empls []Entity) error {
	filename := "./downloads/emp.xlsx"
	f, err := excelize.OpenFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			f = excelize.NewFile()
			if err := f.SaveAs(filename); err != nil {
				return fmt.Errorf("error save new file %w", err)
			}
		} else {
			return fmt.Errorf("%w", err)
		}

	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("error close file %v", err)
		}
	}()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		return err
	}
	headers := []string{"Семинар", "Банк (краткое)", "Банк (полное)",
		"Фамилия", "Имя", "Отчество", "Должность", "Почта", "Дата контакта", "Телефон", "Добавочный", "Мобильный"}
	for i, header := range headers {
		err := f.SetCellValue("Sheet1", cell(i+1, 1), header)
		if err != nil {
			return err
		}
	}
	var cleanEmp []EmployeeClean
	for _, empl := range empls {
		clean := empl.ToCallReport()
		cleanEmp = append(cleanEmp, clean)
	}
	for i, empl := range cleanEmp {
		row := i + 2
		values := []any{
			empl.ProductName,
			empl.ShortBankName,
			empl.FullBankName,
			empl.LastName,
			empl.Name,
			empl.Patronymic,
			empl.JobTitle,
			empl.Email,
			empl.ContactDate,
			empl.Phone,
			empl.ExtensionPhone,
			empl.Mobile,
		}
		for idx, value := range values {
			if err := f.SetCellValue("Sheet1", cell(idx+1, row), value); err != nil {
				return err
			}
		}
	}
	f.SetActiveSheet(index)
	projectRoot := GetProjectRoot()
	finFile := filepath.Join(projectRoot, "downloads", "emp.xlsx")
	if err := f.SaveAs(finFile); err != nil {
		return err
	}
	return nil
}

func cell(col, row int) string {
	name, err := excelize.CoordinatesToCellName(col, row)
	if err != nil {
		return ""
	}
	return name
}
