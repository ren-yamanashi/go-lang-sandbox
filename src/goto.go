package main

import (
	"errors"
	"fmt"
)

func gotoExample() (string, error) {
	var err error
	filename := ""
	data := ""

	filename, err = getFileName()
	if err != nil {
		fmt.Println(err)
		goto Done
	}
	data, err = readFile(filename)
	if err != nil {
		fmt.Println(err)
		goto Done
	}

	fmt.Println(data)
Done:
	return data, err
}

func getFileName() (string, error) {
	return "sample.txt", nil
}

func readFile(filename ...string) (string, error) {
	for _, f := range filename {
		fmt.Println("Reading file:", f)
	}
	return "Hello World!", errors.New("Can't read file")
}
