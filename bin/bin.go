package bin

import (
	"fmt"
	"strconv"
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
	InflateFromCSV(data [][6]string) error
	GetReport() Report
}

func New() *Report {
	return &Report{}
}

// наполняет данными переменную типа Report
// агрументом принимает путь до файла мат отчета
// возвращает ошибку или nil
func (result *Report) InflateFromCSV(data [][6]string) error {

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

func (result Report) GetReport() Report {
	return result
}
