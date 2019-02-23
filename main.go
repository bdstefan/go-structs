package main

import (
	"fmt"
)

func main()  {
	bogusPerson := getBogusPerson()
	bogusPerson.print()

	fmt.Println("\n**** Here is a list of bogus persons ****")

	bogusPersons := getBogusPersons(3)
	fmt.Println(bogusPersons)
}

func getBogusPersons(size int) []person {
	var persons []person

	for i := 0; i < size; i++  {
		p := newPerson(fmt.Sprintf(" No. %d", i), "Jimmy")
		c := newContact(fmt.Sprintf("jimm%v@me.com", i), 734711 + i)
		p.attachContactInfo(c)

		persons = append(persons, p)
	}

	return persons
}

func getBogusPerson() person {
	p := newPerson("Joe", "Doe")
	c := newContact("joe.doe@gmail.com", 718123)
	p.attachContactInfo(c)

	return p
}