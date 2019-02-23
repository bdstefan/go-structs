package main

import "testing"

const firstName string = "Joe"
const lastName string = "Doe"

func TestNewPerson(t *testing.T) {
	p := newPerson(firstName, lastName)

	if p.firstName != firstName {
		t.Errorf("Expected person to have first name %v, got %v", firstName, p.firstName)
	}

	if p.lastName != lastName {
		t.Errorf("Expected person to have last name %v, got %v", lastName, p.lastName)
	}
}

func TestUpdateName(t * testing.T) {
	const newName string = "Jimmy"

	p := newPerson(firstName, lastName)
	p.updateName(newName)

	if p.firstName != newName {
		t.Errorf("Expected person to have new name %v, got %v", newName, p.firstName)
	}
}
