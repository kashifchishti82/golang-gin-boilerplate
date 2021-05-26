package models

type Contact struct {
	ContactQid string `json:"contact_qid" gorm:"primary_key"`
	FirstName string
	LastName string
	DateOfBirth string
	PrimaryEmail string
	PrimaryPhone string
	PrimaryPhoneType string
	HouseholdQid string
}