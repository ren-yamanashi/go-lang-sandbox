package mypkg

type Person struct {
	name string // 小文字の場合はパッケージ外からアクセス不可
	Age  int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.Age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.Age
}
