package models

import "time"

type Tamu struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" binding:"required"`
	Company   string    `json:"company" binding:"required"`
	Visiting  string    `json:"visiting" binding:"required"`
	IDCard    string    `json:"id_card" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
