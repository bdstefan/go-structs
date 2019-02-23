package main

type contactInfo struct {
	email string
	zipCode int
}

func newContact(email string, zipCode int) contactInfo {
	return contactInfo {
		email: email,
		zipCode: zipCode,
	}
}
