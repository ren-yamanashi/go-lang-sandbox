package main

import (
	"fmt"
)

type Book struct {
	title string
}

func _new() {
	bookList := []*Book{}

	for i := 0; i < 10; i++ {
		b := new(Book)
		b.title = fmt.Sprintf("Book %d", i+1)
		bookList = append(bookList, b)
	}

	for _, b := range bookList {
		fmt.Println(b.title)
	}
}
