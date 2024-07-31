package model

import (
	"github.com/google/uuid"
)

type Present struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name" gorm:"unique;not null"`
	Selected bool      `json:"selected" gorm:"default:false"`
	GuestID  uuid.UUID `json:"guest_id" gorm:"type:uuid;default:null"`
}
