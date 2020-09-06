package models

import "time"

//Person contains person structure
type Person struct {
	ID          uint      `gorm:"primary_key"`
	Name        string    `json:"name" gorm:"column:name"`
	LastName    string    `json:"lastname" gorm:"column:lastname"`
	Age         int       `json:"age" gorm:"column:age"`
	Dni         int       `json:"dni" gorm:"column:dni"`
	CreatedAt   time.Time `json:"-" gorm:"column:created_at"`
	ProcessedAt time.Time `json:"-" gorm:"column:processed_at"`
}
