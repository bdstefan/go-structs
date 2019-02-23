package main

import "fmt"

type person struct {
	lastName string
	firstName string
	contactInfo
}

func newPerson(firstName, lastName string) person {
	return person {
		firstName: firstName,
		lastName: lastName,
	}
}

func (p *person) updateName(newFirstName string)  {
	(*p).firstName = newFirstName
}

func (p *person) attachContactInfo(c contactInfo) {
	(*p).contactInfo = c
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

