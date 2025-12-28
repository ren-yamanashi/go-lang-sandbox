package main

import (
	"fmt"
	"local/mypkg"
)

func main() {
	p := mypkg.Person{}
	p.SetPerson("Alice", 30)
	name, age := p.GetPerson()
	fmt.Printf("Name: %s, Age: %d\n", name, p.Age)
}
