package main

import (
	"fmt"
	"log"
	"main/bin"
	"main/csvreport"
)

func main() {
	csvResult, err := csvreport.GetAllDAta("DATA")
	if err != nil {
		log.Println(err)
		return
	}

	var reporter bin.Inflater = bin.New()

	err = reporter.InflateFromCSV(csvResult)
	if err != nil {
		log.Println(err)
		return
	}

	result := reporter.GetReport()
	fmt.Println(result)
}
