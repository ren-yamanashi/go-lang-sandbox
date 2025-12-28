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
		book := new(Book)
		book.title = fmt.Sprintf("Book %d", i+1)
		bookList = append(bookList, book)
	}

	for _, book := range bookList {
		fmt.Println(book.title)
	}
}
