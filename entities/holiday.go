package entities

import "time"

type Holiday struct {
	Id      uint64    `json:"Id" gorm:"primaryKey;autoIncrement;not null;"`
	Name    string    `json:"Name"`
	Date    time.Time `json:"Abb."`
	Regions []Region  `gorm:"many2many:holiday_regions"`
}
