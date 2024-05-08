package csvreader

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type OtchEntity struct {
	Num          int
	Id           int
	Principal    string
	Nomenklature string
	Quantity     float64
	Price        float64
}

func MatOtchCsvRead(path string) (result []OtchEntity, err error) {
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

	res := OtchEntity{}
	var dep string

	for i, row := range record {
		if i == 10 {
			dep = row[0]
		}
		if i > 11 {
			res.Principal = dep

			i, err = strconv.Atoi(string2NumValidator(row[0]))
			if err != nil {
				return
			}
			res.Num = i
			i, err = strconv.Atoi(string2NumValidator(row[1]))
			if err != nil {
				return
			}
			res.Id = i

			res.Nomenklature = row[4]

			f, err := strconv.ParseFloat(string2NumValidator(row[15]), 64)
			if err != nil {
				return nil, err
			}
			res.Quantity = f

			f, err = strconv.ParseFloat(string2NumValidator(row[16]), 64)
			if err != nil {
				log.Println(row)
				return nil, err
			}

			res.Price = f

			result = append(result, res)
		}

	}
	return
}

func string2NumValidator(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", ".", -1)
	if s == "" {
		s = "0"
	}
	return s
}
