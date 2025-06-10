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

func CreateExl(empls []EmployeeDTOResp) error {
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
		err := f.SetCellValue("Sheet1", cell(1, row), empl.ProductName)
		err = f.SetCellValue("Sheet1", cell(2, row), empl.ShortBankName)
		err = f.SetCellValue("Sheet1", cell(3, row), empl.FullBankName)
		err = f.SetCellValue("Sheet1", cell(4, row), empl.LastName)
		err = f.SetCellValue("Sheet1", cell(5, row), empl.Name)
		err = f.SetCellValue("Sheet1", cell(6, row), empl.Patronymic)
		err = f.SetCellValue("Sheet1", cell(7, row), empl.JobTitle)
		err = f.SetCellValue("Sheet1", cell(8, row), empl.Email)
		err = f.SetCellValue("Sheet1", cell(9, row), empl.ContactDate)
		err = f.SetCellValue("Sheet1", cell(10, row), empl.Phone)
		err = f.SetCellValue("Sheet1", cell(11, row), empl.ExtensionPhone)
		err = f.SetCellValue("Sheet1", cell(12, row), empl.Mobile)
		if err != nil {
			return err
		}
	}
	f.SetActiveSheet(index)
	projectRoot := GetProjectRoot()
	finFile := filepath.Join(projectRoot, "downloads", "emp.xlsx")
	if err := f.SaveAs(finFile); err != nil {
		return err
	}
}

func cell(col, row int) string {
	name, err := excelize.CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	return name
}
