package models

import "time"

type Lab struct {
	Id       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Items    []Item `json:"items,omitempty"`
}

type Item struct {
	Id          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	Condition   string `json:"condition"`
	LastChecked time.Time `json:"last_checked"`

	LabID uint `json:"lab_id"`
}

type MaintanceLog struct {
	I uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ItemId uint `json:"item_id"`
	Status string `json:"status"`
	Note string `json:"note"`
	CheckedAt time.Time `json:"checked_at"`
	CheckedBy string `json:"checked_by"`
}