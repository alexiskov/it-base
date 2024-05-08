package csvreader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OtchEntity struct {
	Num          int
	Id           int
	Principal    string
	Nomenklature string
	Quantity     string
	Price        string
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

			i, err = strconv.Atoi(validator(row[0]))
			if err != nil {
				return
			}
			res.Num = i
			i, err = strconv.Atoi(validator(row[1]))
			if err != nil {
				return
			}
			res.Id = i

			res.Nomenklature = row[4]
			res.Quantity = row[15]
			res.Price = row[16]

			result = append(result, res)
		}

	}
	return
}

func validator(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", ".", -1)
	return s
}
