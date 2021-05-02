package main

import "fmt"

type contactInfo struct {
	phone   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	v := person{
		firstName: "Vitor",
		lastName:  "Messias",
		contactInfo: contactInfo{
			phone:   "13",
			zipCode: "13",
		},
	}

	v.updateName("Tais")
	v.print()
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
