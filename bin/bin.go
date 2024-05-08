package bin

import (
	"fmt"
	"main/csvreader"
	"os"
	"strings"
)

func Search(pattern string) (result []csvreader.OtchEntity, warnings []string, err error) {
	res, warnings, err := getAllDAta("DATA")
	if err != nil {
		return
	}

	for _, row := range res {
		if strings.Contains(strings.ToLower(row.Nomenklature), strings.ToLower(pattern)) {
			result = append(result, row)
		}
	}

	return
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

func getAllDAta(path string) (results []csvreader.OtchEntity, warnings []string, err error) {
	f, err := getFoldersPull(path)
	if err != nil {
		return
	}

	for _, dir := range f {
		res, err := csvreader.MatOtchCsvRead(path + "/" + dir + "/otchet.csv")
		if err != nil {
			warnings = append(warnings, err.Error())
			continue
		} else {
			results = append(results, res...)
		}
	}
	return
}
