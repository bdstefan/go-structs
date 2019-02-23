package main

import "testing"

const firstName string = "Joe"
const lastName string = "Doe"

var p = newPerson(firstName, lastName)

func TestNewPerson(t *testing.T) {
	if p.firstName != firstName {
		t.Errorf("Expected person to have first name %v, got %v", firstName, p.firstName)
	}

	if p.lastName != lastName {
		t.Errorf("Expected person to have last name %v, got %v", lastName, p.lastName)
	}
}

func TestUpdateName(t * testing.T) {
	const newName string = "Jimmy"
	p.updateName(newName)

	if p.firstName != newName {
		t.Errorf("Expected person to have new name %v, got %v", newName, p.firstName)
	}
}

func TestAttachContact(t *testing.T) {
	c := newContact("joe@doe.me", 500600)

	p.attachContactInfo(c)

	if p.contactInfo != c {
		t.Errorf("Expected person to have contact %+v, got %+v", c, p.contactInfo)
	}
}
