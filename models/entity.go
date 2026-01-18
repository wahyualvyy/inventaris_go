package models

import "time"

type Lab struct {
	ID      	uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     	string `json:"name"`
	Location 	string `json:"location"`
	Items    	[]Item `json:"items,omitempty"`
}

type Item struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	Condition   string `json:"condition"`
	LastChecked time.Time `json:"last_checked"`

	LabID uint `json:"lab_id"`
}

type MaintenanceLog struct {
	ID 			uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ItemId 		uint `json:"item_id"`
	Status 		string `json:"status"`
	Note 		string `json:"note"`
	CheckedAt 	time.Time `json:"checked_at"`
	CheckedBy 	string `json:"checked_by"`
}

type User struct {
	ID			uint `gorm:"primaryKey;autoIncrement" json:id`
	Username	string `gorm:"unique" json:"username"`
	Password	string `json:"-"`
	Role		string `json:"role"`

}