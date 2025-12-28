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

func checkIsAnimal(a interface{}) {
	animal, isAnimal := a.(Animal)
	if isAnimal {
		fmt.Println(animal.Speak())
	} else {
		fmt.Println("Not an animal")
	}
}
