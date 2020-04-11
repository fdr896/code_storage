package core

import "testing"

func TestCreateUserWithEmptyLogin(t *testing.T) {
	validUser := &User{
		Login: "admin",
	}
	if !validUser.Valid() {
		t.Error("Valid claims that user has blank Login field")
	}

	invalidUser1 := &User{}
	if invalidUser1.Valid() {
		t.Error("Valid claims that user doesn't have blank Login field")
	}

	invalidUser2 := &User{
		Login: "\t\t\n\n  \n\n",
	}
	if invalidUser2.Valid() {
		t.Error("Valid claims that user doesn't have blank Login field")
	}
}
