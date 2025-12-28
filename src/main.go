package main

import (
	"fmt"
)

func main() {
	var person Person
	person.SetPerson("Alice", 30)
	name, age := person.GetPerson()
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	dog := Dog{name: "Buddy"}
	cat := Cat{name: "Whiskers"}
	nonAnimal := "I am not an animal"
	checkIsAnimal(dog)
	checkIsAnimal(cat)
	checkIsAnimal(nonAnimal)

	_pointer()
}
