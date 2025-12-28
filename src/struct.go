package main

type Person struct {
	name string // 小文字の場合はパッケージ外からアクセス不可
	age  int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}
