package models

type Tabler interface {
	TableName() string
}

func (HouseHold) TableName() string  {
	return "households"
}
type HouseHold struct {
	HouseholdQid string `json:"HouseholdQid" gorm:"primary_key"`
	HouseholdName string `json:"householdName"`
	Contacts []Contact `gorm:"foreignKey:HouseholdQid;references:HouseholdQid" json:"contacts"`
}
