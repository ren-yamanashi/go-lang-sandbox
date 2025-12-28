package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.name
}

func (c Cat) Speak() string {
	return "Meow! My name is " + c.name
}

func checkIsAnimal(arg interface{}) {
	a, ok := arg.(Animal)
	if ok {
		fmt.Println(a.Speak())
	} else {
		fmt.Println("Not an animal")
	}
}
