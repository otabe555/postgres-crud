package main

import (
	"testing"

	"github.com/otabe555/postgresql/model"
)

func TestAddContact(t *testing.T) {
	if AddContact(&model.Contact{
		Name:   "Kim Kim",
		Phone:  "+998945121286",
		Gender: "Female",
		Email:  "Kimkim@gmail.com",
	}) == -1 {
		t.Error("Failed to Add user")
	}
}

func TestUpdateContact(t *testing.T) {
	if UpdateContact(6, &model.Contact{
		Name:   "Kim Kim",
		Phone:  "+998906554088",
		Gender: "Female",
		Email:  "Kimkim@gmail.com",
	}) != 1 {
		t.Error("Failed to Update user")
	}
}

func TestDeleteContact(t *testing.T) {
	if DeleteContact(6) != 1 {
		t.Error("Failed to Delete")
	}
}

func TestGetContact(t *testing.T) {
	_, err := GetContact(7)
	if err != nil {
		t.Error("Failed to Get user", err)
	}
}
func TestGetAllContacts(t *testing.T) {
	_, err := GetAllContacts()
	if err != nil {
		t.Error("Failed to GetAll users list ", err)
	}
}
