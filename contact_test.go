package main

import "testing"

func TestNewContact(t *testing.T)  {
	const email string = "joe@doe.com"
	const zipCode int  = 500600

	c := newContact(email, zipCode)

	if c.email != email {
		t.Errorf("Expected contact to have email %v, got %v", email, c.email)
	}

	if c.zipCode !=  zipCode {
		t.Errorf("Expected contact to have zip code %v, got %v", zipCode, c.zipCode)
	}
}
