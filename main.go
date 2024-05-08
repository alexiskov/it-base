package main

import (
	"fmt"
	"log"
	"main/bin"
)

func main() {
	var s string
	fmt.Print("Паттерн: ")
	fmt.Scanln(&s)

	res, warn, err := bin.Search(s)
	if err != nil {
		log.Println(err)
		return
	}

	for _, row := range res {
		fmt.Printf("%+v\n", row)
	}

	fmt.Println(warn)
}
