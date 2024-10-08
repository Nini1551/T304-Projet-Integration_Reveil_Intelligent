package model

import (
	"time"
)

// Alarm représente le modèle d'alarme
type Alarm struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255; not null; uniqueIndex:idx_alarm`
	RingDate  time.Time `gorm:"not null; uniqueIndex:idx_alarm"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	IsActive  bool      `gorm:"default:false"`
}
