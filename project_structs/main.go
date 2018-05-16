package main

import "fmt"

type contactInfo struct {
	email    string
	postCode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Carrey",
		contactInfo: contactInfo{
			email:    "jim@blah.com",
			postCode: "ST7 8KL",
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
