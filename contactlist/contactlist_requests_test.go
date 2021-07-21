package contactlist

import (
	"testing"

	"github.com/otabe555/postgresql/model"
)

var (
	cm         model.ContactManagerInterface
	id         int
	user       = &model.Contact{Name: "Kim Kim", Phone: "998909037980", Gender: "Male", Email: "kim@gmail.com"}
	userforupd = &model.Contact{Name: "jake jake", Phone: "998995555555", Gender: "Male", Email: "kimkim@gmail.com"}
)

func TestNewContactManager(t *testing.T) {
	var err error
	cm, err = NewContactManager()
	if err != nil {
		t.Error("Failed to connect to the db")
	}
}
func TestContactManager_AddContact(t *testing.T) {
	if cm.AddContact(user) != nil {
		t.Error("Failed to Add user")
	}
	id = user.ID
}

func TestContactManager_UpdateContact(t *testing.T) {
	if cm.UpdateContact(10, userforupd) != nil {
		t.Error("Failed to Update Contact")
	}
}

func TestContactManager_DeleteContact(t *testing.T) {
	TestContactManager_AddContact(t)
	if cm.DeleteContact(id) != nil {
		t.Error("Failed to Delete the Contact")
	}
}

func TestContactManager_GetContact(t *testing.T) {
	_, err := cm.GetContact(10)
	if err != nil {
		t.Error("Failed to Get Contact", err)
	}
}
func TestContactManager_GetAllContacts(t *testing.T) {
	_, err := cm.GetAllContacts()
	if err != nil {
		t.Error("Failed to GetAll Contact list ", err)
	}
}
