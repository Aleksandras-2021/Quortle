package models

import "time"

type Game struct {
	ID        uint      `gorm:"primaryKey"`
	WordID    uint      `gorm:"not null"`
	Word      Word      `gorm:"foreignKey:WordID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Game) TableName() string {
	return "quortle.games"
}
