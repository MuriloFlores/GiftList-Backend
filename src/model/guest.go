package model

import (
	"github.com/google/uuid"
)

type Guest struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(16);unique;not null"`
	Presents    []Present `json:"presents"`
}
