package bin

import (
	"fmt"
	"strconv"
	"strings"
)

type ReportEntity struct {
	Num          int
	Id           int
	Principal    string
	Nomenklature string
	Quantity     float64
	Price        float64
}

type Report []ReportEntity

type Inflater interface {
	InflateFromCSV(data [][6]string, searchStr string) error
	GetReport() Report
}

// конструктор
func New() *Report {
	return &Report{}
}

// наполняет данными переменную типа Report
// агрументом принимает слайс данных из csv-отчета
// возвращает ошибку или nil
func (result *Report) InflateFromCSV(data [][6]string, searchStr string) error {
	data = search(data, searchStr)

	res := ReportEntity{}
	for _, row := range data {
		i, err := strconv.Atoi(row[0])
		if err != nil {
			return fmt.Errorf("parse res.Num: %s", err)
		}
		res.Num = i

		i, err = strconv.Atoi(row[1])
		if err != nil {
			return fmt.Errorf("parse res.Id: %s", err)
		}
		res.Id = i

		res.Principal = row[2]
		res.Nomenklature = row[3]

		f, err := strconv.ParseFloat(row[4], 64)
		if err != nil {
			return fmt.Errorf("parse res.Quantity: %s", err)
		}
		res.Quantity = f

		f, err = strconv.ParseFloat(row[5], 64)
		if err != nil {
			return fmt.Errorf("parse res.Price: %s", err)
		}
		res.Price = f

		*result = append(*result, res)
	}
	return nil
}

// возвращает данные интерфейса
func (result Report) GetReport() Report {
	return result
}

// функция ищет строки содержащие подстроку в слайсе [][6]string
// аргументом принимает слайс и подстроку
// возвращает новый слайс [][6]string
func search(data [][6]string, searchStr string) (result [][6]string) {
	if len(searchStr) == 0 {
		result = data
		return
	}
	for _, row := range data {
		for _, val := range row {
			if strings.Contains(strings.ToLower(val), strings.ToLower(searchStr)) {
				result = append(result, row)
			}
		}
	}
	return
}
