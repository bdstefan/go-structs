package main

import "testing"

func TestNewPerson(t *testing.T)  {
	const firstName string = "Joe"
	const lastName string = "Doe"

	p := newPerson(firstName, lastName)

	if p.firstName != firstName {
		t.Errorf("Expected person to have first name %v, got %v", firstName, p.firstName)
	}

	if p.lastName != lastName {
		t.Errorf("Expected person to have last name %v, got %v", lastName, p.lastName)
	}
}
