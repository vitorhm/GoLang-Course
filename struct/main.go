package main

type contactInfo struct {
	phone   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	v := person{
		firstName: "Vitor",
		lastName:  "Messias",
		contact: contactInfo{
			phone:   "13",
			zipCode: "13",
		},
	}

	println(v)
}
