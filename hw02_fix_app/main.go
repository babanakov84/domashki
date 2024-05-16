package main

import (
	"fmt"

	"github.com/babanakov84/domashki/hw02_fix_app/printer"
	"github.com/babanakov84/domashki/hw02_fix_app/reader"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}
	staff, err := reader.ReadJSON(path)
	if err != nil {
		fmt.Println(err)
	} else {
		printer.PrintStaff(staff)
	}
}
