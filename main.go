package main

import (
	"fmt"
	"log"
	"main/bin"
	"main/csvreport"
)

func main() {
	var searchString string
	fmt.Scanln(&searchString)

	csvResult, err := csvreport.GetAllDAta("DATA")
	if err != nil {
		log.Println(err)
		return
	}

	var reporter bin.Inflater = bin.New()

	err = reporter.InflateFromCSV(bin.Search(csvResult, searchString))
	if err != nil {
		log.Println(err)
		return
	}

	result := reporter.GetReport()
	for _, v := range result {
		fmt.Printf("%+v\n", v)
	}
}
