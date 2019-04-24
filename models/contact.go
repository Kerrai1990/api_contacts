package models

import (
	"fmt"

	u "github.com/kerrai1990/api_contacts/utils"
)

// Contact -
type Contact struct {
	ID     uint   `gorm:"primary_key"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"`
	// CreatedAt time.Time  `json:"created_at"`
	// UpdatedAt time.Time  `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at";sql:"index"`
}

// Validate -
func (contact *Contact) Validate() (map[string]interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "No Contact Name found"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "No Phone Number found"), false
	}

	if contact.UserID <= 0 {
		return u.Message(false, "User ID is incorrect"), false
	}

	return u.Message(true, "success"), true

}

// Create -
func (contact *Contact) Create() map[string]interface{} {

	if response, ok := contact.Validate(); !ok {
		fmt.Println(response)
		return response
	}

	GetDB().Create(contact)
	response := u.Message(true, "success")
	response["contact"] = contact

	return response

}

// Get -
func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contact

}

// GetUserContacts -
func GetUserContacts(userID uint) interface{} {

	fmt.Println(userID)

	// contacts := make([]*Contact, 0)
	contacts := &Contact{}

	fmt.Println(contacts)

	// err := GetDB().Table("contacts").Where("user_id = ?", userID).Find(&contacts).Error
	err := GetDB().Table("contacts").Where("user_id = ?", userID).First(contacts).Error

	fmt.Println(err)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
