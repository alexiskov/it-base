package csvreport

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Читает csv файл и отдает данные в виде [][]string, или ошибку
func matOtchCsvRead(path string) (csvReport [][6]string, err error) {
	fs, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("open file error: %s", err)
		return
	}
	defer fs.Close()

	reader := csv.NewReader(fs)
	reader.Comma = ';'
	reader.LazyQuotes = true

	record, err := reader.ReadAll()
	if err != nil {
		err = fmt.Errorf("csv parsing error: %s", err)
		return
	}
	record = record[0 : len(record)-1]

	var report [6]string

	for i, row := range record {
		if i == 10 {
			report[2] = row[0]
		} else if i > 11 {
			report[0] = string2NumValidator(row[0])
			report[1] = string2NumValidator(row[1])
			report[3] = row[4]
			report[4] = string2NumValidator(row[15])
			report[5] = string2NumValidator(row[16])
			csvReport = append(csvReport, report)
		}
	}

	return
}

// подготавливает строку выгруженную из 1С к операции приведения типов:  строка -> число
func string2NumValidator(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\u00a0", "")
	s = strings.ReplaceAll(s, ",", ".")
	if s == "" {
		s = "0"
	}
	return s
}

func getFoldersPull(path string) (folders []string, err error) {
	// Открываем текущую директорию
	dir, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("directory opening error: %s", err)
		return
	}
	defer dir.Close()

	// Получаем список файлов и папок
	files, err := dir.Readdir(-1)
	if err != nil {
		err = fmt.Errorf("folder list reading error: %s", err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			folders = append(folders, f.Name())
		}
	}

	return
}

// получает данные из всех файлов в целевой директории
// возвращает слайс данных или ошибку
func GetAllDAta(path string) (results [][6]string, err error) {
	f, err := getFoldersPull(path)
	if err != nil {
		return
	}

	for _, dir := range f {
		result, err := matOtchCsvRead(path + "/" + dir + "/" + "otchet.csv")
		if err != nil {
			return nil, err
		}
		results = append(results, result...)
	}
	return
}
